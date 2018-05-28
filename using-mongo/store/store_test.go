package store

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func NewSession() *mgo.Session {
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}

	return session
}

func TestCreateUpdate(t *testing.T) {
	one := BookStore{
		Name:  "The Book Store",
		Owner: "me",
		Books: []Book{},
	}

	session := NewSession()
	defer session.Close()
	c := session.DB(DatabaseName).C(CollectionName)

	op := New(session)
	defer op.DropCollection()

	// create
	err := op.Create(one)
	assert.NoError(t, err)

	result := BookStore{}
	err = c.Find(bson.M{"name": one.Name}).One(&result)
	assert.NoError(t, err)

	assert.Equal(t, one, result)

	// update
	newOne := one
	newOne.Owner = "you"
	err = op.Update(newOne)
	assert.NoError(t, err)

	result = BookStore{}
	err = c.Find(bson.M{"name": one.Name}).One(&result)
	assert.NoError(t, err)

	assert.Equal(t, newOne, result)

	// find
	r, err := op.Find(one.Name)
	assert.NoError(t, err)
	assert.Equal(t, newOne, r)
}

func TestAddUpdateDeleteBook(t *testing.T) {
	one := BookStore{
		Name:  "The Book Store",
		Owner: "me",
		Books: []Book{},
	}

	b := Book{
		Title:  "Book One",
		Author: "Maple",
		ISBN:   "123454321",
	}

	session := NewSession()
	defer session.Close()

	op := New(session)
	defer op.DropCollection()

	err := op.Create(one)
	assert.NoError(t, err)

	// add book
	err = op.AddBook(one.Name, b)
	assert.NoError(t, err)

	expect := one
	expect.Books = append(expect.Books, b)

	result, err := op.Find(one.Name)
	assert.NoError(t, err)
	assert.Equal(t, expect, result)

	// update book
	newb := Book{
		Title:  "Book One Point Two",
		Author: "Maple",
		ISBN:   "123454321",
	}
	err = op.UpdateBook(one.Name, newb)
	assert.NoError(t, err)

	expect = one
	expect.Books = append(expect.Books, newb)

	result, err = op.Find(one.Name)
	assert.NoError(t, err)
	assert.Equal(t, expect, result)

	// delete book
	err = op.DeleteBook(one.Name, b.ISBN)
	assert.NoError(t, err)

	expect = one

	result, err = op.Find(one.Name)
	assert.NoError(t, err)
	assert.Equal(t, expect, result)
}

func TestTouchTime(t *testing.T) {

	one := BookStore{
		Name:  "The Book Store",
		Owner: "me",
		Books: []Book{},
	}

	session := NewSession()
	defer session.Close()

	op := New(session)
	defer op.DropCollection()

	// create
	err := op.Create(one)
	assert.NoError(t, err)

	err = op.TouchTime(one.Name)
	assert.NoError(t, err)

	result, err := op.Find(one.Name)
	assert.NoError(t, err)
	assert.Equal(t, time.Now().Unix(), result.UpdateAt.Unix())

}
