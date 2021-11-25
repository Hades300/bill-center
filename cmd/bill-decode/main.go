package main

import (
	"flag"
	"fmt"
	"github.com/gogf/gf/v2/os/gfile"
	billDecode "github.com/hades300/bill-center/pkg/bill-decode"
	"log"
)

var (
	pdfFile   string
	imageFile string
	dir       string
)

func init() {
	flag.StringVar(&pdfFile, "pdf", "", "bill pdf file")
	flag.StringVar(&pdfFile, "image", "", "bill image file")
	flag.StringVar(&dir, "dir", "", "bill dir")
	flag.Parse()
}

func main() {
	var r map[string]string
	if pdfFile != "" {
		r = parseInBatch([]string{pdfFile})
	} else if imageFile != "" {
		r = parseInBatch([]string{imageFile})
	} else if dir != "" {
		l, err := gfile.ScanDir(dir, "*.pdf,*.jpg,*.jpeg,*.png,*.jfif")
		if err != nil {
			log.Fatalln(err)
		}
		r = parseInBatch(l)
	}
	formatOutput(r)
}

func parseInBatch(pathList []string) map[string]string {
	ret := make(map[string]string, 0)
	for _, p := range pathList {
		var (
			t   *billDecode.BillImageQRCodeResult
			s   string
			err error
		)
		ext := gfile.ExtName(p)
		switch ext {
		case "pdf":
			s, err = billDecode.GetFirstQrCodeFromPDF(p)
		case "jfif", "jpeg", "png", "jpg":
			s, err = billDecode.ParseQRCodeFromImage(p)
		}
		if err != nil {
			ret[p] = err.Error()
			continue
		}
		t, err = billDecode.ParseBillImageQRCodeResult(s)
		if err != nil {
			ret[p] = err.Error()
			continue
		}
		ret[p] = t.String()
	}
	return ret
}

func formatOutput(ret map[string]string) {
	for p, des := range ret {
		fmt.Printf("文件名：%s\n%s\n", p, des)
	}
}
