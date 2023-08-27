package main

import "fmt"

// Есть массив nums.
// Вам нужно написать цикл для перебора его значений
// с использованием range.
// В цикле нужно посчитать сумму всех элементов массива.
// Для этого объявлена переменная sum.

func main() {
	var nums = make([]int, 5)
	sum := 0

	nums[0] = 3
	nums[1] = 4
	nums[2] = 2
	nums[3] = 1
	nums[4] = 5

	for _, v := range nums {
		sum += v
	}
	fmt.Println(sum)
}
