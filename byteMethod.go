package main

import (
	"bytes"
	"os"
)

func main() {
	data, _ := os.ReadFile("standard.txt")
	// Treat the entire file as bytes to avoid string overhead
	lines := bytes.Split(data, []byte("\n"))
	input := []byte(os.Args[1])

	var buf bytes.Buffer
	for row := 0; row < 8; row++ {
		for _, char := range input {
			idx := (int(char-32) * 9) + 1 + row
			buf.Write(lines[idx])
		}
		buf.WriteByte('\n')
	}
	os.Stdout.Write(buf.Bytes())
}
