package main

import (
	"bufio"
	"fmt"
	"os"
)

type ASCIIArt struct {
	FontData map[rune][]string
	Height   int
}

func NewASCIIArt(filepath string) (*ASCIIArt, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	art := &ASCIIArt{FontData: make(map[rune][]string), Height: 8}
	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	for i := 32; i <= 126; i++ {
		idx := (i-32)*9 + 1
		art.FontData[rune(i)] = lines[idx : idx+8]
	}
	return art, nil
}

func (a *ASCIIArt) Render(text string) {
	for row := 0; row < a.Height; row++ {
		for _, char := range text {
			fmt.Print(a.FontData[char][row])
		}
		fmt.Println()
	}
}

func main() {
	art, _ := NewASCIIArt("thinkertoy.txt")
	art.Render(os.Args[1])
}
