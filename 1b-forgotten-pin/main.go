package main

import (
	"fmt"
	"sort"
)

func main() {
	var input string
	fmt.Scanf("%s", &input)

	arr := make([]string, 0)

	arr = append(arr, fmt.Sprintf("%s%s%s", string(input[0]), string(input[1]), string(input[2])))
	arr = append(arr, fmt.Sprintf("%s%s%s", string(input[0]), string(input[2]), string(input[1])))
	arr = append(arr, fmt.Sprintf("%s%s%s", string(input[1]), string(input[0]), string(input[2])))
	arr = append(arr, fmt.Sprintf("%s%s%s", string(input[1]), string(input[2]), string(input[0])))
	arr = append(arr, fmt.Sprintf("%s%s%s", string(input[2]), string(input[0]), string(input[1])))
	arr = append(arr, fmt.Sprintf("%s%s%s", string(input[2]), string(input[1]), string(input[0])))

	sort.Strings(arr)

	for _, item := range arr {
		fmt.Println(item)
	}
}

func Some(input string) []string {
	arr := make([]string, 0)

	arr = append(arr, fmt.Sprintf("%s%s%s", string(input[0]), string(input[1]), string(input[2])))
	arr = append(arr, fmt.Sprintf("%s%s%s", string(input[0]), string(input[2]), string(input[1])))
	arr = append(arr, fmt.Sprintf("%s%s%s", string(input[1]), string(input[0]), string(input[2])))
	arr = append(arr, fmt.Sprintf("%s%s%s", string(input[1]), string(input[2]), string(input[0])))
	arr = append(arr, fmt.Sprintf("%s%s%s", string(input[2]), string(input[0]), string(input[1])))
	arr = append(arr, fmt.Sprintf("%s%s%s", string(input[2]), string(input[1]), string(input[0])))

	sort.Strings(arr)
	return arr
}
