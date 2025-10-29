package main

import (
	"bufio"
	"fmt"
	"os"
)

func main(){

	file := OpenFile("file.txt")
	
	defer file.Close()

	writer := bufio.NewWriter(file)

	for i := 0 ; i < 10 ; i++ {
		fmt.Fprintf(writer,"Writing values :- %d\n",i)
	}

	// Without writer.Flush buffer will not write anything to the file
	writer.Flush()
}

// OpenFile is a function that Opens returns a file if it exists, if it doesn't it creates it at the time
func OpenFile(file_name string) *os.File {

	file, err := os.Open(file_name)
	if err != nil {
		fmt.Printf("⚠️ Error :- %s does not exist\n", file_name)
		fmt.Printf("📄 Creating File :- %v\n",file_name)
		file, _ =  os.Create(file_name)
		return file
	}

	return file
}
