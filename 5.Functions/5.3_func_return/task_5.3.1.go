package main

import "fmt"

// Напишите функцию под названием concat,
// которая принимает две строки в качестве аргументов
// и возвращает строку -
// результат их объединения (конкатенации).

func concat(s1, s2 string) string {
	return s1 + s2
}

func main() {
	fmt.Println(concat("Hel", "lo"))
}
