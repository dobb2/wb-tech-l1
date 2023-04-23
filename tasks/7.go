package main

/*
7) Реализовать конкурентную запись данных в map.

Решение: чтобы несколько потоков смогли корректон добавить данные в общий ресурс, нужно разграничивать доступ,
чтобы не вышла запись в один ресурс одновременно нескольких горутин и одно из значений затерялось
(пример суммирование значений в общий ресурс)

Поэтому используем sync.Mutex, который на уровне кода представляет блокирование доступа к общему разделяемому ресурсу.
у мьютекса вызывается метод Lock(), а для разблокировки доступа - метод Unlock().

В качестве примера создадим структуру с map и мьютексом. Перед каждой записью в map вызываем Lock(), а после записи Unlock()

выведем map после записи в нее данные, перед этим обязательно дождемся окончание работы всех горутин(sync.sync.WaitGroup{}),
которые запишут данные и затем прочтем данные из map.
Чтение из общего ресурса необходимо тоже ограничивать Lock() и Unlock()
*/

import (
	"fmt"
	"sync"
)

type Storage struct {
	mux   sync.Mutex
	cache map[int]int
}

func Create() Storage {
	return Storage{
		cache: map[int]int{},
	}
}

func (s Storage) Write(wg *sync.WaitGroup, key int, value int) {
	s.mux.Lock()
	defer s.mux.Unlock()
	s.cache[key] = value
	wg.Done()
}

func (s Storage) Read() {
	s.mux.Lock()
	defer s.mux.Unlock()
	fmt.Println(s.cache)
}

func main() {
	metrics := Create()
	wg := sync.WaitGroup{}
	n := 100
	wg.Add(n)
	for i := 0; i < n; i++ {
		keyValue := i
		metrics.Write(&wg, keyValue, keyValue)
	}

	wg.Wait()
	metrics.Read()
	//fmt.Println(metrics.cache)
}
