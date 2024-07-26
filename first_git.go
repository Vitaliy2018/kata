package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	str := scanner.Text()
	first_sumbol := string(str[0])
	//oper := string(str[1])
	second_sumbol := string(str[2])
	for i := 0; i < len(Arab); i++ {
		iStr := strconv.Itoa(i)
		if first_sumbol == iStr {
			for j := 0; j < len(Arab); j++ {
				jStr := strconv.Itoa(j)
				if second_sumbol == jStr {
					for k := 0; k < len(operation); k++ {
						//kStr := strconv.Itoa(k)
						fmt.Print(operation[k])
					}
				}
			}
		}
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
