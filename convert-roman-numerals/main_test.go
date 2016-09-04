package main

import "testing"

var cases = []struct {
	input  string
	expect int
}{
	{"CXVI", 116},
}

func TestRoman2Arabic(t *testing.T) {
	for _, c := range cases {
		result := Roman2Arabic(c.input)
		if result != c.expect {
			t.Errorf("Roman2Arabic: expect %d, but %d, when input is %s", c.expect, result, c.input)
		}
	}
}
