package main

import (
	"fmt"
)

func main() {
	var num1Str, num2Str, operator, fakeOperator string
	fmt.Print("Введите выражение (например, 1 + 2 или V - IV): ")
	fmt.Scanln(&num1Str, &operator, &num2Str, &fakeOperator)

	// Проверяем, что введены только два числа и оператор
	if fakeOperator != "" {
		panic("Паника.")
	}

	var num1num, num2num int
	// Проверяем диапазон чисел (1-10 для арабских, I-X для римских)
	switch num1Str {
	case "1":
		num1num = 1
	case "2":
		num1num = 2
	case "3":
		num1num = 3
	case "4":
		num1num = 4
	case "5":
		num1num = 5
	case "6":
		num1num = 6
	case "7":
		num1num = 7
	case "8":
		num1num = 8
	case "9":
		num1num = 9
	case "10":
		num1num = 10
	}

	switch num2Str {
	case "1":
		num2num = 1
	case "2":
		num2num = 2
	case "3":
		num2num = 3
	case "4":
		num2num = 4
	case "5":
		num2num = 5
	case "6":
		num2num = 6
	case "7":
		num2num = 7
	case "8":
		num2num = 8
	case "9":
		num2num = 9
	case "10":
		num2num = 10
	}

	var isValidArabic bool

	if num1num >= 1 && num2num >= 1 {
		isValidArabic = true //Верный формат чисел: числа должны быть от 1 до 10
	}
	isValidRoman := (num1Str == "I" || num1Str == "II" || num1Str == "III" || num1Str == "IV" || num1Str == "V" || num1Str == "VI" || num1Str == "VII" || num1Str == "VIII" || num1Str == "IX" || num1Str == "X") &&
		(num2Str == "I" || num2Str == "II" || num2Str == "III" || num2Str == "IV" || num2Str == "V" || num2Str == "VI" || num2Str == "VII" || num2Str == "VIII" || num2Str == "IX" || num2Str == "X")
	if isValidArabic == false && isValidRoman == false {
		panic("Паника.")
	}

	var answer int
	if isValidArabic == true {
		switch operator {
		case "+":
			answer = num1num + num2num
		case "-":
			answer = num1num - num2num
		case "*":
			answer = num1num * num2num
		case "/":
			answer = num1num / num2num
		default:
			panic("Паника.")

		}
		fmt.Println(answer)
	}
	var num1Rnum, num2Rnum int
	if isValidRoman == true {
		switch num1Str {
		case "I":
			num1Rnum = 1
		case "II":
			num1Rnum = 2
		case "III":
			num1Rnum = 3
		case "IV":
			num1Rnum = 4
		case "V":
			num1Rnum = 5
		case "VI":
			num1Rnum = 6
		case "VII":
			num1Rnum = 7
		case "VIII":
			num1Rnum = 8
		case "IX":
			num1Rnum = 9
		case "X":
			num1Rnum = 10

		}
		switch num2Str {
		case "I":
			num2Rnum = 1
		case "II":
			num2Rnum = 2
		case "III":
			num2Rnum = 3
		case "IV":
			num2Rnum = 4
		case "V":
			num2Rnum = 5
		case "VI":
			num2Rnum = 6
		case "VII":
			num2Rnum = 7
		case "VIII":
			num2Rnum = 8
		case "IX":
			num2Rnum = 9
		case "X":
			num2Rnum = 10
		}

		switch operator {
		case "+":
			answer = num1Rnum + num2Rnum
		case "-":
			answer = num1Rnum - num2Rnum
		case "*":
			answer = num1Rnum * num2Rnum
		case "/":
			answer = num1Rnum / num2Rnum
		default:
			panic("Паника.")

		}

		if answer < 1 {
			panic("Паника.")

		}

		units := []string{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"}
		tens := []string{"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"}
		hundreds := []string{"", "C"}
		new_answer := fmt.Sprintf("%s%s%s", hundreds[answer/100], tens[(answer%100)/10], units[answer%10])
		fmt.Println(new_answer)
	}
}
