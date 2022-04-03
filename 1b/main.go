package main

import (
	"fmt"
	"sort"
)

/*
Склады

Компания-застройщик занимается строительством недвижимости промышленного назначения, а именно – складских помещений.
Каждый квадратный метр площади стоит L рублей. Компания имеет достаточный запас оборотных средств,
чтобы продавать построенные помещения не сразу после завершения строительства, а в тот момент, когда ей это выгодно.
Рынок недвижимости очень динамичный, поэтому цена квадратного метра меняется каждый день.

Аналитики застройщика смогли точно спрогнозировать цену на ближайшие N дней (пронумеруем дни в хронологическом порядке от 0 до N-1).
Теперь требуется определить, в какие из этих дней нужно продавать, чтобы по истечению N дней заработать как можно больше денег.

Известно, что стройка новых площадей происходит с равномерной скоростью S кв. метров в сутки.
А к 0-му дню объем построенной площади составлял S кв. метров, при том что S = 1

Формат выходных данных

Вывести одно число - максимальное кол-во денег, которое может заработать компания-застройщик.

Примеры

Входные данные:
7
81 22 31 44 41 78 98
Выходные данные:
686
Входные данные:
5
81 14 88 2 22
Выходные данные:
308

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
