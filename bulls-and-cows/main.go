package main

import (
	"fmt"
)
import "time"
import "math/rand"

func main() {
	fmt.Println("Let's play bulls and cows")

	var playerGuess string
	answer := CreateAnswerNumber()

	fmt.Println("The answer is ready, please input 4 numbers.")

	for {
		fmt.Scanf("%s", &playerGuess)
		fmt.Printf("You input %d number : %s\n", len(playerGuess), playerGuess)

		AA, BB := CheckGuessNumber(answer, playerGuess)
		if AA == 4 {
			break
		}

		fmt.Printf("%dA%dB\n", AA, BB)
	}

	fmt.Println("4A0B, Congrate, you guess the right numbers.")

}

// CreateAnswerNumber create a random number as answer
func CreateAnswerNumber() string {
	baseStrings := [10]string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
	var randN int
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	randN = r1.Intn(10)
	ans1 := baseStrings[randN]
	baseStrings[randN], baseStrings[9] = baseStrings[9], baseStrings[randN]

	randN = r1.Intn(9)
	ans2 := baseStrings[randN]
	baseStrings[randN], baseStrings[8] = baseStrings[8], baseStrings[randN]

	randN = r1.Intn(8)
	ans3 := baseStrings[randN]
	baseStrings[randN], baseStrings[7] = baseStrings[7], baseStrings[randN]

	randN = r1.Intn(7)
	ans4 := baseStrings[randN]

	return ans1 + ans2 + ans3 + ans4
}

// CheckGuessNumber return A, B result
func CheckGuessNumber(ans string, guess string) (int, int) {
	resultA := 0
	resultB := 0

	for i := 0; i < 4; i++ {
		if ans[i] == guess[i] {
			resultA++
		}

		for j := 0; j < 4; j++ {
			if j == i {
				continue
			}
			if ans[j] == guess[i] {
				resultB++
			}
		}
	}

	return resultA, resultB
}
