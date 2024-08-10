package functions

import (
	"fmt"
	"log"
	"strconv"
)

func turnBin(contentSli string) string {
	binValueInt, err := strconv.ParseInt(contentSli, 2, 64)
	if err != nil {
		log.Fatal("Error Converting to binary")
	}
	return strconv.Itoa(int(binValueInt))
}

func thisisnot() {
	fmt.Print("this is a test")
}
