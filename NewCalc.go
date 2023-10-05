package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Scan1() string {
	in := bufio.NewScanner(os.Stdin)
	in.Scan()
	if err := in.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Ошибка ввода:", err)
	}
	return strings.ToUpper(in.Text())
}

func main() {
	fmt.Print("Введите выражение: ")
	digits := strings.Fields(Scan1())
	switch {
	case len(digits) == 1:
		fmt.Println("Ошибка, так как строка не является математической операцией")
		return
	case len(digits) == 2:
		fmt.Println("Ошибка, так как строка не является математической операцией")
		return
	case len(digits) > 3:
		fmt.Println("Ошибка, так как формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *)")
		return
	}
	var felix string = digits[1]
	var x, y int
	var lola bool
	switch {
	case romanToInt(digits[0]) != 0 && romanToInt(digits[2]) != 0:
		x = romanToInt(digits[0])
		y = romanToInt(digits[2])
		lola = true
	case (romanToInt(digits[0]) == 0 && romanToInt(digits[2]) != 0) || (romanToInt(digits[0]) != 0 && romanToInt(digits[2]) == 0):
		fmt.Println("Ошибка, так как используются одновременно разные системы счисления")
		return
	case romanToInt(digits[0]) == 0 && romanToInt(digits[2]) == 0:
		x, _ = strconv.Atoi(digits[0])
		y, _ = strconv.Atoi(digits[2])
	}

	if x > 10 || x < 1 || y > 10 || y < 1 {
		fmt.Println("Превышено число максимума по заданию")
		return
	}
	res, err := Calc(x, felix, y)
	if err == nil {
		if lola {
			if res < 0 || res == 0 {
				fmt.Println("Ошибка, так как в Римской системе нет нуля и отрицательных чисел!")
				return
			}
			fmt.Println(Roman(res))
		} else {
			fmt.Println(res)
		}
	} else {
		fmt.Println(err)
	}
}

func Calc(x int, felix string, y int) (int, error) {
	switch felix {
	case "*":
		return x * y, nil
	case "-":
		return x - y, nil
	case "+":
		return x + y, nil
	case "/":
		return x / y, nil
	default:
		return 0, errors.New("Недопустимая операция!")
	}
}

// Функция преобразования Арабских чисел в Римские найдена в Интернете
func Roman(number int) string {
	conversions := []struct {
		value int
		digit string
	}{
		{1000, "M"},
		{900, "CM"},
		{500, "D"},
		{400, "CD"},
		{100, "C"},
		{90, "XC"},
		{50, "L"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}

	roman := ""
	for _, conversion := range conversions {
		for number >= conversion.value {
			roman += conversion.digit
			number -= conversion.value
		}
	}
	return roman
}

// Функция преобразования Римских чисел в арабские найдена в Интернете
func romanToInt(s string) int {
	rMap := map[string]int{"I": 1, "V": 5, "X": 10, "L": 50, "C": 100, "D": 500, "M": 1000}
	result := 0
	for k := range s {
		if k < len(s)-1 && rMap[s[k:k+1]] < rMap[s[k+1:k+2]] {
			result -= rMap[s[k:k+1]]
		} else {
			result += rMap[s[k:k+1]]
		}
	}
	return result
}
