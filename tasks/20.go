package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/*
Разработать программу, которая переворачивает слова в строке.
Пример: «snow dog sun — sun dog snow».

так как fmt.Scan() считывает строку до первого пробела, полностью считать строку за раз с помощью bufio.NewReader()
Полагаем, что слова разделены только пробелом

Полученную строку с помощью функции Split разделяем на элементы массива,
разделителем в строке которой является пробел.

Полученные массив слов переворачиваем и затем применяем функцию split,
которая полученный массив строк соединяет в строку, с заданным разделителем.

*/

func main() {
	text, _ := bufio.NewReader(os.Stdin).ReadString('\n') // считываем вводимую строку до переноса строки
	text = strings.Trim(text, "\n")                       // убираем символ переноса строки из строки
	arrayWords := strings.Split(text, " ")

	revArrayWords := []string{}
	for i := range arrayWords {
		revArrayWords = append(revArrayWords, arrayWords[len(arrayWords)-1-i])
	}
	resultString := strings.Join(revArrayWords, " ")
	fmt.Println(resultString)
}
