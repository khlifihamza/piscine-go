package main

import (
	"io"
	"os"

	"github.com/01-edu/z01"
)

func printS(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
}

func main() {
	if len(os.Args) == 1 {
		io.Copy(os.Stdout, os.Stdin)
	} else {
		args := os.Args[1:]
		for _, v := range args {
			file, err := os.Open(v)
			if err != nil {
				printS("ERROR: ")
				printS(err.Error())
				z01.PrintRune('\n')
				os.Exit(1)
			}
			arr, err := io.ReadAll(file)
			printS(string(arr))
			file.Close()
		}
	}
}
