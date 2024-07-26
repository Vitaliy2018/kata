package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	sort   = "Цифры не одной системы"
	biggir = "Входные данные слишком большие"
)

var (
	Arab      = [11]string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "10"}
	operation = [4]string{"+", "-", "/", "*"}
	Rim       = [82]string{"0", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX",
		"X", "XI", "XII", "XIII", "XIV", "XV", "XVI", "XVII", "XVIII", "XIX",
		"XX", "XXI", "XXII", "XXIII", "XXIV", "XXV", "XXVI", "XXVII", "XXVIII", "XXIX",
		"XXX", "XXXI", "XXXII", "XXXIII", "XXXIV", "XXXV", "XXXVI", "XXXVII", "XXXVIII", "XXXIX",
		"XL", "XLI", "XLII", "XLIII", "XLIV", "XLV", "XLVI", "XLVII", "XLVIII", "XLIX",
		"L", "LI", "LII", "LIII", "LIV", "LV", "LVI", "LVII", "LVIII", "LIX",
		"LX", "LXI", "LXII", "LXIII", "LXIV", "LXV", "LXVI", "LXVII", "LXVIII", "LXIX",
		"LXX", "LXXI", "LXXII", "LXXIII", "LXXIV", "LXXV", "LXXVI", "LXXVII", "LXXVIII", "LXXIX",
		"LXXX", "LXXXI"}
)

func main() {
	for i := 0; i < len(Rim); i++ {
		fmt.Println(i, " ", Rim[i])
	}
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	str := scanner.Text()
	fmt.Print(str)
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
