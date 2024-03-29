package piscine

func LoafOfBread(str string) string {
	if len(str) == 0 {
		return "\n"
	} else if len(str) < 5 {
		return "Invalid Output\n"
	} else {
		res := ""
		c := 0
		for _, char := range str {
			if c%6 != 5 && char == ' ' {
				continue
			}
			if c%6 == 5 {
				res += " "
			} else {
				res += string(char)
			}
			c++
		}
		for i := len(res) - 1; i >= 0; i-- {
			if res[i] != ' ' {
				res = res[:i+1]
				break
			}
		}
		return res + "\n"
	}
}
