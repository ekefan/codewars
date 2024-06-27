package main

/*
 DESCRIPTION:
Create a function taking a positive integer between 1 and 3999 (both included) as its parameter and returning a string containing the Roman Numeral representation of that integer.

Modern Roman numerals are written by expressing each digit separately starting with the leftmost digit and skipping any digit with a value of zero. There cannot be more than 3 identical symbols in a row.
*/
import (
	// "fmt"
	"strings"
)

func intToRoman(number int) string {
	val := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	sym := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	var b strings.Builder
	for i := 0; i < len(val); i++ {
		for number >= val[i] {
			b.WriteString(sym[i])
			number -= val[i]
		}
	}
	return b.String()
}

// func main() {
// 	// Test the function with some examples
// 	fmt.Println(intToRoman(1990)) // Output: MCMXC
// 	fmt.Println(intToRoman(2008)) // Output: MMVIII
// 	fmt.Println(intToRoman(1666)) // Output: MDCLXVI
// }
