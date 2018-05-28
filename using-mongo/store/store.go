package store

import (
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Book info
type Book struct {
	Title  string
	Author string
	ISBN   string
}

// BookStore have many books
type BookStore struct {
	Name  string
	Owner string
	Books []Book
}

// DatabaseName for stroe
const DatabaseName = "store"

// CollectionName for book store
const CollectionName = "bookstore"

// Operator control database
type Operator struct {
	c *mgo.Collection
}

// New an Operator
func New(session *mgo.Session) *Operator {
	return &Operator{
		c: session.DB(DatabaseName).C(CollectionName),
	}
}

// Create BookStore
func (o *Operator) Create(new BookStore) error {
	return o.c.Insert(new)
}

// Update BookStore
func (o *Operator) Update(new BookStore) error {
	return o.c.Update(
		bson.M{"name": new.Name},
		bson.M{"$set": bson.M{"owner": new.Owner}},
	)
}

// Find BookStore
func (o *Operator) Find(name string) (result BookStore, err error) {
	err = o.c.Find(bson.M{"name": name}).One(&result)
	return
}

// AddBook to BookStore
func (o *Operator) AddBook(name string, book Book) error {
	return o.c.Update(
		bson.M{"name": name},
		bson.M{"$addToSet": bson.M{"books": book}},
	)
}

// UpdateBook in BookStore
func (o *Operator) UpdateBook(name string, book Book) error {
	return o.c.Update(
		bson.M{"name": name, "books.isbn": book.ISBN},
		bson.M{"$set": bson.M{"books.$": book}},
	)
}

// DeleteBook in BookStore
func (o *Operator) DeleteBook(name string, isbn string) error {
	return o.c.Update(
		bson.M{"name": name},
		bson.M{"$pull": bson.M{"books": bson.M{"isbn": isbn}}},
	)
}

// DropCollection empty Collection
func (o *Operator) DropCollection() error {
	return o.c.DropCollection()
}
