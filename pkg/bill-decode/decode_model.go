package bill_decode

import (
	"encoding/json"
	"errors"
	"strings"
)


type BaiduResponse struct {
	Errno int `json:"errno"`
	Msg string `json:"msg"`
}

type BaiduOCRResponse struct {
	BaiduResponse
	Data struct {
		WordResult BillImageBaiduResult `json:"words_result"`
	} `json:"data"`
}

func (b *BaiduOCRResponse) getResult() (*BillImageBaiduResult, error) {
	if b.Errno!=0{
		return nil, errors.New(b.Msg)
	}
	return &b.Data.WordResult, nil
}


type BillImageBaiduResult struct {
	CheckCode     string // 校验码
	InvoiceCode   string // 发票代码
	InvoiceNumber string // 发票号码
	InvoiceDate   string // 开票日期

	TotalTax        string // 税额
	TotalAmount     string // 合计金额(不含税)
	AmountInWords   string // 金额汉字大写
	AmountInFigures string `json:"AmountInFiguers"`// 金额数字大写

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
	CheckCode     string // 校验码
	InvoiceCode   string // 发票代码
	InvoiceNumber string // 发票号码
	InvoiceDate   string // 开票日期
	AmountInFigures string // 合计金额
	TotalAmount     string // 合计金额(不含税)
}

// parse qrcode result into BillImageQRCodeResult
// example: 01,10,036001900111,09781653,17.70,20211017,81045826961248021535,134F,
func ParseBillImageQRCodeResult(content []byte) (*BillImageQRCodeResult, error) {
	s:=string(content)
	partList:=strings.Split(s, ",")
	if len(partList)!=9{
		return nil,errors.New("qrcode result is invalid")
	}
	var ret BillImageQRCodeResult
	ret.CheckCode=partList[6]
	ret.InvoiceCode=partList[2]
	ret.InvoiceNumber=partList[3]
	ret.InvoiceDate=partList[5]
	ret.AmountInFigures=""
	ret.TotalAmount = partList[4]
	return &ret, nil
}
