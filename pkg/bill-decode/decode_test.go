package bill_decode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const pdfFilePath = "../../resource/凉茶发票.pdf"

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