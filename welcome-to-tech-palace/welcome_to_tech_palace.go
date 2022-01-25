package techpalace

import (
	"fmt"
	"strings"
)

const welcomeMessageText = "Welcome to the Tech Palace, %v"

const messageWithBorder = "%v\n%v\n%v"

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
	return fmt.Sprintf(
		welcomeMessageText,
		strings.ToUpper(customer),
	)
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	border := ""
	for i := 1; i <= numStarsPerLine; i++ {
		border += "*"
	}
	return fmt.Sprintf(
		messageWithBorder,
		border,
		welcomeMsg,
		border,
	)
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
	return strings.TrimSpace(
		strings.ReplaceAll(oldMsg, "*", ""),
	)
}
