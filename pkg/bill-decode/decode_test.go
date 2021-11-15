package bill_decode

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

const pdfFilePath = "../../resource/凉茶发票.pdf"
const imageFilePath = "../../resource/fapiao.png"

func TestGetImageFromPDF(t *testing.T) {
	imageList,err:=GetImageFromPDF(pdfFilePath)
	assert.NoError(t,err)
	t.Logf("GetImageFromPDF success imageNum:%d\n",len(imageList))
}



// test pdf file contain 3 image in page 1,only the third image is valid
func TestParseQRCode(t *testing.T){
	// get image from pdf
	imageList,err:=GetImageFromPDF(pdfFilePath)
	assert.NoError(t,err)

	for _,image:=range imageList{
		// parse qrcode
		qrCode,err:=ParseQRCode(image)
		assert.NoError(t,err)
		t.Logf("ParseQRCode success qrCode:%s\n",qrCode)
	}
}

func TestGetFirstQrCodeFromPDF(t *testing.T){
	secret,err:=GetFirstQrCodeFromPDF(pdfFilePath)
	assert.NoError(t,err)
	t.Logf("get secret:%s",secret)
}


func TestBaiduImageOCRByURL(t *testing.T){
	c:=NewBdClient()
	res,err:=c.GetTextByImageURL("https://ai.bdstatic.com/file/F58F1C22248D412FBBF1632CC5776524")
	assert.NoError(t,err)
	fmt.Print(res)
}


func TestBaiduImageOCRByFile(t *testing.T){
	c:=NewBdClient()
	res,err:=c.GetTextByImageFile(imageFilePath)
	assert.NoError(t,err)
	fmt.Print(res)
}


func TestParseBillImageQRCodeResult(t *testing.T){
	s:="01,10,036001900111,09781653,17.70,20211017,81045826961248021535,134F,"
	ret,err:=ParseBillImageQRCodeResult([]byte(s))
	assert.NoError(t,err)
	ans:=&BillImageQRCodeResult{
		CheckCode:"81045826961248021535",
		InvoiceCode: "036001900111",
		InvoiceNumber: "09781653",
		InvoiceDate: "20211017",
		TotalAmount: "17.70",
	}
	assert.Equal(t,ans,ret)
}