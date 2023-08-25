package main

import "fmt"

// На вход подаются три целых числа.
// Необходимо сосчитать
// и вывести их сумму и произведение на разных строках.

func main() {
	var a, b, c int
	fmt.Scan(&a)
	fmt.Scan(&b)
	fmt.Scan(&c)

	fmt.Println(a + b + c)
	fmt.Println(a * b * c)
}
