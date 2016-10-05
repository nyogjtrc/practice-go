package main

import (
	"reflect"
	"testing"
)

func TestCreateAnswerNumber(t *testing.T) {
	result := CreateAnswerNumber()
	if reflect.TypeOf(result).Kind() != reflect.String {
		t.Errorf("CreateAnswerNumber: return value should be string.")
	}
}
