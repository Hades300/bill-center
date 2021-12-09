package utils

import (
	"crypto/hmac"
	"crypto/md5"
	"fmt"
	"io/ioutil"
)

// generate file hash digest
func CalcFileHash(file string) (string, error) {
	content, err := ioutil.ReadFile(file)
	if err != nil {
		return "", err
	}
	return CalcContentHash(content)
}

// generate file hash digest
func CalcContentHash(content []byte) (string, error) {
	h := hmac.New(md5.New, []byte("bill-center"))
	h.Write(content)
	return fmt.Sprintf("%x", h.Sum(nil)), nil
}
