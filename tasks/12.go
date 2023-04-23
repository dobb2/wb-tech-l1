package main

import "fmt"

/*
Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее собственное множество.

Решение:
обходим массив строк и добавляем i строку в map[string]struct{}
*/

func main() {
	animals := []string{"cat", "cat", "dog", "cat", "tree"}

	setAnimals := make(map[string]struct{})

	for _, value := range animals {
		setAnimals[value] = struct{}{}
	}

	fmt.Println(setAnimals)
}
