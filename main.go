package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func pal(s string) (result string) {
	for _, r := range s {
		result = string(r) + result
	}
	return result
}
func hex(h string) (int64, error) {
	return strconv.ParseInt(h, 16, 64)
}
func bin(b string) (int64, error) {
	return strconv.ParseInt(b, 2, 64)
}
func ispunc(s string) bool {
	return strings.ContainsAny(s, "',.:;!")
}
func AtoAn(s string) string {
	r := strings.ToLower(string(s[0]))
	switch r {
	case "a", "e", "i", "o", "u", "h":
		return "an " + s
	default:
		return "a " + s
	}

}
func readsC(r string) (string, error) {
	data, error := os.ReadFile(r)
	if error != nil {
		return r, nil
	}
	return string(data), error
}
func count(s string) map[string]int {
	r := make(map[string]int)
	for _, w := range strings.Fields(s) {
		r[w]++
	}
	return r
}
func main() {
	fmt.Println(pal("a man a plan a canal panama"))
	fmt.Println(hex("1E"))
	fmt.Println(bin("1010"))
	fmt.Println(ispunc("."))
	fmt.Println(AtoAn("book"))
	fmt.Println(readsC("Hello Go"))
	fmt.Println(count("g0 go 0ur god"))

}
