package main

import "fmt"

/*
Удалить i-ый элемент из слайса.
*/

func main() {
	s := []int{1, 2, 3, 4, 5}
	i := 2

	if len(s) != 0 && i < len(s) { // защищаемся от паники
		s = append(s[:i], s[i+1:]...)
	}
	fmt.Println(s) // [1 2 4 5]
}
