package main

import (
	"fmt"
	"reflect"
)

/*
Разработать программу,
которая в рантайме способна определить тип переменной: int, string, bool, channel из переменной типа interface{}.

Решение:
Можно использовать переключатели, но переключатели типов работают с определенными типами.
Поэтому можно использовать пакет Reflect, который позволяет во время рантайма определять тип,
метаинформацию о типе и динамически создавать и менять типы в рантайме
*/

func WhatType(value interface{}) {
	if reflect.ValueOf(value).IsNil() {
		fmt.Println("nil")
		return
	}

	switch reflect.ValueOf(value).Kind() {
	case reflect.Int, reflect.Int8, reflect.Int32, reflect.Int64:
		fmt.Println("int")
	case reflect.Float32, reflect.Float64:
		fmt.Println("float64")
	case reflect.String:
		fmt.Println("string")
	case reflect.Bool:
		fmt.Println("bool")
	case reflect.Chan:
		fmt.Println("channel")
	default:
		fmt.Println("unknown")
	}
}

func main() {
	a := 3
	b := make(chan int)
	c := "dgsg"
	d := 44.3
	WhatType(a)
	WhatType(c)
	WhatType(d)
	WhatType(b)
}
