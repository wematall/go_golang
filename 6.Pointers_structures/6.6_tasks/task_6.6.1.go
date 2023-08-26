package main

// написать фунцию changeStrings,
// которая должна принимать указатели на две строки
// и затем менять местами их значения.

import "fmt"

func changeStrings(s1 *string, s2 *string) {
	*s1, *s2 = *s2, *s1
}

func main() {
	s1 := "World"
	s2 := "Hello"
	fmt.Println(s1, s2)
	changeStrings(&s1, &s2)
	fmt.Println(s1, s2)
}
