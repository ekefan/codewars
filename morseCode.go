package main

import (
	"strings"
	"fmt"
)
func DecodeMorse(morseCode string) []string {
	str := strings.Split(morseCode, "   ")
	return str
}

func main() {
	str := DecodeMorse("I am   a blood   a baby from my")
	for _, words := range str {
		word := strings.Split(words, " ")
		for _, w := range word {
			fmt.Printf("%s\n", w)
		}
		fmt.Printf("***\n")
	}
}
