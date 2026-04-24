package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		return
	}

	// Step 1: Pre-process the font into a Map
	fontMap := make(map[rune][]string)
	file, _ := os.Open("thinkertoy.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var allLines []string
	for scanner.Scan() {
		allLines = append(allLines, scanner.Text())
	}

	for i := 32; i <= 126; i++ {
		start := (i-32)*9 + 1
		if start+8 <= len(allLines) {
			fontMap[rune(i)] = allLines[start : start+8]
		}
	}

	// Step 2: Render using easy lookups
	input := strings.Split(os.Args[1], "\\n")
	for _, word := range input {
		for row := 0; row < 8; row++ {
			for _, char := range word {
				fmt.Print(fontMap[char][row])
			}
			fmt.Println()
		}
	}
}
