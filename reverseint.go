package reverseint

import (
	"strconv"

	rev "github.com/shawnwollenberg/goReverseString"
)

// --- Directions
// Given an integer, return an integer that is the reverse
// ordering of numbers.
// --- Examples
//   reverseInt(15) === 51
//   reverseInt(981) === 189
//   reverseInt(500) === 5
//   reverseInt(-15) === -51
//   reverseInt(-90) === -9

//Adjusted Solution found online - Doesnt require converting to string (cleaner)
func reverseInt(n int) int {
	mult := 1
	if n < 0 {
		mult = -1
		n = -n
	}
	newInt := 0

	for n > 0 {
		//remainder of each spot (right to left), next time through mult by 10 and add new remainder and continue dividing source int by 10 until ==0
		remainder := n % 10
		newInt *= 10
		newInt += remainder
		n /= 10
	}
	return mult * newInt
}

func reverseIntConvString(x int) int {
	imult := 1
	str := strconv.Itoa(x)
	if str[:1] == "-" {
		imult = -1
		str = str[1:]
	}
	y, _ := strconv.Atoi(rev.Reversestring(str))
	return imult * y
}
