package main

import (
	"fmt"
	"strconv"
	"strings"
)

//.	Write a function that fixes spacing inside single quotes // "' awesome '" -> "'awesome'" // "' hello world '" -> "'hello world'" func fixSingleQuotes(text string) string { }

func fixquotes(s string) string {
	s = strings.ReplaceAll(s, "' ", "'")
	s = strings.ReplaceAll(s, " '", "'")
	return s
}

// Write a function that processes a full sentence and fixes all "a" → "an" corrections
// "There it was. A amazing rock. A honest man. A book." // -> "There it was. An amazing rock. An honest man. A book." func fixArticles(text string) string { }
func atoans(s string) string {
	s = strings.ReplaceAll(s, "A a", "An")
	s = strings.ReplaceAll(s, "a e", "an")
	s = strings.ReplaceAll(s, "a i", "an")
	s = strings.ReplaceAll(s, "A o", "An")
	s = strings.ReplaceAll(s, "a u", "an")
	s = strings.ReplaceAll(s, "a h", "an")
	return s
}

// Write a function that determines whether to use "a" or "an" before a given word
// "apple" -> "an" // "horse" -> "an" // "book" -> "a"
// "honest" -> "an" (starts with h) func aOrAn(nextWord string) string { }
func atan(s string) string {
	r := strings.ToLower(string(s[0]))
	switch r {
	case "a", "e", "i", "o", "u", "h":
		return "an " + s

	}
	return "a " + s
}

//	Write a function that fixes spacing around punctuation for a simple case
//
// Given a slice of tokens (words + punctuation), remove the space before punctuation marks.
// ["hello", ",", "world", "!"] -> "hello, world!" func joinWithPunctuation(tokens []string) string { }
func fixspace(s []string) string {
	a := strings.Join(s, " ")
	b := strings.ReplaceAll(a, " ,", ",")
	c := strings.ReplaceAll(b, " !", "!")
	return c
}

// Write a function that checks if a string is a punctuation mark from the project's list
// "," -> true, "!" -> true, "x" -> false func isPunctuation(s string) bool { }
func ispunc(s string) bool {
	return strings.ContainsAny(s, "!,.:;'")
}

// Write a function that takes a slice of words and applies uppercase to the last N words
// words = ["this", "is", "so", "exciting"], n = 2 // -> ["this", "is", "SO", "EXCITING"] func uppercaseLastN(words []string, n int) []string { }
func last(s []string, n int) []string {
	for i := len(s) - n; i < len(s); i++ {
		s[i] = strings.ToUpper(s[i])
	}
	return s
}

// 7.	Write a function that capitalizes only the first letter of a word, lowercasing the rest
// "hELLO" -> "Hello", "WORLD" -> "World" func capitalizeWord(word string) string { }
func cap(s string) string {
	return strings.ToUpper(string(s[0])) + strings.ToLower(s[1:])
}

// 8.	Write a function that converts a binary string to decimal // "10" -> 2, "1010" -> 10, "11111111" -> 255 func binToDecimal(binStr string) (int64, error) { }
func bin(b string) (int64, error) {
	return strconv.ParseInt(b, 2, 64)
}

// 10.	1.      Write a program that count how many times each word appears “go go lang exercise exercise”

func count(s string) map[string]int {
	m := make(map[string]int)
	for _, w := range strings.Fields(s) {
		m[w]++
	}
	return m
}
func main() {
	fmt.Println(fixquotes("' awesome '"))
	fmt.Println(fixquotes("' hello world '"))
	// 2
	fmt.Println(atoans("There it was. An amazing rock. An honest man. A book."))
	// 3
	fmt.Println(atan("apple"))
	fmt.Println(atan("horse"))
	fmt.Println(atan("book"))
	// 4
	fmt.Println(fixspace([]string{"hello", ",", "world", "!"}))
	// 5
	fmt.Println(ispunc(","))
	fmt.Println(ispunc("!"))
	fmt.Println(ispunc("x"))
	// 6
	fmt.Println(last([]string{"this", "is", "so", "exciting"}, 2))
	// 7
	fmt.Println(cap("hello, word"))
	// 8
	fmt.Println(bin("10"))
	fmt.Println(bin("11111111"))
	fmt.Println(bin("1010"))
	// 9

	fmt.Println(count(" execise execise"))
	fmt.Println(count("go go"))
	fmt.Println(count("lang"))

}
