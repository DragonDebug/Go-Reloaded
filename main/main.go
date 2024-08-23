package main

import (
	"fmt"
	fn "goreloaded"
	"os"
	re "regexp"
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
	file, err := os.OpenFile("outputfile.txt", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Printf("Error opening file '%s': %s\n", output, err)
		return
	}
	defer file.Close()

	// Converting the Data to a readable form and spliting it
	content := string(data)

	// Regexp command for finding the required patterns
	pattern := `(?i)((?:\w+\s*)+)(?:[.,;:!?])*\s*\(\s*(cap|up|low|bin|hex)\s*(?:,\s*(\d+)\s*)?\)`
	reMain := re.MustCompile(pattern)

	// Iterate through the matches and replace them with the appropriate modifications
	newContent := reMain.ReplaceAllStringFunc(content, func(match string) string {
		submatches := reMain.FindStringSubmatch(match)
		if len(submatches) >= 3 {
			sentence := submatches[1]
			command := submatches[2]
			count := 1 // default count is 1 if no number is provided
			if submatches[3] != "" {
				count = fn.StringToInt(submatches[3])
			}

			// Check the command and make the correct modifications
			if fn.IsCap(command) {
				return fn.TurnCapital(sentence, count)
			} else if fn.IsUp(command) {
				return fn.TurnUpper(sentence, count)
			} else if fn.IsLow(command) {
				return fn.TurnLower(sentence, count)
			} else if fn.IsBin(command) {
				return fn.TurnBinToDec(sentence)
			} else if fn.IsHex(command) {
				return fn.TurnHexToDec(sentence)
			}
		}
		return match // Return original if not matched correctly
	})

	// Cleaning the spaces before the special characters and points
	newContent = fn.CleanSpecial(newContent)

	// Fixing the quotation spaces
	newContent = fn.FixQuotationSpaces(newContent)

	// Fixing the grammer of a and an
	newContent = fn.FixGrammar(newContent)

	// Writing the new content into the output file
	_, err = file.WriteString(newContent)
	if err != nil {
		fmt.Printf("Error writing to file %s: %s\n", output, err)
		return
	}
}
