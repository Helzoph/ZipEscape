package method

import (
	"bufio"
	"log"
	"os"
	"path/filepath"
)

func FileAbs(file *os.File) string {
	abs, err := filepath.Abs(file.Name())
	if err != nil {
		log.Fatal(err)
	}

	return abs
}

func FileHeader(file *os.File) [4]byte {
	file.Seek(0, 0)
	reader := bufio.NewReader(file)
	header := make([]byte, 4)
	_, err := reader.Read(header)
	if err != nil {
		log.Fatal(err)
	}

	return [4]byte{header[0], header[1], header[2], header[3]}
}

func FileType(file *os.File) string {
	header := FileHeader(file)

	if header == [4]byte{0x50, 0x4B, 0x03, 0x04} {
		return "zip"
	} else if header == [4]byte{0x25, 0x50, 0x44, 0x46} {
		return "pdf"
	} else {
		return "unknown"
	}
}
