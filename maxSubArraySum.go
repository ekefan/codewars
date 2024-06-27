package main

// The maximum sum subarray problem consists
// in finding the maximum sum of a contiguous
// subsequence in an array or list of integers
// If the list is made up of only negative numbers, return 0 instead

// Using Kadane's Algorithm
func MaximumSubarraySum(numbers []int) int {
	maxHere, maxSubSum := 0, 0

	for _, number := range numbers {
		if number > maxHere+number {
			maxHere = number
		} else {
			maxHere += number
		}
		if maxHere > maxSubSum && maxHere > 0 {
			maxSubSum = maxHere
		}
	}
	return maxSubSum
}

// func max(a, b int) int {
// 	if a > b {
// 		return a
// 	} else {
// 		return b
// 	}
// }

// func MaxSubarraySum(numbers []int) int {
// 	res, sum := 0, 0
// 	for i := range numbers {
// 		sum += numbers[i]
// 		res = max(res, sum)
// 		sum = max(sum, 0)
// 	}
// 	return res
// }
