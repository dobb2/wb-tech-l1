package main

import (
	"fmt"
	"strings"
)

/*
Разработать программу, которая проверяет, что все символы в строке уникальные
(true — если уникальные, false etc). Функция проверки должна быть регистронезависимой.

Например:
abcd — true
abCdefAaf — false
	aabcd — false

Решение:
1) c помощью функции strings.ToLower(str) все символы в нижнем регистре
обходим посимвольно строку и добавляем значение литерала руны в мапу,
если значение в мапе уже есть, значит символ неуникальный и возвращаем false
*/

func CountLetters(str string) bool {
	mapLetter := make(map[int]struct{})
	lowerRune := []rune(strings.ToLower(str))
	for _, value := range lowerRune {
		_, ok := mapLetter[int(value)]
		if ok {
			return false
		} else {
			mapLetter[int(value)] = struct{}{}
		}
	}
	return true
}

func main() {
	fmt.Println(CountLetters("abcd"))
	fmt.Println(CountLetters("abCdefAaf"))
	fmt.Println(CountLetters("aabcd"))
}
