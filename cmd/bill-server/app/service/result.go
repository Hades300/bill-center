package service

import (
	"context"
	"crypto/hmac"
	"crypto/md5"
	"errors"
	"fmt"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/hades300/bill-center/cmd/bill-server/app/dao"
	"github.com/hades300/bill-center/cmd/bill-server/app/model"
	"github.com/hades300/bill-center/cmd/bill-server/library/convert"
	Decode "github.com/hades300/bill-center/pkg/bill-decode"
	"io/ioutil"
)

var (
	ErrNoExtFound         = errors.New("no file ext is found")
	ErrFileTypeNotAllowed = errors.New("only pdf or image(jpeg\\jpg\\png) is allowed")
)

type resultServiceI interface {
	cachedResult(fileHash string) (result *model.Result, err error)
	saveResult(recordArgs *model.UserResultCreateServiceArgs, result *model.ResultCreateServiceArgs) error

	Parse(file string) (*model.Result, error)
	//ParseRaw(typ string,reader io.Reader)(*model.Result,error)
	parseByBaidu(filename string) (result *model.Result, err error)
	parseByQrcode(filename string) (result *model.Result, err error)
	parseByOCR(filepath string) (result *model.Result, err error)
}

type resultService struct {
	client *Decode.BdClient
}

var _ resultServiceI = (*resultService)(nil)

var Result = NewResultService()

func NewResultService() resultServiceI {
	return &resultService{
		client: Decode.NewBdClient(),
	}
}

// return result if file hash exists
func (r *resultService) cachedResult(fileHash string) (*model.Result, error) {
	var userResult model.UserResult
	var result model.Result
	err := dao.UserResult.Ctx(nil).Where("file_hash=?", fileHash).Scan(&userResult)
	if err != nil {
		return nil, err
	}
	err = dao.Result.Ctx(nil).Where("id=?", userResult.ResultId).Scan(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// save result, tx: save the result and got the result id,then save the user_result
func (r *resultService) saveResult(recordArgs *model.UserResultCreateServiceArgs, result *model.ResultCreateServiceArgs) error {
	bizFunc := func(ctx context.Context, tx *gdb.TX) error {
		// insert result
		id, err := dao.Result.Ctx(nil).InsertAndGetId(result)
		if err != nil {
			return err
		}

		// insert user_result
		recordArgs.ResultId = int(id)
		if _, err = dao.UserResult.Ctx(nil).Insert(recordArgs); err != nil {
			return err
		}
		return nil
	}
	return dao.UserResult.Transaction(nil, bizFunc)
}

func (r *resultService) parseByBaidu(file string) (*model.Result, error) {
	resp, err := r.client.GetTextByImageFile(file)
	if err != nil {
		return nil, err
	}
	var ret model.Result
	ret.ParseType = "baidu"

	ret.AmountInFigures = resp.AmountInFigures
	ret.CheckCode = resp.CheckCode
	ret.AmountInWords = resp.AmountInWords
	ret.InvoiceCode = resp.InvoiceCode
	// ret.InvoiceDate = resp.InvoiceDate
	ret.InvoiceNumber = resp.InvoiceNumber
	ret.Province = resp.Province
	ret.SellerName = resp.SellerName
	ret.TotalTax = resp.TotalTax
	ret.InvoiceType = resp.InvoiceType

	return &ret, nil
}

func (r *resultService) Parse(file string) (*model.Result, error) {
	// calc file hash
	hash, err := calcFileHash(file)
	if err != nil {
		return nil, err
	}
	// found previous result in db
	if ret, err := r.cachedResult(hash); err == nil {
		return ret, nil
	}
	// parse by type
	typ, err := getFileType(file)
	if err != nil {
		return nil, err
	}
	var result *model.Result
	switch typ {
	case "pdf":
	case "jpeg", "jpg", "png", "jfif":
		result, err = r.parseByBaidu(file)
	default:
		return nil, ErrFileTypeNotAllowed
	}
	if err != nil {
		result = &model.Result{ErrMsg: err.Error()}
	}
	//cache result
	recordArgs := model.UserResultCreateServiceArgs{
		UserId:   0,
		FileHash: hash,
		FileUrl:  "",
		ResultId: 0,
	}
	var resultArgs model.ResultCreateServiceArgs
	if err := convert.Transform(result, &resultArgs); err != nil {
		return nil, err
	}
	if err := r.saveResult(&recordArgs, &resultArgs); err != nil {
		return nil, err
	}
	return result, err
}

func (r *resultService) parseByQrcode(filename string) (result *model.Result, err error) {
	panic("implement me")
}

func (r *resultService) parseByOCR(filepath string) (result *model.Result, err error) {
	panic("implement me")
}

func getFileType(filename string) (string, error) {
	pos := len(filename) - 1
	for pos >= 0 {
		if pos != 0 && pos != len(filename)-1 && filename[pos] == '.' {
			return filename[pos+1:], nil
		}
		pos--
	}
	return "", ErrNoExtFound
}

// generate file hash digest
func calcFileHash(file string) (string, error) {
	content, err := ioutil.ReadFile(file)
	if err != nil {
		return "", err
	}
	return calcContentHash(content)
}

// generate file hash digest
func calcContentHash(content []byte) (string, error) {
	h := hmac.New(md5.New, []byte("bill-center"))
	h.Write(content)
	return fmt.Sprintf("%x", h.Sum(nil)), nil
}
