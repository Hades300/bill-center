package bill_decode

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/pdfcpu/pdfcpu/pkg/api"
	"github.com/tuotoo/qrcode"
)

const MaxImageSize = 1024 * 1024 * 10 // 10MB
var ErrNoValidQrCodeInPDF = fmt.Errorf("no valid qrcode in pdf")

// get image from pdf file
func GetImageFromPDF(filename string) ([][]byte, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	imageList, err := api.ExtractImagesRaw(f, []string{"1"}, nil)
	if err != nil {
		return nil, err
	}

	var ret [][]byte
	for _, image := range imageList {
		buff := make([]byte, MaxImageSize)
		n, err := image.Read(buff)
		if err != nil {
			// failed to read one image
			return nil, err
		}
		ret = append(ret, buff[:n])
	}
	return ret, nil
}

// get QR code from pdf file (page 1)
func GetFirstQrCodeFromPDF(filename string) (string, error) {
	imageList, err := GetImageFromPDF(filename)
	if err != nil {
		return "", err
	}
	for _, image := range imageList {
		qrCode, err := ParseQRCode(image)
		if err != nil {
			continue
		}
		return qrCode, nil
	}
	return "", ErrNoValidQrCodeInPDF
}

// parse QR code from image file
func ParseQRCodeFromImage(filename string) (string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer f.Close()
	content, err := ioutil.ReadAll(f)
	if err != nil {
		return "", err
	}
	return ParseQRCode(content)
}

// parse QR code from binary data
func ParseQRCode(content []byte) (string, error) {
	file := bytes.NewReader(content)
	img, err := qrcode.Decode(file)
	if err != nil {
		return "", err
	}
	return img.Content, nil
}
