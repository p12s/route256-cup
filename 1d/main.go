package main

import (
	"fmt"
	"sort"
)

/*
Продажи

В базе данных есть таблица Invoice следующего вида:

Напишите запрос, который покажет общие продажи по всем странам, отсортированным в порядке возрастания. На выходе в первой колонке должно быть название страны, а во второй показатель общих продаж.

Примечание
Для решения задачи используется база данных Chinook Database в формате Sqlite - см. файл  Chinook_Sqlite.sqlite в разделе Данные под описанием задачи.

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
