package main

// Напишите функцию isEven(),
// которая принимает в качестве аргумента одно целое число
// и возвращает true если оно четное и false в ином случае.

import "fmt"

func isEven(a int) bool {
	if a%2 == 0 {
		return true
	} else {
		return false
	}
}

func main() {
	fmt.Println(isEven(4))
	fmt.Println(isEven(7))
}
