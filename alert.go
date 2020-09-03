package cui

import (
	"fmt"
	"os"

	"github.com/gookit/color"
)

func alertSuccess(message string, details ...string) string {
	msg := color.Success.Sprint("√ ", message, "\n")
	for _, detail := range details {
		msg += color.Primary.Sprint("- " + detail + "\n")
	}

	return msg
}

// Success prints a success message and exit status 0.
func Success(message string, details ...string) {
	msg := alertSuccess(message, details...)
	fmt.Print(msg)

	os.Exit(0)
}

func alertError(message string, err ...error) string {
	msg := color.Danger.Sprint("X ", message, "\n")
	for _, e := range err {
		msg += color.Secondary.Sprint("  " + e.Error() + "\n")
	}

	return msg
}

// Error prints a success message and exit status 1.
func Error(message string, err ...error) {
	msg := alertError(message, err...)
	fmt.Print(msg)

	os.Exit(1)
}

func alertWarning(message string, details ...string) string {
	msg := color.Warn.Sprint("! ", message, "\n")
	for _, detail := range details {
		msg += fmt.Sprint("  " + detail + "\n")
	}

	return msg
}

// Warning prints a warning message and exit status 1.
func Warning(message string, details ...string) {
	msg := alertWarning(message, details...)
	fmt.Print(msg)

	os.Exit(1)
}

func alertInfo(message string) string {
	msg := color.Primary.Sprint("* ", message, "\n")

	return msg
}

// Info prints an information message without exit.
func Info(message string) {
	msg := alertInfo(message)
	fmt.Print(msg)
}
