package main

/*
Разработать программу, которая будет последовательно отправлять значения в канал, а с другой стороны канала — читать.
По истечению N секунд программа должна завершаться.

Используем контекст, который отменяется по истечении времени, и как отправим его в горутину,
которая будет отправлять данные в канал, как только получим уведомление от контекста,
горутина закроет канал и получатель таким образом прекратит чтение канала и программа завершится.

Значение параметра N зададим с помощью флага при компиляции программы
*/

import (
	"context"
	"fmt"
	"time"
)

const (
	N = 3 * time.Second
)

func main() {
	ctx, _ := context.WithTimeout(context.Background(), N)
	ch1 := make(chan string)

	go func(context.Context, chan string) {
		var value string
		for {
			select {
			case <-ctx.Done():
				close(ch1) // no statement
				fmt.Println("Время истекло")
				return
			default:
				fmt.Scan(&value)
				ch1 <- value
			}
		}
	}(ctx, ch1)

	for value := range ch1 {
		fmt.Println("Get data ", value)
	}
	fmt.Println("Завершение программы")
}
