package main

import (
	"errors"

	"github.com/erdaltsksn/cui"
)

func main() {
	err := errors.New("Wrapped error message")
	if err != nil {
		cui.Error("Error message", err)
	}
}
