package main

import "fmt"

type Subject interface {
	Register(o Observer)
	Remove(o Observer)
	Notify()
}

type Observer interface {
	Name() string
	Update(mssage string)
}

type SubjectA struct {
	observers []Observer
}

func (s *SubjectA) Register(o Observer) {
	fmt.Println("register observer", o.Name())
	s.observers = append(s.observers, o)
}

func (s *SubjectA) Remove(o Observer) {
	fmt.Println("remove observer", o.Name())
	newObs := []Observer{}
	for i := range s.observers {
		if s.observers[i] == o {
			continue
		}
		newObs = append(newObs, s.observers[i])
	}

	s.observers = newObs
}

func (s *SubjectA) Notify() {
	for i := range s.observers {
		s.observers[i].Update("new event")
	}
}

func (s *SubjectA) UpdateMessage(msg string) {}

type ObserverPrintln struct {
	name string
}

func NewObserverA(name string) *ObserverPrintln {
	return &ObserverPrintln{name}
}

func (o *ObserverPrintln) Name() string {
	return o.name
}

func (o *ObserverPrintln) Update(message string) {
	fmt.Println("observer", o.name, "Got message:", message)
}

func main() {

	sA := new(SubjectA)
	oA := NewObserverA("A")
	oB := NewObserverA("B")

	sA.Register(oA)
	sA.Register(oB)
	sA.Notify()

	sA.Remove(oA)
	sA.Notify()
}
