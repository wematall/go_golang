package main

// вывести Победа, если переменная result равна 3,
// Ничья, если она будет равна 1,
// и Поражение во всех остальных случаях.

import "fmt"

func main() {
	result := 3

	if result == 3 {
		fmt.Println("Победа")
	} else if result == 1 {
		fmt.Println("Ничья")
	} else {
		fmt.Println("Поражение")
	}
}
