package main

import "fmt"

/*
Реализовать пересечение двух неупорядоченных множеств.

Решение: д
ля реализации множества используем map, где значением будет пустая структура struct{}(или bool=true),
но struct{} занимает 0 байт, а ключом заданный тип множества

Пересечение множеств - это множество содержащие элементы, которые содержатся в обоих множествах.
Достаточно обойти одно из множеств(map) и проверять есть ли данные элемент во втором множестве,
если элемент есть, то добавляем его в результирующее множество

*/

func main() {
	set1 := make(map[int]struct{})
	for i := 0; i < 7; i++ {
		set1[i] = struct{}{}
	}

	fmt.Printf("Множество 1: ")
	for k := range set1 {
		fmt.Printf("%d ", k)
	}
	fmt.Printf("\n")

	set2 := make(map[int]struct{})
	for i := 4; i < 15; i++ {
		set2[i] = struct{}{}
	}

	fmt.Printf("Множество 2: ")
	for k := range set2 {
		fmt.Printf("%d ", k)
	}
	fmt.Printf("\n")

	intersecSets := make(map[int]struct{})
	for k := range set1 {
		if _, ok := set2[k]; ok {
			intersecSets[k] = struct{}{}
		}
	}

	fmt.Printf("Пересечение двух множеств: ")
	for k := range intersecSets {
		fmt.Printf("%d ", k)
	}
	fmt.Printf("\n")

}
