package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	// Define a pattern to match text that you want to remove
	// Example: remove all occurrences of "remove this"
	pattern := `remove this`
	text := "Please remove this text. And also remove this one."

	// Compile the pattern into a regex
	re := regexp.MustCompile(pattern)

	// Remove the matched text by replacing it with an empty string
	newText := re.ReplaceAllString(text, "")
	newText = strings.Title(newText)

	fmt.Println("Modified text:", newText)
}
