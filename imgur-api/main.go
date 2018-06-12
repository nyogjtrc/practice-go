package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/nyogjtrc/practice-go/imgur-api/imgur"
)

func main() {

	ic, err := imgur.NewFromFile("config.json")
	if err != nil {
		log.Fatal(err)
	}

	client, err := ic.HTTPClient()
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.Get("https://api.imgur.com/3/account/me/images/count")
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(resp.Status)
	fmt.Println(string(body))
}
