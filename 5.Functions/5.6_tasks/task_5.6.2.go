package main

// Напишите функцию calc()
// которая должна принимать на вход одно целое число,
// а возвращать два значения -
// число умноженное на два и число возведенное в квадрат.

import "fmt"

func calc(a int) (int, int) {
	return a * 2, a * a
}

func main() {
	x, y := calc(3)
	fmt.Println(x)
	fmt.Println(y)
}
