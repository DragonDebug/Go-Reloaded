package goreloaded

import (
	"fmt"
	"log"
	"strconv"
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
