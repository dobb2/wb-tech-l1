package main

import "fmt"

/*
Разработать программу, которая переворачивает подаваемую на ход строку
(например: «главрыба — абырвалг»).
Символы могут быть unicode.

Решение:
так как строка это неизменяемая последовательность байт и при этом символы могут быть unicode,
то будем использовать []rune
Тип []rune можно изменять и 1 символ состоит из 1 кода(int32), в отличие от []byte,
где 1 символ может занимать более 1 байта.

Преобразуем подаваемую строку в руну и обойдем руну до п/2 и поменяем местами первый символ с последним,
второй с предпоследним и так до n/2

Чтобы вывести измененную строк, необходимо измененный тип []rune преобразовать обратно в строку
*/

func main() {
	var inputString string
	fmt.Scan(&inputString)
	inputRune := []rune(inputString)
	n := len(inputRune)
	for i := 0; i < n/2; i++ {
		inputRune[i], inputRune[n-i-1] = inputRune[n-i-1], inputRune[i]
	}
	outString := string(inputRune)
	fmt.Println(outString)

}
