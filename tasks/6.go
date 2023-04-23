package main

import (
	"context"
	"fmt"
	"time"
)

/*
6) Реализовать все возможные способы остановки выполнения горутины.

1) отправить в горутину сигнальный канал, как только из вне придет значение в него,
тогда в select statement будет выполнен блок case <-signal

2) отарвить в функцию горутины контекст, можно вызвать функцию вручную cancel(), либо использовать
ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second), который отменяет по истечении заданного времени

3) остановить работу горутины изнутри с помощью таймера, который по истечении времени сработает и в функции ChannelStop
по истечении заданного времени завершить работу горутина

4) можно завершить main, тогда и горутины тоже завершат свою работу, так делать обычно не стоит
*/

func ChannelStop(signal chan struct{}) {
	sum := 0
	defer fmt.Println("goroutine with channel work done")
	timer := time.NewTimer(8 * time.Second)
	for {
		select {
		case <-signal:
			return
		case <-timer.C:
			return
		default:
			sum++
			time.Sleep(1 * time.Second)
		}
	}
}

func CtxStop(ctx context.Context) {
	doworkctx := 0
	defer fmt.Println("goroutine with context work done")
	for {
		select {
		case <-ctx.Done():
			return
		default:
			doworkctx++
			time.Sleep(2 * time.Second)
		}
	}
}

func main() {
	ch := make(chan struct{})
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	time.AfterFunc(7*time.Second, cancel)

	go ChannelStop(ch)
	go CtxStop(ctx)

	// убить main
	//time.Sleep(3 * time.Second)
	//return

	time.Sleep(5 * time.Second)
	ch <- struct{}{}

	time.Sleep(10 * time.Second)
	fmt.Println("main is work")
}
