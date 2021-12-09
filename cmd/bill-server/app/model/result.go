package model

import (
	"errors"
	"github.com/gogf/gf/v2/os/gtime"
)

type ResultCreateServiceArgs struct {
	InvoiceNumber   string      `orm:"invoice_number"   `  // 发票号码
	InvoiceCode     string      `orm:"invoice_code"     `  // 发票代码
	CheckCode       string      `orm:"check_code"       `  // 校验码
	TotalTax        string      `orm:"total_tax"        `  // 总税额
	InvoiceDate     *gtime.Time `orm:"invoice_date"     `  // 开票日期
	AmountInWords   string      `orm:"amount_in_words"   ` // 总金额 文字
	TotalAmount     string      `orm:"total_amount"     `  // 合计金额（不含税）
	AmountInFigures string      `orm:"amount_in_figures" ` // 总金额 数字
	Province        string      `orm:"province"        `   // 省份
	InvoiceType     string      `orm:"invoice_type"     `  // 发票类型
	SellerName      string      `orm:"seller_name"      `  // 卖方名称
	ParseType       string      `orm:"parse_type"       `  // qrcode\baidu\ocr
	ErrMsg          string      `orm:"err_msg"`            // empty for success
}

type UserResultCreateServiceArgs struct {
	UserId       int    `orm:"user_id"   ` // 用户id
	FileHash     string `orm:"file_hash" ` // 文件哈希
	FileUrl      string `orm:"file_url"  ` // 若解析失败，上传文件
	ResultId     int    `orm:"result_id" ` // 结果id
	CollectionId string // 集合id
}

type ResultParseApiArgs struct {
	CollectionId string `v:"required#集合id是必要的"`
}

type ResultVO struct {
	InvoiceNumber string `json:"invoice_number"   ` // 发票号码
	InvoiceCode   string `json:"invoice_code"     ` // 发票代码
	CheckCode     string `json:"check_code"       ` // 校验码
	InvoiceDate   string `json:"invoice_date"     ` // 开票日期
	TotalAmount   string `json:"total_amount"     ` // 合计金额（不含税）
	InvoiceType   string `json:"invoice_type"     ` // 发票类型
	SellerName    string `json:"seller_name"      ` // 卖方名称
	ParseType     string `json:"parse_type"       ` // qrcode\baidu\ocr
	FileURL       string `json:"file_url"`
}

func ToResultVO(result *Result) (*ResultVO, error) {
	if result.ErrMsg != "" {
		return nil, errors.New(result.ErrMsg)
	}
	timeStr := result.InvoiceDate.Format("Ymd")
	ret := ResultVO{
		InvoiceNumber: result.InvoiceNumber,
		InvoiceCode:   result.InvoiceCode,
		CheckCode:     result.CheckCode,
		InvoiceDate:   timeStr,
		TotalAmount:   result.TotalAmount,
		InvoiceType:   result.InvoiceType,
		SellerName:    result.SellerName,
		ParseType:     result.ParseType,
	}
	return &ret, nil
}
