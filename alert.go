package cui

import (
	"fmt"
	"os"

	"github.com/gookit/color"
)

// Success prints a success message and exit status 0
func Success(message string, details ...string) {
	color.Success.Println("√", message)
	if len(details) > 0 {
		for _, detail := range details {
			color.Primary.Println("-", detail)
		}
	}
	os.Exit(0)
}

// Error prints a success message and exit status 1
func Error(message string, err ...error) {
	color.Danger.Println("X", message)
	if len(err) > 0 {
		for _, e := range err {
			fmt.Println(" ", e.Error())
		}
	}
	os.Exit(1)
}

// Warning prints a warning message and exit status 1
func Warning(message string, details ...string) {
	color.Warn.Println("!", message)
	if len(details) > 0 {
		for _, detail := range details {
			fmt.Println(" ", detail)
		}
	}
	os.Exit(1)
}

// Info prints an information message without exit
func Info(message string) {
	color.Primary.Println("*", message)
}
