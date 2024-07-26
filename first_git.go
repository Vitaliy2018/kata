package main

import (
	"fmt"
	"strconv"
)

const (
	sort   = "Цифры не одной системы"
	biggir = "Входные данные слишком большие"
)

func main() {
	var (
		first_number_str, second_number_str string
		chek                                int
	)
	fmt.Scan(&first_number_str, &second_number_str)
	for i := 0; i <= 10; i++ {
		if first_number_str == strconv.Itoa(i) || second_number_str == strconv.Itoa(i) {
			if first_number_str == second_number_str {
				chek = 2
			} else {
				chek++
			}
		}
	}
	if chek == 2 {
		fmt.Print("good")
	} else {
		fmt.Print("bad")
	}
}

func sum(first_number int, second_number int) int { // Функиця суммы
	return first_number + second_number
}
func sub(first_number int, second_number int) int { // Функция разности
	return first_number - second_number
}
func div(first_number int, second_number int) int { // Функция деление
	return first_number / second_number
}
func mult(first_number int, second_number int) int { // Функция умножение
	return first_number * second_number
}
