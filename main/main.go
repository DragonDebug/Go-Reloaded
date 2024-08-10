package main

import (
	"fmt"
	"os"
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

	for i := 0; i < len(contentSli); i++ {
		if contentSli[i] == "(cap)" {
			newContent = append(newContent, strings.Title((contentSli[i-1])))
		} else if contentSli[i] == "(up)" {
			newContent = append(newContent, strings.ToUpper(contentSli[i-1]))
		} else if contentSli[i] == "(low)" {
			newContent = append(newContent, strings.ToLower(contentSli[i-1]))
		} else if contentSli[i] == "(bin)" {
			// binValue := turnBin(contentSli[i-1])
			return
		} else if contentSli[i] == "(hex)" {
			return
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
