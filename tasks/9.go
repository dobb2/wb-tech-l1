package main

import "fmt"

/*
Разработать конвейер чисел. Даны два канала: в первый пишутся числа (x) из массива, во второй — результат операции x*2,
после чего данные из второго канала должны выводиться в stdout.

Решение:
Используем несколько функций
функция gen записывает данные из массива в канал, а именно записав первое значение она блокируется до тех пор,
пока из канала не прочтут данные и затем дальше продолжает запись.

Вторая функция считывает данные из 1 канала, преобразовывает значение и записывает в следующи канал.
Считывание из первого канала происходит до тех пор, пока его не закроют.

Затем в функции main в цикле происходит считывание из второго канала, который возвращает в stdout полученное значение

*/

func gen(array []int) chan int {
	out := make(chan int)
	go func() {
		for _, value := range array {
			out <- value
		}
		close(out)
	}()
	return out
}

func sq(in chan int) chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out // вернули канал и затем в него постепоенно поступают данные
}

func main() {
	array := []int{2, 4, 6, 8, 10}

	for n := range sq(gen(array)) {
		fmt.Println(n)
	}

}