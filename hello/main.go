package main

import "fmt"

func main() {
	fmt.Printf(hello() + "\n")
}

func hello() string {
	return "hello, world."
}
