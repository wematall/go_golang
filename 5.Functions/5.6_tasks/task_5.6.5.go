package main

// Вам нужно написать функцию oneOrTwo(),
// которая принимает два целых числа и строку.
// Строка может иметь значения one, two или любой другой текст.

// Возвращать из функции вам нужно два значения.
// Если строка равна one, нужно вернуть первое число и саму строку.
// Если строка равна two, нужно вернуть второе число и саму строку.
// Если строка другая - нужно вернуть 0 и саму строку.

import "fmt"

func oneOrTwo(a, b int, s string) (int, string) {
	if s == "one" {
		return a, s
	} else if s == "two" {
		return b, s
	} else {
		return 0, s
	}
}

func main() {
	fmt.Println(oneOrTwo(2, 5, "two"))
	fmt.Println(oneOrTwo(2, 5, "one"))
	fmt.Println(oneOrTwo(2, 5, "abc"))
}
