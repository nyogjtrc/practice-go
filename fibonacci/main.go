package main

import "fmt"

func main() {
	fmt.Println("fibonacci")

	i := 30
	r := fib(i)

	fmt.Println(i, r)
}

func fib(i int) int {
	if i <= 1 {
		return 1
	}

	return fib(i-1) + fib(i-2)
}
