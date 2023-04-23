package main

import "fmt"

/*
1) Дана структура Human (с произвольным набором полей и методов).
Реализовать встраивание методов в структуре Action от родительской структуры Human (аналог наследования).


Решение:
в Go вместо наследования используется композиция. Эмбеддинг, то есть встраивание — это реализация композиции в Go.
Структуры могут в себя включать другие структуры и типы.

Все поля и методы Human будут переданы в структуру Action, как если бы она сама их содержала.
Это позволяет переиспользовать код сложных структур, встраивая одни в другие.

При этом Action будет включать в себя Human. Это похоже на наследование, но есть существенные отличия.
Структура Action не является Human, можно сказать, что она включает её в себя.

Объект Action не может быть приведён к типу Human с помощью приведения типов, то есть
имеется в виду классическое приведение типов в ООП языках,
где экземпляр производного класса может выступать в качестве экземпляра базового класса.
Условно говоря такой конструкции, как human := Action(human) для встроенных типов в Go нет.


*/

type Human struct {
	Name string
	age  int //
}

func (h Human) Talk() {
	fmt.Println("I am ", h.Name)
}

func (h Human) talk() {
	fmt.Println("Данный метод нельзя встроить в другую структуры из другого пакета(не main). Из этого пакета вызвать можем")
}

func (h Human) String() string {
	return fmt.Sprintf("Имя: %s, возвраст: %d", h.Name, h.age)
}

func (h Human) Do() {
	fmt.Println("Данный метод структуры Human")
}

type Action struct {
	typeAction string
	Human
}

func (a Action) String() string {
	return fmt.Sprintf("%s любит %s", a.Name, a.typeAction)
}

func (a Action) Do() {
	fmt.Println("Данный метод затеняет метод встроенной структуры Human")
}

func main() {
	h1 := Human{Name: "Вася", age: 23}
	h1.Talk()
	h1.Do()
	a := Action{typeAction: "бег", Human: h1}
	fmt.Println(a)
	fmt.Println(a.Human) // указав явно тип, можем обратиться к методу String
	a.Talk()             // cтруктура Action имееть доступ к методу структуры Action
	a.Do()
	a.Human.Do() // указав явно тип получим метод структуры Human

}