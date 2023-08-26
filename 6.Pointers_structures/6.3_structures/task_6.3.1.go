package main

// создать структуру Car,
// которая должна иметь три поля: color и brand строчного типа,
// а также year целочисленного

import "fmt"

type Car struct {
	color string
	brand string
	year  int
}

func main() {
	a := Car{"white", "Volvo", 3}

	fmt.Println(a.color)
	fmt.Println(a.brand)
	fmt.Println(a.year)
}
