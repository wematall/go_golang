package main

// Результаты матчей
// Вы создаете программу для анализа результатов спортивных матчей
// и подсчета очков заданной команды.
// Результаты матчей хранятся в массиве results.
// Каждый матч имеет один из следующих результатов:
// "w" - выиграл
// "l" - проиграл
// "d" - ничья

// Победа добавляет три очка,
// ничья - одно очко,
// а проигранный матч не добавляет очков.

// Ваша программа должна принять на вход результат
// последнего матча и добавить его в массив результатов.
// После этого необходимо вычислить и вывести общее количество очков,
// набранных командой по результатам матчей из массива results.

import "fmt"

func main() {
	var a string
	fmt.Scan(&a)

	results := []string{"w", "l", "w", "d", "w", "l", "l", "l", "d", "d", "w", "l", "w", "d"}
	results = append(results, a)
	sum := 0
	for _, v := range results {
		if v == "w" {
			sum += 3
		} else if v == "d" {
			sum += 1
		} else {
			sum += 0
		}
	}
	fmt.Println(sum)

}
