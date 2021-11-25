package bill_decode

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"strings"
)

type BaiduResponse struct {
	Errno int    `json:"errno"`
	Msg   string `json:"msg"`
}

type BaiduOCRResponse struct {
	BaiduResponse
	Data struct {
		WordResult BillImageBaiduResult `json:"words_result"`
	} `json:"data"`
}

func (b *BaiduOCRResponse) getResult() (*BillImageBaiduResult, error) {
	if b.Errno != 0 {
		return nil, errors.New(b.Msg)
	}
	return &b.Data.WordResult, nil
}

type BillImageBaiduResult struct {
	CheckCode     string // 校验码
	InvoiceCode   string // 发票代码
	InvoiceNumber string `json:"InvoiceNum"` // 发票号码
	InvoiceDate   string // 开票日期

	TotalTax        string // 税额 单位：元
	TotalAmount     string // 合计金额(不含税)
	AmountInWords   string // 金额汉字大写
	AmountInFigures string `json:"AmountInFiguers"` // 金额数字大写

	Province    string // 省
	InvoiceType string // 发票类型
	SellerName  string // 卖方名称
}

func (b *BillImageBaiduResult) Marshal() ([]byte, error) {
	return json.Marshal(&b)
}

func (b *BillImageBaiduResult) String() string {
	content, _ := json.Marshal(b)
	return string(content)
}

type BillImageQRCodeResult struct {
	CheckCode     string `bill:"校验码"`  // 校验码
	InvoiceCode   string `bill:"发票代码"` // 发票代码
	InvoiceNumber string `bill:"发票号码"` // 发票号码
	InvoiceDate   string `bill:"开票日期"` // 开票日期
	//AmountInFigures string `bill:"合计金额"` // 合计金额
	TotalAmount string `bill:"合计金额(不含税)"` // 合计金额(不含税)
}

func (b BillImageQRCodeResult) String() string {
	var ret string
	v := reflect.ValueOf(b)
	t := reflect.TypeOf(b)
	for i := 0; i < v.NumField(); i++ {
		f := v.Field(i)
		ret += fmt.Sprintln(t.Field(i).Tag.Get("bill"), "\t", f.Interface())
	}
	return ret
}

// ParseBillImageQRCodeResult parse qrcode result into BillImageQRCodeResult
// example: 01,10,036001900111,09781653,17.70,20211017,81045826961248021535,134F,
func ParseBillImageQRCodeResult(s string) (*BillImageQRCodeResult, error) {
	partList := strings.Split(s, ",")
	if len(partList) != 9 {
		fmt.Println(s)
		return nil, errors.New("qrcode result is invalid")
	}
	var ret BillImageQRCodeResult
	ret.CheckCode = partList[6]
	ret.InvoiceCode = partList[2]
	ret.InvoiceNumber = partList[3]
	ret.InvoiceDate = partList[5]
	//ret.AmountInFigures = ""
	ret.TotalAmount = partList[4]
	return &ret, nil
}
