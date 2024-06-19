package main

import (
	"fmt"
	"strings"
	"strconv"
)
var (
	highest int
	lowest int
)
func HighAndLow(in string) string {
	numbers := strings.Fields(in)
	highest, _  = strconv.Atoi(numbers[0])
	lowest = highest
	for _, l := range(numbers) {
		number, _  := strconv.Atoi(l)
		if number < lowest { lowest = number }
		if number > highest { highest = number }
	}

	return fmt.Sprintf("%v, %v", highest, lowest)
}

func main() {
	fmt.Printf("%v", HighAndLow("8 3 -5 42 -1 0 0 -9 4 7 4 -4"))
}
