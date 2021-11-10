package main

import (
	"flag"
	"fmt"
)

var pdfFile string
var pdfDir string

func init(){
	flag.StringVar(&pdfFile, "pdf", "", "pdf file")
	flag.StringVar(&pdfDir, "dir", "", "pdf dir")
}

func main() {

    wordPtr := flag.String("word", "foo", "a string")

    numbPtr := flag.Int("numb", 42, "an int")
    forkPtr := flag.Bool("fork", false, "a bool")

    var svar string
    flag.StringVar(&svar, "svar", "bar", "a string var")

    flag.Parse()

    fmt.Println("word:", *wordPtr)
    fmt.Println("numb:", *numbPtr)
    fmt.Println("fork:", *forkPtr)
    fmt.Println("svar:", svar)
    fmt.Println("tail:", flag.Args())
	fmt.Println("pdfFile:", pdfFile)
	fmt.Println("pdfDir:", pdfDir)
}