package main

// вывести Вы гений!, если переменная iq больше 120.

import "fmt"

func main() {
	iq := 121
	s := "Вы гений!"

	if iq >= 120 {
		fmt.Println(s)
	}
}
