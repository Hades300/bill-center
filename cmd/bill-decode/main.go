package main

import (
	"flag"
	"fmt"
	"log"

	billDecode "github.com/hades300/bill-center/pkg/bill-decode"
)
var pdfFile string
// var pdfDir string

func init(){
	flag.StringVar(&pdfFile, "pdf", "", "bil pdf file")
	// flag.StringVar(&pdfDir, "dir", "", "bil pdf dir")
	flag.Parse()
}


func main() {
	ret,err:=billDecode.GetFirstQrCodeFromPDF(pdfFile)
	if err!=nil{
		log.Fatal(err)
	}
	fmt.Println(ret)
}