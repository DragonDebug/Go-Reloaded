package goreloaded

import (
	"fmt"
	"log"
	re "regexp"
	"strconv"
	"strings"
)

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

func TurnCapital(part string, strNum string) string {
	if len(strNum) == 0 {
		return strings.Title(part)
	}
	number, err := strconv.Atoi(strNum)
	if err != nil {
		log.Fatal("Error Converting String to int")
	}
	words := strings.Fields(part)
	if len(words) <= number {
		return strings.Title(part)
	}

	for i := number; i > 0; i++ {

	}
	return part
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
