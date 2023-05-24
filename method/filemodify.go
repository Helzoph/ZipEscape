package method

import (
	"fmt"
	"log"
	"os"
)

func ToZip(file *os.File) {
	file.Seek(0, 0)
	_, err := file.Write([]byte{0x50, 0x4B, 0x03, 0x04})
	if err != nil {
		log.Fatal(err)
	}
	file.Seek(0, 0)

	ModifyExt(file, ".zip")

	fmt.Println("to ZIP")
}

func ToPdf(file *os.File) {
	file.Seek(0, 0)
	_, err := file.Write([]byte{0x25, 0x50, 0x44, 0x46})
	if err != nil {
		log.Fatal(err)
	}
	file.Seek(0, 0)

	ModifyExt(file, ".pdf")

	fmt.Println("to PDF")
}

func ModifyExt(file *os.File, ext string) {
	file.Close()
	oldFileName := FileAbs(file)
	newFileName := oldFileName[:len(oldFileName)-len(ext)] + ext
	err := os.Rename(FileAbs(file), newFileName)
	if err != nil {
		log.Fatal(err)
	}
}
