package main

import (
	"bufio"
	"os"
)

func main() {

	file, _ := os.Open("file.txt")
	reader := bufio.NewReaderSize(file, 8*1024)
	writer := bufio.NewWriterSize(file, 16*1024)
}
