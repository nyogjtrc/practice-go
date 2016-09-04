package main

import "testing"

func TestHello(t *testing.T) {
	expect := "hello, world."
	value := hello()
	if value != expect {
		t.Error("expect", expect, "value", value)
		t.Errorf("Expect: %s, but it was %s.", expect, value)
	}
}
