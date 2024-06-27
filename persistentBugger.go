package main

import (
	// "fmt"
)

func Persistence(n int) int {
	if n%10 == n {
		return 0
	}
	persistence := 0
	for n%10 != n {
		nums := getNums(n)
		n = getProduct(nums)
		persistence++
	}
	return persistence
}

func getProduct(n []int) int {
	product := 1
	for _, num := range n {
		product *= num
	}
	return product
}

func getNums(n int) []int {
	var individualNums []int
	for n != 0 {
		individualNums = append(individualNums, n%10)
		n /= 10
	}
	return individualNums
}

// func main() {
// 	fmt.Println(Persistence(290))
// 	fmt.Println(Persistence(9))
// 	fmt.Println(Persistence(456))
// }
