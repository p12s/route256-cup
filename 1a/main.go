package main

import (
	"fmt"
)

/*
Маски

Никита собирается открывать свой пункт выдачи заказов. Для обеспечения безопасности посетителей Никита решил обеспечить ПВЗ
одноразовыми медицинскими масками, которые будут поставляться каждый месяц в количестве n штук.
Местные поставщики продают маски по цене 0.55 рублей за одну штуку, но можно сэкономить,
купив упаковку из 24 масок за 11 рублей или коробку из 16 упаковок за 160 рублей.
Помогите Никите, определив, сколько коробок, упаковок и отдельных масок он должен купить, чтобы заплатить как можно меньше денег.
Примечание: Никита готов купить больше масок, чем ему нужно, если это позволит сэкономить.

Формат выходных данных

Требуется вывести три числа в одну строку через пробел - количество отдельных масок, упаковок и коробок, которые надо купить.
Примеры

Входные данные:
385
Выходные данные:
1 0 1
Входные данные:
23
Выходные данные:
0 1 0
Входные данные:
27
Выходные данные:
3 1 0

0.55 - 1 шт
11 - 24 шт - упаковка
160 - 16 упаковок - коробка
"количество отдельных масок" - "упаковок" - "коробок"

*/
func main() {
	var input string
	fmt.Scanf("%s", &input)

	arr := make([]int, 0)

	for _, item := range arr {
		fmt.Println(item)
	}
}

func Some(input int) []int {
	arr := make([]int, 0)

	return arr
}
