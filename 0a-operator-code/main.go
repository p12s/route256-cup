package main

import (
	"fmt"
)

func main() {
	var phoneNumber string
	fmt.Scanf("%s", &phoneNumber)
	if len(phoneNumber) >= 4 {
		fmt.Println(phoneNumber[1:4])
	}
}
