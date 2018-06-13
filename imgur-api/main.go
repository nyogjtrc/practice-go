package main

import (
	"fmt"
	"log"

	"github.com/nyogjtrc/practice-go/imgur-api/imgur"
)

func main() {
	ic, err := imgur.NewFromFile("config.json")
	if err != nil {
		log.Fatal(err)
	}

	r, err := ic.MyImageCount()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v", r)
}
