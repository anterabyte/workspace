package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Enter your input here (Ctrl+D to Stop)  :- ")

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println("You typed :-",line)
	}
}
