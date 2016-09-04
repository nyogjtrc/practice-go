package main

import "fmt"

func main() {
	fmt.Println(Roman2Arabic(""))
}

var rMap = map[string]int{
	"I": 1,
	"V": 5,
	"X": 10,
	"L": 50,
	"C": 100,
	"D": 500,
	"M": 1000,
}

// Roman2Arabic convert roman numerals to arabic
func Roman2Arabic(s string) int {
	var preChar string
	result := 0

	for _, singleRune := range s {
		currentChar := string(singleRune)
		if preChar != "" && rMap[preChar] < rMap[currentChar] {
			result -= rMap[preChar] * 2
		}
		result += rMap[currentChar]
		preChar = currentChar
	}

	return result
}
