package main

import (
	"fmt"
	"github.com/flotzilla/pdf_parser")

func main() {
	pdf , err := pdf_parser.ParsePdf("CV.pdf")

	if err != nil {
		fmt.Println(err.Error())
	}
	// fmt.Println(file.GetTitle())
	fmt.Println(pdf.GetTitle(),
	pdf.GetAuthor(),
	pdf.GetCreator(),
	pdf.GetISBN(),
	pdf.GetPublishers() ,
	pdf.GetLanguages() ,
	pdf.GetPagesCount())
	

}