package functions

import (
	"fmt"
	"log"
	"strconv"
)

func TurnBin(contentSli string) string {
	binValueInt, err := strconv.ParseInt(contentSli, 2, 64)
	if err != nil {
		log.Fatal("Error Converting to binary")
	}
	return fmt.Sprint(int(binValueInt))
}
