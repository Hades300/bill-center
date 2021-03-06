package service

import (
	"fmt"
	"log"
	"testing"

	"github.com/gogf/gf/v2/os/gtime"
	"github.com/hades300/bill-center/cmd/bill-server/library/convert"
	bill_decode "github.com/hades300/bill-center/pkg/bill-decode"
	"github.com/stretchr/testify/assert"
)

const testImageFile = "../../../../resource/fapiao.png"

func TestResultService_Parse(t *testing.T) {
	rs := NewResultService()
	ret, err := rs.Parse(testImageFile)
	assert.NoError(t, err)
	log.Println(ret)
}

func TestCalcFileHash(t *testing.T) {
	hash, err := calcFileHash(testImageFile)
	assert.NoError(t, err)
	log.Println(hash)
}

func TestExtractFields(t *testing.T) {
	s := bill_decode.BillImageQRCodeResult{
		CheckCode:     "81045826961248021535",
		InvoiceCode:   "036001900111",
		InvoiceNumber: "09781653",
		InvoiceDate:   "20211017",
		TotalAmount:   "17.70",
	}
	fmt.Println(convert.ExtractFields(s))
}

func TestParseTime(t *testing.T) {
	s := "2021年10月17日"
	ret := gtime.NewFromStrFormat(s, "Y年n月d日")
	assert.NotNil(t, ret)
}
