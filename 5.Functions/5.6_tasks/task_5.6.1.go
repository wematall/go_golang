package main

// Напишите функцию max()
// которая должна принимать в качестве параметров два целых числа
// и возвращать большее из них.

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func main() {
	fmt.Println(max(3, 7))
}
