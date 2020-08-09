package Reverseint

import (
	"fmt"
	"strconv"

	rev "github.com/shawnwollenberg/goReverseString"
)

//Adjusted Solution found online - Doesnt require converting to string (cleaner)
func Reverse_int(n int) int {
	mult := 1
	if n < 0 {
		mult = -1
		n = -n
	}
	new_int := 0

	for n > 0 {
		//remainder of each spot (right to left), next time through mult by 10 and add new remainder and continue dividing source int by 10 until ==0
		remainder := n % 10
		new_int *= 10
		new_int += remainder
		fmt.Println(remainder, new_int, n)
		n /= 10
	}
	return mult * new_int
}

func ReverseIntConvString(x int) int {
	imult := 1
	str := strconv.Itoa(x)
	if str[:1] == "-" {
		imult = -1
		str = str[1:]
	}
	y, _ := strconv.Atoi(rev.Reversestring(str))
	return imult * y
}
