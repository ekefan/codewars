package main

import "fmt"

func ballBounced(height, bounce, window float64) int {
	if height < 0 || bounce < 0 || bounce > 1 || window > height {
		return -1
	}
	ballSeenFromMothersWindow := 0 
	bounceHeight := height 
	for bounceHeight > window {
		ballSeenFromMothersWindow++
		bounceHeight *= bounce
		if bounceHeight > window {
			ballSeenFromMothersWindow++
		}
	}
	return ballSeenFromMothersWindow
}

func main() {
	fmt.Println(ballBounced(3, 0.66, 1.5))
}
