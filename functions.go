package goreloaded

import (
	"fmt"
	"log"
	re "regexp"
	"strconv"
	"strings"
)

func CleanSpecial(content string) string {
	newContent := content
	// first loop for removing the spaces before the special character
	re1 := re.MustCompile(`\s+([.,;:!?]+)`)
	matches := re1.FindAllStringSubmatch(content, -1)
	for _, match := range matches {
		newContent = strings.Replace(newContent, match[0], match[1], 1)
	}

	// second loop for adding the spaces after the special characters
	re2 := re.MustCompile(`(\w+[.,;:!?]+)`)
	matches = re2.FindAllStringSubmatch(newContent, -1)
	for _, match := range matches {
		newContent = strings.Replace(newContent, match[0], match[1]+" ", -1)
	}
	return newContent
}



// Converts the string to an int
func StringToInt(s string) int {
	value, err := strconv.Atoi(s)
	if err != nil {
		return 1 // return 1 as a defaul value if there is an error
	}
	return value
}

func TurnBinToDec(contentSli string) string {
	decValueInt, err := strconv.ParseInt(contentSli, 2, 64)
	if err != nil {
		log.Fatal("Error Converting from binary to decimal")
	}
	return fmt.Sprint(int(decValueInt))
}

func TurnHexToDec(contentSli string) string {
	decValueInt, err := strconv.ParseInt(contentSli, 16, 64)
	if err != nil {
		log.Fatal("Error Converting from hex to decimal")
	}
	return fmt.Sprint(int(decValueInt))
}

func TurnCapital(sentence string, count int) string {
	words := strings.Fields(sentence)
	if count >= len(words) {
		for i := 0; i < len(words); i++ {
			words[i] = strings.Title(strings.ToLower(words[i]))
		}
	} else {
		for i := len(words) - 1; i > count; i-- {
			words[i] = strings.Title(strings.ToLower(words[i]))
		}
	}
	return strings.Join(words, " ")
}

func TurnUpper(sentence string, count int) string {
	words := strings.Fields(sentence)
	if count >= len(words) {
		for i := 0; i < len(words); i++ {
			words[i] = strings.ToUpper(words[i])
		}
	} else {
		for i := len(words) - 1; i >= 0; i-- {
			words[i] = strings.ToUpper(words[i])
		}
	}
	return strings.Join(words, " ")
}

func TurnLower(sentence string, count int) string {
	words := strings.Fields(sentence)
	if count >= len(words) {
		for i := 0; i < len(words); i++ {
			words[i] = strings.ToLower(words[i])
		}
	} else {
		for i := len(words) - 1; i >= 0; i-- {
			words[i] = strings.ToLower(words[i])
		}
	}
	return strings.Join(words, " ")
}

func IsUp(command string) bool {
	re2 := re.MustCompile(`(?i)up`)
	return re2.MatchString(command)
}

func IsCap(command string) bool {
	re2 := re.MustCompile(`(?i)cap`)
	return re2.MatchString(command)
}

func IsLow(command string) bool {
	re2 := re.MustCompile(`(?i)low`)
	return re2.MatchString(command)
}

func IsHex(command string) bool {
	re2 := re.MustCompile(`(?i)lhex`)
	return re2.MatchString(command)
}

func IsBin(command string) bool {
	re2 := re.MustCompile(`(?i)bin`)
	return re2.MatchString(command)
}
