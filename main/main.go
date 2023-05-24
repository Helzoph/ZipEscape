package main

import (
	"fmt"
	"log"
	"os"
	"zipescape/method"
)

func main() {
	var (
		filepath   string
		filetype   string
		fileHeader [4]byte
	)

	fmt.Println("Please enter the file path：")

	fmt.Scanln(&filepath)

	file, err := os.OpenFile(filepath, os.O_RDWR, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	fileHeader = method.FileHeader(file)
	filetype = method.FileType(file)

	fmt.Printf("File Header: %02x, File Type: %v\n", fileHeader, filetype)

	if filetype == "zip" {
		method.ToPdf(file)
	} else if filetype == "pdf" {
		method.ToZip(file)
	} else {
		log.Fatal("unknown file type")
	}
}
