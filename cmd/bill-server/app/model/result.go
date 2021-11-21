package model

import "github.com/gogf/gf/v2/os/gtime"

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
}

type UserResultCreateServiceArgs struct {
	UserId   int    `orm:"user_id"   ` // 用户id
	FileHash string `orm:"file_hash" ` // 文件哈希
	FileUrl  string `orm:"file_url"  ` // 若解析失败，上传文件
	ResultId int    `orm:"result_id" ` // 结果id
}
