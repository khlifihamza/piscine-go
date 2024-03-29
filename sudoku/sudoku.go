package main

import (
	"fmt"
	"os"

	"github.com/01-edu/z01"
)

func CheckRow(su [9][9]rune, n rune, row int) bool {
	for i := 0; i < 9; i++ {
		if su[row][i] == n {
			return true
		}
	}
	return false
}

func CheckCol(su [9][9]rune, n rune, col int) bool {
	for i := 0; i < 9; i++ {
		if su[i][col] == n {
			return true
		}
	}
	return false
}

func CheckBox(su [9][9]rune, n rune, row int, col int) bool {
	lrow := row - row%3
	lcol := col - col%3
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if su[lrow+i][lcol+j] == n {
				return true
			}
		}
	}
	return false
}

func ValidatePlacement(su [9][9]rune, n rune, row int, col int) bool {
	if !CheckRow(su, n, row) && !CheckCol(su, n, col) && !CheckBox(su, n, row, col) {
		return true
	} else {
		return false
	}
}

func IsEmpty(su [9][9]rune) (int, int, bool) {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if su[i][j] == '0' {
				return i, j, true
			}
		}
	}
	return 0, 0, false
}

func solveSudoku(su [9][9]rune) (bool, [9][9]rune) {
	row, col, ie := IsEmpty(su)
	if !ie {
		return true, su
	}

	for n := '1'; n <= '9'; n++ {
		if ValidatePlacement(su, n, row, col) {
			su[row][col] = n

			if f, res := solveSudoku(su); f {
				return true, res
			}

			su[row][col] = '0'
		}
	}

	return false, su
}

func main() {
	args := os.Args[1:]
	su := [9][9]rune{}
	if len(args) != 9 {
		fmt.Println("Error")
		return
	} else {
		for y, v := range args {
			if len(v) != 9 {
				fmt.Println("Error")
				return
			} else {
				for x, n := range v {
					if (n < '1' || n > '9') && n != '.' {
						fmt.Println("Error")
						return
					}
					if n == '.' {
						su[y][x] = '0'
					} else {
						su[y][x] = n
					}
				}
			}
		}
		for y, v := range args {
			for x := range v {
				for i := 0; i < 8; i++ {
					for j := i + 1; j < 9; j++ {
						if su[y][i] == su[y][j] && su[y][i] != '0' {
							fmt.Println("Error")
							return
						}
					}
				}
				for i := 0; i < 8; i++ {
					for j := i + 1; j < 9; j++ {
						if su[i][x] == su[j][x] && su[i][x] != '0' {
							fmt.Println("Error")
							return
						}
					}
				}
			}
		}
		for k := 0; k < 3; k++ {
			for l := 0; l < 3; l++ {
				for i := 0; i < 8; i++ {
					for j := i + 1; j < 9; j++ {
						in := 3*k + i/3
						jn := 3*l + i%3
						i := 3*k + j/3
						j := 3*l + j%3
						if su[in][jn] == su[i][j] && su[in][jn] != '0' {
							fmt.Println("Error")
							return
						}
					}
				}
			}
		}
	}

	if f, res := solveSudoku(su); f {
		for _, v := range res {
			for i, x := range v {
				z01.PrintRune(x)
				if i < len(v)-1 {
					z01.PrintRune(' ')
				}
			}
			z01.PrintRune('\n')
		}
	} else {
		fmt.Println("Error")
	}
}
