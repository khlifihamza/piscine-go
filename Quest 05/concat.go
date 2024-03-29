package piscine

func Concat(str1 string, str2 string) string {
	nstr1 := []byte(str1)
	nstr2 := []byte(str2)
	ns := append(nstr1, nstr2...)
	return string(ns)
}
