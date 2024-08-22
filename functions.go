package goreloaded

import (
	"fmt"
	"log"
	re "regexp"
	"strconv"
	"strings"
)

func CleanSpecial(content string) string {
	// first loop for removing the spaces before the special character
	re1 := re.MustCompile(`\s+([.,;:!?]+)`)
	matches := re1.FindAllStringSubmatch(content, -1)
	for _, match := range matches {
		content = strings.Replace(content, match[0], match[1], 1)
	}

	// second loop for adding the spaces after the special characters
	re2 := re.MustCompile(`(\w+[.,;:!?]+)`)
	matches = re2.FindAllStringSubmatch(content, -1)
	for _, match := range matches {
		content = strings.Replace(content, match[0], match[1]+" ", -1)
	}
	return content
}

func FixQuotationSpaces(content string) string {
	// Pattern to handle single quotes
	reSingle := re.MustCompile(`'\s*([^']*?)\s*'`)
	// Pattern to handle double quotes
	reDouble := re.MustCompile(`"\s*([^"]*?)\s*"`)

	// Process single quotes
	fixedText := reSingle.ReplaceAllStringFunc(content, func(match string) string {
		submatches := reSingle.FindStringSubmatch(match)
		if len(submatches) >= 2 {
			// Trim spaces inside single quotes
			return fmt.Sprintf(`'%s'`, submatches[1])
		}
		return match // Return original if not matched correctly
	})

	// Process double quotes
	fixedText = reDouble.ReplaceAllStringFunc(fixedText, func(match string) string {
		submatches := reDouble.FindStringSubmatch(match)
		if len(submatches) >= 2 {
			// Trim spaces inside double quotes
			return fmt.Sprintf(`"%s"`, submatches[1])
		}
		return match // Return original if not matched correctly
	})

	return fixedText
}

// Converts the string to an int
func StringToInt(s string) int {
	value, err := strconv.Atoi(s)
	if err != nil {
		return 1 // return 1 as a defaul value if there is an error
	}
	return value
}

func TurnBinToDec(sentence string) string {
	words := strings.Fields(sentence)
	binNumStr := words[len(words)-1]
	decValueInt, err := strconv.ParseInt(binNumStr, 2, 64)
	if err != nil {
		log.Fatal("Error Converting from binary to decimal")
	}
	newContent := strings.Join(words[:len(words)-1], " ")
	decValueStr := fmt.Sprint(int(decValueInt))
	return newContent + " " + decValueStr
}

func TurnHexToDec(sentence string) string {
	words := strings.Fields(sentence)
	HexNumStr := words[len(words)-1]
	decValueInt, err := strconv.ParseInt(HexNumStr, 16, 64)
	if err != nil {
		log.Fatal("Error Converting from hex to decimal")
	}
	newContent := strings.Join(words[:len(words)-1], " ")
	decValueStr := fmt.Sprint(int(decValueInt))
	return newContent + " " + decValueStr
}

func TurnCapital(sentence string, count int) string {
	words := strings.Fields(sentence)
	if count >= len(words) {
		for i := 0; i < len(words); i++ {
			words[i] = strings.Title(strings.ToLower(words[i]))
		}
	} else {
		for i := len(words) - 1; count > 0; i-- {
			words[i] = strings.Title(strings.ToLower(words[i]))
			count--
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
		for i := len(words) - 1; count > 0; i-- {
			words[i] = strings.ToUpper(words[i])
			count--
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
		for i := len(words) - 1; count > 0; i-- {
			words[i] = strings.ToLower(words[i])
			count--
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
