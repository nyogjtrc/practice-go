package main

import (
	"fmt"
	"log"

	"github.com/nyogjtrc/practice-go/using-mongo/store"
	mgo "gopkg.in/mgo.v2"
)

func main() {
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	op := store.New(session)

	err = op.Create(store.BookStore{
		Name:  "Maple Store",
		Owner: "Maple",
	})
	if err != nil {
		log.Fatal(err)
	}

	result, err := op.Find("Maple Store")
	fmt.Println(result, err)

	err = op.AddBook("Maple Store", store.Book{
		Title:  "Maple Story",
		Author: "Maple",
		ISBN:   "abc",
	})
	if err != nil {
		log.Fatal(err)
	}

	result, err = op.Find("Maple Store")
	fmt.Println(result, err)

	err = op.DropCollection()
	if err != nil {
		log.Fatal(err)
	}

}
