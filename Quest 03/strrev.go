package piscine

func StrRev(s string) string {
	var rev string
	strbyte := []byte(s)
	j := 0
	for i := len(s) - 1; i >= 0; i-- {
		strbyte[j] = s[i]
		j++
	}
	rev = string(strbyte)
	return rev
}
