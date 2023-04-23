package main

import "fmt"

/* Дана переменная int64. Разработать программу которая устанавливает i-й бит в 1 или 0.

Решение:
1) С помощью битового сдвига получим минимальное число с заданным количеством бит(не считая число 0),
которое будем применять для поразрядных операций с заданным числом для изменение значения i-го бита.
пример:
2 бит, тогда минимальное число в двоичной системе 10 - 2 в десятичной
3 бит 1000 - 4 в десятичной

2) в зависимости от того на какое значение хотим поменять i-бит то применяем определенную поразрядную операцию.

 1. Если  заменить i-бит на 1, то поразрядня дизъюнкция возвращает 1 для заданного бита,
если хотя бы у одного числа данный бит равен 1

2. Если  заменить i-бит на 0, то сброс бита операцией (И НЕ) z= x &^ y заменить соответствующий бит на 0 числа z,
если соответствующий бит числа y равен 1
*/

func main() {
	var digit, res int64
	var i, bit int
	fmt.Println("Введите число")
	fmt.Scan(&digit)
	fmt.Println("Введите бит числа(начиная с 0 бита), который необходимо изменить")
	fmt.Scan(&i)
	fmt.Println("На какое значение заданного бита заменить в двочной системе(0 или 1)")
	fmt.Scan(&bit)

	b := int64(1 << i)
	if bit == 1 {
		res = digit | b
	} else {
		res = digit &^ b
	}
	fmt.Println(res)
}
