package main

import "testing"

var cases = []struct {
	input  string
	expect int
}{
	{"IV", 4},
	{"CXVI", 116},
	{"X", 10},
	{"LXXXIV", 84},
	{"MD", 1500},
	{"DCCCXC", 890},
	{"MDCCC", 1800},
	{"DI", 501},
}

func TestRoman2Arabic(t *testing.T) {
	for _, c := range cases {
		result := Roman2Arabic(c.input)
		if result != c.expect {
			t.Errorf("Roman2Arabic: expect %d, but %d, when input is %s", c.expect, result, c.input)
		}
	}
}
