package main

import (
	"fmt"
	"sort"
)

/*
Картины

Максим решил заняться продажей картин на NFT-площадках, а для этого нужно придумать что-то свое и оригинальное. Максиму очень нравится, как выглядят цифры на черном фоне. Нужно написать программу, которая будет рисовать квадрат, состоящий из целых чисел и выводить информацию о его размере.

Формат выходных данных

В первой строке выводится сообщение “Квадрат со стороной N”.
В следующих строках требуется вывести квадрат, сформированный из набора чисел или символов, длиной стороны является длина введенного массива.
Примеры

Входные данные:
1 2 3 4 5
Выходные данные:
Квадрат со стороной 5
1 2 3 4 5
1 2 3 4 5
1 2 3 4 5
1 2 3 4 5
1 2 3 4 5
*/
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
