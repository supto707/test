package main

import (
	"bytes"
	"fmt"

	"github.com/ledongthuc/pdf"
)

func main() {
	pdf.DebugOn = true
	content, err := readPdf("CV.pdf") // Read local pdf file
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(content)
	return
}

func readPdf(path string) (string, error) {
	file, read, err := pdf.Open(path)
	// remember close file
    defer file.Close()
	if err != nil {
		return "", err
	}
	var buf bytes.Buffer
    b, err := read.GetPlainText()
    if err != nil {
        return "", err
    }
    buf.ReadFrom(b)
	return buf.String(), nil
}