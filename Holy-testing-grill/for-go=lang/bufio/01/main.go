package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	file, err := os.Open("test.txt")
	if err != nil{
		fmt.Println(" ⛔ Error :- File does not exist")
		return
	}

	defer file.Close()

	reader := bufio.NewReader(file)

	for {

		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("⚠️ Error :- file is empty")
			break
		}
		fmt.Println(line)
		
	}
}
