package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(RGB(-20, 275, 125))
}

func checkInt(val int) int {
	if val > 255 {
		return 255
	}
	if val < 0 {
		return 0
	}
	return val
}

func getHex(num int) string {
	remQueue := []int{}
	var reminder int = 0

	for num >= 0 {
		reminder = num % 16
		remQueue = append(remQueue, reminder)
		num /= 16
		if num == 0 { break }
	}
	output := addHexValue(remQueue)
	if len(output) < 2 {
		output = fmt.Sprintf("0%s", output)
	}
	return output
}
func addHexValue(nums []int) string {
	hexMap := make(map[int]string)

	hexMap[10] = "A"
	hexMap[11] = "B"
	hexMap[12] = "C"
	hexMap[13] = "D"
	hexMap[14] = "E"
	hexMap[15] = "F"
	var output string
	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] > 9 {
			output += hexMap[nums[i]]
		} else {
			output += strconv.Itoa(nums[i])
		}
	}
	return output
}

func RGB(r, g, b int) string {
	r = checkInt(r)
	g = checkInt(g)
	b = checkInt(b)

	rHex := getHex(r)
	gHex := getHex(g)
	bHex := getHex(b)
	return fmt.Sprintf("%v%v%v", rHex, gHex, bHex)
}
