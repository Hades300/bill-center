package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/hades300/bill-center/cmd/bill-server/app/dao"
	"github.com/hades300/bill-center/cmd/bill-server/app/model"
	"github.com/hades300/bill-center/cmd/bill-server/app/task"
	"github.com/hades300/bill-center/cmd/bill-server/library/convert"
	"github.com/hades300/bill-center/cmd/bill-server/library/utils"
	Decode "github.com/hades300/bill-center/pkg/bill-decode"
)

var (
	ErrNoExtFound         = errors.New("no file ext is found")
	ErrFileTypeNotAllowed = errors.New("only pdf or image(jpeg\\jpg\\png) is allowed")
)

const (
	MaxFileDayLife = 30
)

type resultServiceI interface {
	cachedResult(fileHash string) (result *model.Result, err error)
	saveResult(file string, recordArgs *model.UserResultCreateServiceArgs, result *model.ResultCreateServiceArgs) error

	getFileURL(resultId int64) (u string, err error)

	Parse(file string, collectionId string) (*model.Result, error)
	parseByBaidu(filename string) (result *model.Result, err error)
	parseByQrcode(filepath string) (*model.Result, error)
	parseByOCR(filepath string) (result *model.Result, err error)
}

type resultService struct {
	client *Decode.BdClient
}

var _ resultServiceI = (*resultService)(nil)

var Result = NewResultService()

func NewResultService() resultServiceI {
	return &resultService{
		client: Decode.DefaultClient,
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
func (r *resultService) saveResult(file string, recordArgs *model.UserResultCreateServiceArgs, result *model.ResultCreateServiceArgs) error {
	var (
		resultId int64
		err      error
	)
	bizFunc := func(ctx context.Context, tx *gdb.TX) error {
		// insert result
		resultId, err = dao.Result.Ctx(nil).InsertAndGetId(result)
		if err != nil {
			return err
		}

		// insert user_result
		recordArgs.ResultId = int(resultId)
		if _, err = dao.UserResult.Ctx(nil).Insert(recordArgs); err != nil {
			return err
		}
		return nil
	}
	err = dao.UserResult.Transaction(nil, bizFunc)
	if err != nil {
		return err
	}
	// upload and set ttl
	go func() {
		u, err := task.Upload(file)
		if err == nil {
			_ = task.DeleteAfterDay(u, MaxFileDayLife)
		}
		_, err = dao.UserResult.Ctx(gctx.New()).Update(g.Map{"file_url": u}, "result_id = ?", resultId)
		if err != nil {
			// save file url in cache
			//_ = cache.Set(gctx.New(),recordArgs.FileHash,u,0)
			_ = cache.Set(gctx.New(), fileUrlKey(resultId), u, 0)
		}
	}()
	return nil
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
	ret.InvoiceDate = gtime.NewFromStrFormat(resp.InvoiceDate, "Y年n月d日")
	ret.InvoiceNumber = resp.InvoiceNumber
	ret.Province = resp.Province
	ret.SellerName = resp.SellerName
	ret.TotalTax = resp.TotalTax
	ret.InvoiceType = resp.InvoiceType

	return &ret, nil
}

// Parse fast return result if already in db(query by file hash),else save result(including file url) after recognize.
func (r *resultService) Parse(file string, collectionId string) (*model.Result, error) {
	// calc file hash
	hash, err := utils.CalcFileHash(file)
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
		result, err = r.parseByQrcode(file)
	case "jpeg", "jpg", "png", "jfif":
		result, err = r.parseByQrcode(file)
		if err != nil {
			result, err = r.parseByBaidu(file)
		}
	default:
		return nil, ErrFileTypeNotAllowed
	}
	if err != nil {
		result = &model.Result{ErrMsg: err.Error()}
	}
	//cache result
	recordArgs := model.UserResultCreateServiceArgs{
		UserId:       0,
		FileHash:     hash,
		FileUrl:      "",
		ResultId:     0,
		CollectionId: collectionId,
	}
	var resultArgs model.ResultCreateServiceArgs
	if err := convert.Transform(result, &resultArgs); err != nil {
		return nil, err
	}
	resultArgs.InvoiceDate = result.InvoiceDate
	if err := r.saveResult(file, &recordArgs, &resultArgs); err != nil {
		return nil, err
	}
	return result, err
}

func (r *resultService) parseByQrcode(filepath string) (*model.Result, error) {
	ext := gfile.ExtName(filepath)
	var (
		s   string
		err error
	)
	// get string from qrcode
	if ext == "pdf" {
		s, err = Decode.GetFirstQrCodeFromPDF(filepath)
	}
	s, err = Decode.ParseQRCodeFromImage(filepath)
	if err != nil {
		return nil, err
	}
	// parse string to bill result
	ret, err := Decode.ParseBillImageQRCodeResult(s)
	if err != nil {
		return nil, err
	}
	// convert to uni model.Result
	var result model.Result
	err = convert.Transform(ret, &result)
	if err != nil {
		return nil, err
	}
	// fix time type
	result.InvoiceDate = gtime.NewFromStrFormat(ret.InvoiceDate, "Ymd")
	result.ParseType = "qrcode"
	return &result, nil
}

func (r *resultService) parseByOCR(filepath string) (*model.Result, error) {
	panic("implement me")
}

func (r *resultService) getFileURL(resultId int64) (u string, err error) {
	// get from cache
	link, err := cache.Get(gctx.New(), fileUrlKey(resultId))
	if link != nil {
		return link.String(), err
	}
	var userResult model.UserResult
	err = dao.UserResult.Ctx(gctx.New()).Where("result_id = ?", resultId).Scan(&userResult)
	if err != nil {
		return "", err
	}
	err = cache.Set(gctx.New(), fileUrlKey(resultId), userResult.FileUrl, 0)
	if err != nil {
		// save url in cache
		return "", err
	}
	return userResult.FileUrl, nil
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

func fileUrlKey(resultId int64) string {
	return fmt.Sprintf("%d-file_url", resultId)
}
