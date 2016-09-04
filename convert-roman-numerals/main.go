// convert-roman-numerals
// reference: https://discuss.leetcode.com/topic/51161/convert-roman-numerals-to-arabic/6

package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) <= 1 {
		fmt.Println("convert-roman-numerals <roman>")
		return
	}
	arguments := os.Args[1]
	fmt.Println(Roman2Arabic(arguments))
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
