package main

import (
	"fmt"
	"sync"
	"time"
)

/*
Реализовать структуру-счетчик, которая будет инкрементироваться в конкурентной среде.
По завершению программа должна выводить итоговое значение счетчика.

Решение:
Создаем структуру с мьютексом и переменной счестика, которую отправляем в каждую вызываемую горутину
и для которой вызываем функцию, котораяя инкрементирует счетчик с использованием мьютекса,
чтобы не избежать перезаписи значений счетчика
*/

type Counter struct {
	count int
	mux   sync.Mutex
}

func (c *Counter) Increment() {
	c.mux.Lock()
	defer c.mux.Unlock()
	c.count++
}

func (c Counter) Value() int {
	return c.count
}

func main() {
	cnt := Counter{count: 0}
	wg := sync.WaitGroup{}

	for i := 0; i < 40; i++ {
		wg.Add(1)
		go func(Counter, sync.WaitGroup) {
			defer wg.Done()
			cnt.Increment()
			time.Sleep(3 * time.Second) // do hard work for example

		}(cnt, wg)
	}

	wg.Wait()
	fmt.Println(cnt.Value())
}
