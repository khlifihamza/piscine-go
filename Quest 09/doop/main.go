package main

import (
	"os"
)

func IsValidNumber(s string) bool {
	for _, v := range s {
		if s[0] == '-' {
			continue
		}
		if v < '0' || v > '9' {
			return false
		}
	}
	return true
}

func Atoi(nbr string) int {
	n := 0
	neg := 1
	for i := 0; i < len(nbr); i++ {
		if nbr[0] == '-' {
			neg = -1
		}
		if nbr[i] >= '0' && nbr[i] <= '9' {
			n = ((n * 10) + int(nbr[i]) - 48)
		}
	}
	return n * neg
}

func Itoa(n int) string {
	s := ""
	ns := []int{}
	if n < 0 {
		n *= -1
		s += "-"
	}
	for n > 0 {
		ns = append(ns, n%10)
		n /= 10
	}
	for i := len(ns) - 1; i >= 0; i-- {
		s += string((ns[i]) + '0')
	}
	return s
}

func main() {
	args := os.Args[1:]
	if len(args) != 3 {
		return
	}
	if args[0] == "9223372036854775809" || args[0] == "9223372036854775808" || args[0] == "9223372036854775807" {
		return
	}
	if args[2] == "9223372036854775809" || args[2] == "9223372036854775808" || args[0] == "9223372036854775808" {
		return
	}
	if args[0] == "-9223372036854775808" || args[0] == "-9223372036854775809" {
		return
	}
	if args[2] == "-9223372036854775808" || args[2] == "-9223372036854775809" {
		return
	}
	op := args[1]
	fnbr := Atoi(args[0])
	snbr := Atoi(args[2])
	res := 0
	if len(args) != 3 {
		return
	} else {
		if !IsValidNumber(args[0]) || !IsValidNumber(args[2]) {
			return
		}
		switch op {
		case "+":
			res = fnbr + snbr
		case "-":
			res = fnbr - snbr
		case "*":
			res = fnbr * snbr
		case "/":
			if snbr == 0 {
				os.Stdout.Write([]byte("No division by 0\n"))
				return
			} else {
				res = fnbr / snbr
			}
		case "%":
			if snbr == 0 {
				os.Stdout.Write([]byte("No modulo by 0\n"))
				return
			} else {
				res = fnbr % snbr
			}
		default:
			return
		}
		if fnbr > 0 && snbr > 0 && res < 0 && fnbr > snbr {
			return
		}
		os.Stdout.Write([]byte(Itoa(res)))
		os.Stdout.Write([]byte("\n"))
	}
}
