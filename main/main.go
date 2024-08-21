package main

import (
	"fmt"
	fn "goreloaded"
	"os"
	re "regexp"
	"strings"
)

func main() {
	// Checking the number of Arguments
	if len(os.Args) < 3 {
		fmt.Println("please enter the <input file> <output file>")
		return
	}

	// Taking the arguments into variables
	input := os.Args[1]
	output := os.Args[2]

	// Reeding the input file
	data, err := os.ReadFile(input)
	if err != nil {
		fmt.Printf("Error Reading file '%s': %s\n", input, err)
		return
	}

	// Opening the output file
	file, err := os.OpenFile("output.txt", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Printf("Error opening file '%s': %s\n", output, err)
		return
	}
	defer file.Close()

	// Converting the Data to a readable form and spliting it
	content := string(data)
	newContent := content

	// Regexp command for finding the required patterns
	pattern := `(?i)((?:\w+\s*)+)(?:[.,;:!?])*\s*\(\s*(cap|up|low|bin|hex)\s*(?:,\s*(\d+)\s*)?\)`
	re := re.MustCompile(pattern)

	// Find all matches
	matches := re.FindAllStringSubmatch(content, -1)

	// Iterate through the matches and replace them with the appropriate modifications
	for _, match := range matches {
		sentence := match[1]
		command := match[2]
		count := 1 // default count is 1 if no number is provided
		if match[3] != "" {
			count = fn.StringToInt(match[3])
		}

		// Replace based on the command
		var replacement string

		if fn.IsCap(command) {
			replacement = fn.TurnCapital(sentence, count)
		} else if fn.IsUp(command) {
			replacement = fn.TurnUpper(sentence, count)
		} else if fn.IsLow(command) {
			replacement = fn.TurnLower(sentence, count)
		} else if fn.IsBin(command) {
			replacement = fn.TurnBinToDec(sentence)
		} else if fn.IsHex(command) {
			replacement = fn.TurnHexToDec(sentence)
		}

		// Replace the content with the modified word
		newContent = strings.Replace(newContent, match[0], replacement, 1)

	}

	// Cleaning the spaces before the special characters and points
	newContent = fn.CleanSpecial(newContent)

	// Writing the new content into the output file
	_, err = file.WriteString(newContent)
	if err != nil {
		fmt.Printf("Error writing to file %s: %s\n", output, err)
		return
	}
}

// words := strings.Fields(part)
// fmt.Print(words)

// Writing the new content into the output file
// _, err = file.WriteString(newContentComb)
// if err != nil {
// 	fmt.Printf("Error writing to file %s: %s\n", output, err)
// 	return
// }

// for i := 0; i < len(contentSli); i++ {
// 	if contentSli[i] == "(cap)" {
// 		fmt.Println(strings.ToUpper(contentSli[i-1]))
// 		_, err := file.WriteString(strings.ToUpper(contentSli[i-1]))
// 		if err != nil {
// 			fmt.Printf("Error writing to file %s: %s\n", output, err)
// 			return
// 		}
// 	}
// }
