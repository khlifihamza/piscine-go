package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strconv"
)

func main() {
	output, _ := ioutil.ReadAll(os.Stdin)
	x := 0
	y := 0
	for _, v := range output {
		if v == '\n' {
			break
		}
		x++
	}
	for _, v := range output {
		if v == '\n' {
			y++
		}
	}
	if x == 0 || y == 0 {
		fmt.Println("Not a quad function")
		return
	}
	sx := strconv.Itoa(x)
	sy := strconv.Itoa(y)
	start := "./"
	quads := []string{"quadA", "quadB", "quadC", "quadD", "quadE"}
	valid := []string{}
	for _, v := range quads {
		cmd := exec.Command(start+v, sx, sy)
		res, _ := cmd.CombinedOutput()
		if string(output) == string(res) {
			valid = append(valid, "["+v+"] "+"["+sx+"] "+"["+sy+"]")
		}
	}
	if len(valid) == 0 {
		fmt.Print("Not a quad function")
	} else {
		for i, v := range valid {
			fmt.Print(v)
			if i < len(valid)-1 {
				fmt.Printf(" || ")
			}

		}
	}
	println()
}
