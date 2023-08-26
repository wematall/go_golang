package main

// получился правильный оператор switch
// под названием number:

import "fmt"

func main() {
	var number int
	switch number {
	case 0:
		fmt.Println("i equivals 0")
	case 1:
		fmt.Println("i equivals 1")
	case 2:
		fmt.Println("i equivals 2")
	default:
		fmt.Println("i not equivals 0, 1, or 2")
	}
}
