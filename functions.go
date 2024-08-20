package goreloaded

import (
	"fmt"
	"log"
	re "regexp"
	"strconv"
	"strings"
)

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

func TurnCapital(word string, count int) string {
	words := strings.Fields(word)
	for i := len(words) - 1; i >= 0 && i >= len(words)-count; i-- {
		words[i] = strings.Title(strings.ToLower(words[i]))
	}
	return strings.Join(words, " ")
}

func TurnUpper(word string, count int) string {
	words := strings.Fields(word)
	for i := len(words) - 1; i >= 0 && i >= len(words)-count; i-- {
		words[i] = strings.ToUpper(words[i])
	}
	return strings.Join(words, " ")
}

func TurnLower(word string, count int) string {
	words := strings.Fields(word)
	for i := len(words) - 1; i >= 0 && i >= len(words)-count; i-- {
		words[i] = strings.ToLower(words[i])
	}
	return strings.Join(words, " ")
}

func IsUp(match string) bool {
	re2 := re.MustCompile(`(?i)\(\s*up\s*(?:,\s*(\d+)\s*)?\)`)
	return re2.MatchString(match)
}

func IsCap(match string) bool {
	re2 := re.MustCompile(`(?i)\(\s*cap\s*(?:,\s*(\d+)\s*)?\)`)
	return re2.MatchString(match)
}

func IsLow(match string) bool {
	re2 := re.MustCompile(`(?i)\(\s*low\s*(?:,\s*(\d+)\s*)?\)`)
	return re2.MatchString(match)
}

func IsHex(match string) bool {
	re2 := re.MustCompile(`(?i)\(\s*hex\s*(?:,\s*(\d+)\s*)?\)`)
	return re2.MatchString(match)
}

func IsBin(match string) bool {
	re2 := re.MustCompile(`(?i)\(\s*bin\s*(?:,\s*(\d+)\s*)?\)`)
	return re2.MatchString(match)
}
