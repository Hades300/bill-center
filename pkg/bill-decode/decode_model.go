package bill_decode

import (
	"encoding/json"
	"errors"
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