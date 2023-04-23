package main

import (
	"fmt"
	"math/big"
	"strconv"
)

/*
Разработать программу, которая перемножает, делит, складывает,
вычитает две числовых переменных a,b, значение которых > 2^20.
*/

const Base int = 1000000000

type BigInt struct {
	digits         []int
	SignIsNegative bool
}

func (b *BigInt) Read(number string) {
	if len(number) == 0 {
		b.SignIsNegative = false
	} else {
		if number[0] == '-' {
			b.SignIsNegative = true
		} else {
			b.SignIsNegative = false
		}
		for i := len(number); i > 0; i -= 9 {
			if i < 9 {
				n, _ := strconv.Atoi(number[:i])
				b.digits = append(b.digits, n)
			} else {
				n, _ := strconv.Atoi(number[i-9 : i])
				b.digits = append(b.digits, n)
			}
		}
		b.RemoveLeadingZeros()
	}
}

func (b *BigInt) RemoveLeadingZeros() {
	for len(b.digits) > 1 && b.digits[len(b.digits)-1] == 0 {
		b.digits = b.digits[:len(b.digits)-1]
	}

	if len(b.digits) == 1 && b.digits[0] == 0 {
		b.SignIsNegative = false
	}
}

func (b BigInt) Write() string {
	result := ""
	if len(b.digits) == 0 {
		result += "0"
	} else {
		if b.SignIsNegative {
			result += "-"
		}
		// следующие числа нам нужно печатать группами по 9 цифр
		// поэтому сохраним текущий символ-заполнитель, а потом восстановим его
		result += fmt.Sprintf("%d", b.digits[len(b.digits)-1])
		for i := len(b.digits) - 2; i >= 0; i-- {
			result += fmt.Sprintf("%09d", b.digits[i])
		}

	}
	return result
}

func (b *BigInt) Addition(a BigInt) {
	carry := 0
	n := Max(len(b.digits), len(a.digits))
	for i := 0; i < n || i < carry; i++ {
		if i == len(b.digits) {
			b.digits = append(b.digits, 0)
		}

		if i < len(a.digits) {
			b.digits[i] += carry + a.digits[i]
		} else {
			b.digits[i] += carry
		}

		if b.digits[i] >= Base {
			carry = 1
		} else {
			carry = 0
		}

		if carry == 1 {
			b.digits[i] -= Base
		}

	}
}

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func (b *BigInt) Subtracting(a BigInt) {
	carry := 0
	a_digit := a.digits
	if b.isLess(a) {
		b.SignIsNegative = true
		a_digit = b.digits
		b.digits = a.digits
	}
	for i := 0; i < len(a_digit); i++ {
		if i < len(b.digits) {
			b.digits[i] -= carry + a_digit[i]
		} else {
			b.digits[i] -= carry
		}
		if b.digits[i] < 0 {
			carry = 1
		} else {
			carry = 0
		}

		if carry == 1 {
			b.digits[i] += Base
		}
	}
	for len(b.digits) > 1 && b.digits[len(b.digits)-1] == 0 {
		b.digits = b.digits[:len(b.digits)-1]
	}
}

func (b BigInt) isLess(a BigInt) bool {
	if len(b.digits) > len(a.digits) {
		return false
	} else if len(b.digits) < len(a.digits) {
		return true
	}

	for i := len(b.digits) - 1; i >= 0; i-- {
		if b.digits[i] != a.digits[i] {
			return b.digits[i] < a.digits[i]
		}
	}
	return false
}

func (b BigInt) isEqual(a BigInt) bool {
	if len(b.digits) != len(a.digits) {
		return false
	}

	for i := len(b.digits) - 1; i >= 0; i-- {
		if b.digits[i] != a.digits[i] {
			return false
		}
	}
	return true
}

func (left *BigInt) Multiplication(right BigInt) {
	if len(right.digits) == 1 {
		left.MultiplicationByShort(right.digits[0])
		return
	}
	c := make([]int, len(left.digits)+len(right.digits))
	for i := 0; i < len(left.digits); i++ {
		carry := 0
		for j := 0; j < len(right.digits) || j < carry; j++ {
			cur := 0
			if j < len(right.digits) {
				cur = c[i+j] + left.digits[i]*right.digits[j] + carry
			} else {
				cur = c[i+j] + left.digits[i]*0 + carry
			}
			c[i+j] = cur % Base
			carry = cur / Base
		}
	}
	for len(c) > 1 && c[len(c)-1] == 0 {
		c = c[:len(c)-1]
	}
	left.digits = c
}

func (left *BigInt) MultiplicationByShort(right int) {
	carry := 0
	for i := 0; i < len(left.digits) || carry != 0; i++ {
		if i == len(left.digits) {
			left.digits = append(left.digits, 0)
		}
		cur := carry + left.digits[i]*right
		left.digits[i] = cur % Base
		carry = cur / Base

	}
	for len(left.digits) > 1 && left.digits[len(left.digits)-1] == 0 {
		left.digits = left.digits[:len(left.digits)-1]
	}
}

func (left *BigInt) Division(right BigInt) {
	leftBig := new(big.Int)
	leftBig.SetString(left.Write(), 10)

	rightBig := new(big.Int)
	rightBig.SetString(right.Write(), 10)

	result := new(big.Int)
	result.Div(leftBig, rightBig)

	res := BigInt{}
	res.Read(result.String())
	left.digits = res.digits
}

func main() {
	// Деление нацело
	division := BigInt{}
	division.Read("3000000000000")
	division2 := BigInt{}
	division2.Read("300000000090")
	division.Division(division2)
	fmt.Println(division.Write())

	// сумма
	add := BigInt{}
	add.Read("4000000000")
	add2 := BigInt{}
	add2.Read("4234400000")
	add.Addition(add2)
	fmt.Println(add.Write())

	// разность
	min := BigInt{}
	min.Read("555555555555")
	min2 := BigInt{}
	min2.Read("455555555555")
	min.Subtracting(min2)
	fmt.Println(min.Write())

	// умножение
	mult := BigInt{}
	mult.Read("555555555555")
	mult2 := BigInt{}
	mult2.Read("455555555555")
	mult.Multiplication(mult2)
	fmt.Println(mult.Write())
}


