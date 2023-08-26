package main

// Напишите функцию под названием mult,
// которая должна принимать два целочисленных аргумента
// и выводить результат их произведения.

import "fmt"

func mult(a, b int) {
	fmt.Println(a * b)
}

func main() {
	mult(3, 2)
}
