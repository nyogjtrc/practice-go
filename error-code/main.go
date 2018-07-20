package main

import (
	"fmt"

	"github.com/nyogjtrc/practice-go/error-code/apperror"
)

func main() {

	err := apperror.New(apperror.ErrBadRequestData, "some bad request", nil)
	fmt.Println(err.Error())

	err2 := apperror.New(apperror.ErrNotFound, "ng", err)
	fmt.Println(err2.ErrorSummry())

}
