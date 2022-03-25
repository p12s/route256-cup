package main

import (
	"fmt"
	"strconv"
)

func main() {
	var input int
	fmt.Scanf("%d", &input)

	delta := 0
	for !IsPalindrom(input + delta) {
		delta++
	}

	fmt.Println(delta)
}

func IsPalindrom(number int) bool {
	str := strconv.Itoa(number)
	for i := 0; i < len(str)/2; i++ {
		if str[i] != str[len(str)-i-1] {
			return false
		}
	}
	return true
}
