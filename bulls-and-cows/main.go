package main

import (
	"fmt"
)

func main() {
	fmt.Println("Let's play bulls and cows")

	var playerGuess string
	answer := CreateAnswerNumber()

	fmt.Println("The answer is ready, please input 4 numbers.")

	for {
		fmt.Scanf("%s", &playerGuess)

		if CheckGuessNumber(answer, playerGuess) {
			fmt.Println("Congrate, you guess the right numbers.")
			break
		}

		fmt.Println("Sorry, wrong number, please input 4 numbers again.")
	}

}

// CreateAnswerNumber create a random number as answer
func CreateAnswerNumber() string {
	return "0000"
}

// CheckGuessNumber return true when answer is the same of guess number
func CheckGuessNumber(ans string, guess string) bool {
	return false
}
