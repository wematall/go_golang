package main

// создать структуру Student,
// которая должна иметь два поля: name и university строчного типа.

import "fmt"

type Student struct {
	name       string
	university string
}

func main() {
	sr := Student{"Ivan", "FMT"}
	fmt.Println(sr.name, sr.university)
}
