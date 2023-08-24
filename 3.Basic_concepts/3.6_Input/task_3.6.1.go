package main

// программа считывала два целых числа со входа
// и выводила их сумму

import "fmt"

func main() {
	var a, b int

	fmt.Scan(&a)
	fmt.Scan(&b)

	fmt.Println(a + b)

}
