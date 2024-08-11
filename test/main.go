package main

import (
	"fmt"
	"regexp"
)

func main() {
	content := []byte("this is no this")
	re := regexp.MustCompile(`this`)
	fmt.Println(re.FindAllIndex(content, 1))
	fmt.Println(re.FindAllIndex(content, -1))
}
