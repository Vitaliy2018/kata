package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	sort     = "Выдача паники, так как используются одновременно разные системы счисления."
	error_op = "Выдача паники, так как формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *)."
	rim_otr  = "Выдача паники, так как в римской системе нет отрицательных чисел."
	dan      = "Выдача паники, так как не правильные данные"
)

var (
	Arab      = [11]string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "10"}
	operation = [4]string{"+", "-", "/", "*"}
	Rim       = [101]string{"0", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX",
		"X", "XI", "XII", "XIII", "XIV", "XV", "XVI", "XVII", "XVIII", "XIX",
		"XX", "XXI", "XXII", "XXIII", "XXIV", "XXV", "XXVI", "XXVII", "XXVIII", "XXIX",
		"XXX", "XXXI", "XXXII", "XXXIII", "XXXIV", "XXXV", "XXXVI", "XXXVII", "XXXVIII", "XXXIX",
		"XL", "XLI", "XLII", "XLIII", "XLIV", "XLV", "XLVI", "XLVII", "XLVIII", "XLIX",
		"L", "LI", "LII", "LIII", "LIV", "LV", "LVI", "LVII", "LVIII", "LIX",
		"LX", "LXI", "LXII", "LXIII", "LXIV", "LXV", "LXVI", "LXVII", "LXVIII", "LXIX",
		"LXX", "LXXI", "LXXII", "LXXIII", "LXXIV", "LXXV", "LXXVI", "LXXVII", "LXXVIII", "LXXIX",
		"LXXX", "LXXXI", "LXXXII", "LXXXIII", "LXXXIV", "LXXXV", "LXXXVI", "LXXXVII", "LXXXVIII", "LXXXIX",
		"XC", "XCI", "XCII", "XCIII", "XCIV", "XCV", "XCVI", "XCVII", "XCVIII", "XCIX", "C"}
)

func main() {
	first_number := 111
	oper_int := 5
	second_number := 111
	rim_count := 0
	first_rim_number := 0
	second_rim_number := 0
	second_arab_number := 0
	first_arab_number := 0
	arab_count := 0
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	str := scanner.Text()
	words := strings.Split(str, " ")
	if len(words) != 3 {
		fmt.Print(error_op)
	} else {
		first_sumbol := string(words[0])
		oper := string(words[1])
		second_sumbol := string(words[2])
		for i := 0; i < len(Arab); i++ {
			iStr := strconv.Itoa(i)
			if first_sumbol == iStr {
				arab_count++
				first_arab_number = i
			}
		}
		for i := 0; i < len(Arab); i++ {
			iStr := strconv.Itoa(i)
			if second_sumbol == iStr {
				arab_count++
				second_arab_number = i
			}
		}
		for i := 0; i < 11; i++ {
			if first_sumbol == Rim[i] {
				first_rim_number = i
				rim_count++
			}
		}
		for i := 0; i < 11; i++ {
			if second_sumbol == Rim[i] {
				second_rim_number = i
				rim_count++
			}
		}
		for i := 0; i < len(operation); i++ {
			if operation[i] == oper {
				oper_int = i
			}
		}
		if oper_int == 5 {
			fmt.Print(error_op)
		}
	}
	if rim_count == 2 || arab_count == 2 {
		if rim_count == 2 {
			first_number = first_rim_number
			second_number = second_rim_number
			switch oper_int {
			case 0:
				fmt.Print(Rim[sum(first_number, second_number)])
			case 1:
				if sub(first_number, second_number) <= 0 {
					fmt.Print(rim_otr)
				} else {
					fmt.Print(Rim[sub(first_number, second_number)])
				}
			case 2:
				if div(first_number, second_number) <= 0 {
					fmt.Print(rim_otr)
				} else {
					fmt.Print(Rim[div(first_number, second_number)])
				}
			case 3:
				fmt.Print(Rim[mult(first_number, second_number)])
			}

		} else {
			first_number = first_arab_number
			second_number = second_arab_number
			switch oper_int {
			case 0:
				fmt.Print(sum(first_number, second_number))
			case 1:
				fmt.Print(sub(first_number, second_number))
			case 2:
				fmt.Print(div(first_number, second_number))
			case 3:
				fmt.Print(mult(first_number, second_number))
			}
		}
	}
	if arab_count == 1 && rim_count == 1 {
		fmt.Print(sort)
	}
	if first_number >= 0 && first_number <= 10 && second_number >= 0 && second_number <= 10 {
	} else {
		fmt.Print(dan)
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
