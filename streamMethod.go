package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	input := os.Args[1]
	file, _ := os.Open("thinkertoy.txt")
	defer file.Close()

	var font []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		font = append(font, scanner.Text())
	}

	subStrings := strings.Split(input, "\\n")
	for _, sub := range subStrings {
		for row := 0; row < 8; row++ {
			for _, char := range sub {
				// Immediately "stream" out the text to the terminal
				lineIndex := int(char-32)*9 + 1 + row
				fmt.Fprint(os.Stdout, font[lineIndex])
			}
			fmt.Fprintln(os.Stdout)
		}
	}
}
