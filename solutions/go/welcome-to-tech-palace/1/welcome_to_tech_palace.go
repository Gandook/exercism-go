package techpalace

import (
    "strings"
    "unicode/utf8"
)

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
	return "Welcome to the Tech Palace, " + strings.ToUpper(customer)
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	stars := strings.Repeat("*", numStarsPerLine)
    return stars + "\n" + welcomeMsg + "\n" + stars
}

func findTheBeginning(s string) int {
    for i, r := range s {
        if r != '*' && r != ' ' && r != '\n' {
            return i
        }
    }
    return -1
}

func reverseString(s string) string {
    r := []rune(s)

    for i, j := 0, len(r) - 1; i < j; i, j = i + 1, j - 1 {
        r[i], r[j] = r[j], r[i]
    }

    return string(r)
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
	l := findTheBeginning(oldMsg)
    runeCnt := utf8.RuneCountInString(oldMsg)
    r := runeCnt - findTheBeginning(reverseString(oldMsg))

    msgAsRunes := []rune(oldMsg)
    return string(msgAsRunes[l:r])
}
