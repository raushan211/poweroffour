package main

import (
	"fmt"
	"math"
)

func main() {
	n := 5
	fmt.Println(isPowerOfFour(n))

}
func isPowerOfFour(n int) bool {
	for i := 0; i < 31; i++ {
		y := float64(i)
		if float64(n) == math.Pow(4, y) {
			return true
		}
	}
	return false
}
