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
	file, err := os.OpenFile(output, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Printf("Error opening file '%s': %s\n", output, err)
		return
	}
	defer file.Close()

	// Converting the Data to a readable form and spliting it
	content := string(data)
	contentSli := strings.Split(content, " ")
	newContent := []string{}

	// Regexp command for finding the required patterns
	pattern := `(?i)\(\s*(?:cap|up|low|bin|hex)\s*(?:,\s*(\d+)\s*)?\)`
	re := re.MustCompile(pattern)

	// Find all matches
	matches := re.FindAllStringSubmatch(content, -1)
	
	// parts := re.Split(content, -1)
	parts:= re.ReplaceAllString(content, func(Match string) string{



	})



	// Iterate through the matches and parts of the content
	for _, match := range matches {
		for _, part := range parts {
			if len(part)>1 && len(match)>1 {
				if fn.IsCap(match[0]) {
					newContent = append(newContent, strings.Title((contentSli[i-1])))
				} else if fn.IsUp(match[0]) {
					newContent = append(newContent, strings.ToUpper(contentSli[i-1]))
				} else if fn.IsLow(match[0]) {
					newContent = append(newContent, strings.ToLower(contentSli[i-1]))
				} else if fn.IsBin(match[0]) {
					decimalValue = fn.TurnBinToDec(contentSli[i-1])
					newContent = append(newContent, decimalValue)
				} else if fn.IsHex(match[0]) {
					decimalValue = fn.TurnHexToDec(contentSli[i-1])
					newContent = append(newContent, decimalValue)
				}
			}
		}
	}

	words := strings.Fields(part)
	fmt.Print(words)

	//

	var decimalValue string
	for i := 0; i < len(contentSli); i++ {
		if contentSli[i] == "(cap)" {
			newContent = append(newContent, strings.Title((contentSli[i-1])))
		} else if contentSli[i] == "(up)" {
			newContent = append(newContent, strings.ToUpper(contentSli[i-1]))
		} else if contentSli[i] == "(low)" {
			newContent = append(newContent, strings.ToLower(contentSli[i-1]))
		} else if contentSli[i] == "(bin)" {
			decimalValue = fn.TurnBinToDec(contentSli[i-1])
			newContent = append(newContent, decimalValue)
		} else if contentSli[i] == "(hex)" {
			decimalValue = fn.TurnHexToDec(contentSli[i-1])
			newContent = append(newContent, decimalValue)
		}
	}

	newContentComb := strings.Join(newContent, " ")

	// Writing the new content into the output file
	_, err = file.WriteString(newContentComb)
	if err != nil {
		fmt.Printf("Error writing to file %s: %s\n", output, err)
		return
	}

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
}
