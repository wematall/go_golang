package main

// Напишите функцию под названием myStr,
// которая принимает две строки в качестве аргументов
// и возвращает эти же две строки.

import "fmt"

func myStr(s1, s2 string) (string, string) {
	return s1, s2
}

func main() {
	var s1, s2 string
	s1, s2 = myStr("Ivan", "Sidor")
	fmt.Println(s1, s2)
}
