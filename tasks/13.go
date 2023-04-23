package main

import "fmt"

/*
Поменять местами два числа без создания временной переменной.

Решение:
поиграться со сложением и вычитанием
*/

func main() {
	a := 5
	b := 10

	a += b
	b = a - b
	a = a - b

	fmt.Printf("a : %d \n", a)
	fmt.Printf("b : %d \n", b)
}
