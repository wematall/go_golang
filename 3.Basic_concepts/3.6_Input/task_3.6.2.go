package main

// Надо считать со стандартного входа 4 числа
// и вывести их в одну строку через пробел.

import "fmt"

func main() {
	var a, b, c, d int

	fmt.Scan(&a, &b, &c, &d)
	fmt.Println(a, b, c, d)
}
