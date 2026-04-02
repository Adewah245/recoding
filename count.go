package main

import (
	"fmt"
	"strings"
)

func countfile(s string) map[string]int {
	m := make(map[string]int)
	for _, w := range strings.Fields(s) {
		m[w]++
	}
	return m
}
func main() {
	r := count("God is Able to take care of our needs")
	for _, i := range r {
		fmt.Println(i, r)
	}
}
