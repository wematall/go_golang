package main

// На стандартный вход подается три целых числа,
// в одной строке через пробел.
// Вам нужно их считать и вывести их сумму.

import "fmt"

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)
	fmt.Println(a + b + c)
}
