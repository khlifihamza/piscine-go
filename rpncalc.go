package piscine

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func rpncalc() {
	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Println("Error")
		return
	} else {
		ns := strings.Split(args[0], " ")
		nss := []string{}
		for _, v := range ns {
			if v != "" {
				nss = append(nss, v)
			}
		}
		nbrs := []int{}
		for _, v := range nss {
			nbr, err := strconv.Atoi(v)
			if err == nil {
				nbrs = append(nbrs, nbr)
				continue
			}
			n := len(nbrs)
			if n < 2 {
				fmt.Println("Error")
				continue
			}
			if v == "+" {
				nbrs[n-2] += nbrs[n-1]
				nbrs = nbrs[:n-1]
			}
			if v == "-" {
				nbrs[n-2] -= nbrs[n-1]
				nbrs = nbrs[:n-1]
			}
			if v == "*" {
				nbrs[n-2] *= nbrs[n-1]
				nbrs = nbrs[:n-1]
			}
			if v == "/" {
				nbrs[n-2] /= nbrs[n-1]
				nbrs = nbrs[:n-1]
			}
			if v == "%" {
				nbrs[n-2] += nbrs[n-1]
				nbrs = nbrs[:n-1]
			}
		}
		if len(nbrs) == 1 {
			fmt.Println(nbrs[0])
		} else {
			fmt.Println("Error")
		}
	}
}
