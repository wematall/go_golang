package main

// Напишите функцию squareSumm(),
// которая должна принимать на вход два целых числа a и b
// и возвращать сумму квадратов чисел от a до b (включительно)

import "fmt"

func squareSumm(a, b int) int {
	result := 0
	for i := a; i <= b; i++ {
		result += i * i
	}

	return result
}

func main() {
	fmt.Println(squareSumm(2, 6))
}
