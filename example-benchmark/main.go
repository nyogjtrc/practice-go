package main

func Sum(s []int) int {
	sum := 0
	for i := 0; i < len(s); i++ {
		sum += s[i]
	}
	return sum
}

type Book struct {
	Title  string `json:"title"`
	Author string `json:"author"`
}
