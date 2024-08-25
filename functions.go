package goreloaded

import (
	"fmt"
	"log"
	re "regexp"
	"strconv"
	"strings"
)

func ClearSpaces(content string) string {
	//compile regex
	re1 := re.MustCompile(`(\s+)`)

	fixedText := re1.ReplaceAllString(content, " ")

	return fixedText
}

func FixGrammar(content string) string {
	// Compile the regular expressions
	re1 := re.MustCompile(`(?i)(a)(\s+[AEIOUH]\w+)`)
	re2 := re.MustCompile(`(?i)(an)(\s+[^AEIOUH]\w+)`)

	// Fixes the grammar with "a" being used incorrectly before a word
	fixedText := re1.ReplaceAllStringFunc(content, func(match string) string {
		submatches := re1.FindStringSubmatch(match)
		if len(submatches) > 2 {
			switch submatches[1] {
			case "a":
				return fmt.Sprintf(`an%s`, submatches[2])
			case "A":
				return fmt.Sprintf(`AN%s`, submatches[2])
			}
		}
		return match // Return original if not matched correctly
	})

	// Fixes the grammar with "an" being used incorrectly before a word
	fixedText = re2.ReplaceAllStringFunc(fixedText, func(match string) string {
		submatches := re2.FindStringSubmatch(match)
		if len(submatches) > 2 {
			switch submatches[1] {
			case "an":
				return fmt.Sprintf(`a%s`, submatches[2])
			case "AN":
				return fmt.Sprintf(`A%s`, submatches[2])
			case "An":
				return fmt.Sprintf(`A%s`, submatches[2])
			}
		}
		return match // Return original if not matched correctly
	})

	return fixedText
}

func CleanSpecial(content string) string {
	// Compile the regular expressions
	re1 := re.MustCompile(`\s+([.,;:!?]+)`)
	re2 := re.MustCompile(`([.,;:!?]+)\s*(\w+)`)

	// Removing the spaces before the special characters
	fixedText := re1.ReplaceAllStringFunc(content, func(match string) string {
		submatches := re1.FindStringSubmatch(match)
		if len(submatches) > 1 {
			return fmt.Sprint(submatches[1])
		}
		return match // Return original if not matched correctly
	})

	// Removing the extra spaces after the special character and the next word
	fixedText = re2.ReplaceAllStringFunc(fixedText, func(match string) string {
		submatches := re2.FindStringSubmatch(match)
		if len(submatches) > 2 {
			return fmt.Sprintf(`%s %s`, submatches[1], submatches[2])
		}
		return match // Return original if not matched correctly
	})

	return fixedText
}

func FixQuotationSpaces(content string) string {
	// Compile the regular expressions for both single and double quotation marks
	reSingle := re.MustCompile(`(?:\s*)'\s*([^']*?)\s*'(?:\s*)`)
	reDouble := re.MustCompile(`(?:\s*)"\s*([^"]*?)\s*"(?:\s*)`)

	// Process single quotes
	fixedText := reSingle.ReplaceAllStringFunc(content, func(match string) string {
		submatches := reSingle.FindStringSubmatch(match)
		if len(submatches) >= 2 {
			return fmt.Sprintf(` '%s' `, submatches[1])
		}
		return match // Return original if not matched correctly
	})

	// Process double quotes
	fixedText = reDouble.ReplaceAllStringFunc(fixedText, func(match string) string {
		submatches := reDouble.FindStringSubmatch(match)
		if len(submatches) >= 2 {
			return fmt.Sprintf(` "%s" `, submatches[1])
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
	re2 := re.MustCompile(`(?i)hex`)
	return re2.MatchString(command)
}

func IsBin(command string) bool {
	re2 := re.MustCompile(`(?i)bin`)
	return re2.MatchString(command)
}
