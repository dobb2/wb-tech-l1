package main

import (
	"fmt"
	"sync"
)

/*
Дана последовательность чисел: 2,4,6,8,10.
Найти сумму их квадратов(2^2+3^2+4^2….) с использованием конкурентных вычислений.

Решение:
1) Аналогично как в 2, переменную для суммирования и мьютекс, примитив синхронизации,
который изолирует доступ к общей переменной sum, тем самым в переменную добавим значение квадрата числа.

*/

func main() {
	array := [5]int{2, 4, 6, 8, 10}
	ch := make(chan int)
	wg := new(sync.WaitGroup)
	mutex := new(sync.Mutex)
	sum := 0
	for i := range array {
		digit := array[i]
		wg.Add(1)
		go func(int, int, chan int, *sync.WaitGroup) {
			defer wg.Done()
			mutex.Lock()
			sum += digit * digit
			mutex.Unlock()
		}(digit, sum, ch, wg)
	}
	wg.Wait()
	fmt.Println(sum)

}
