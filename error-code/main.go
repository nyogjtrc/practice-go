package main

import (
	"fmt"

	"github.com/nyogjtrc/practice-go/error-code/apperror"
	"github.com/nyogjtrc/practice-go/error-code/gateerror"
)

func main() {

	err := apperror.New(apperror.ErrBadRequestData, "some bad request", nil)
	fmt.Println(err.Error())

	err2 := apperror.New(apperror.ErrNotFound, "ng", err)
	fmt.Println(err2.ErrorSummry())

	errG := gateerror.New(400, "bad request", err.Code())
	fmt.Println(errG.Error())

	errG2 := gateerror.New(500, "api service error", err2.Code())
	fmt.Println(errG2.Error())

}
