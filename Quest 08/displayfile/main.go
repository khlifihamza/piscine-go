package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) > 1 {
		fmt.Println("Too many arguments")
		return
	}
	if len(args) == 0 {
		fmt.Println("File name missing")
		return
	} else {
		file, _ := os.Open("quest8.txt")
		arr := make([]byte, 14)
		file.Read(arr)
		fmt.Println(string(arr))
		file.Close()
	}
}
