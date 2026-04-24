package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Print("no arguments")
	}
	banner := "thinkertoy.txt"
	if len(os.Args) == 3 {
		banner = os.Args[2]
	}

	input := os.Args[1]

	val, err := os.ReadFile(banner)
	if err != nil {
		fmt.Print("error reading %v", input)
	}

	s := strings.Split(string(val), "\n")
	f := strings.Split(input, "\\n")
	var res string
	for _, r := range f {
		for row := 0; row < 8; row++ {
			for _, t := range r {
				res += s[row+(int(t-32)*9)+1]
			}
			res += "\n"
		}
	}
	fmt.Println(res)
}
