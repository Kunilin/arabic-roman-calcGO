package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

var a, b, oper string

var str string

func toRoman(data int) string {
	ones := [101]string{"O", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX", "X", "XI", "XII", "XIII", "XIV", "XV", "XVI", "XVII", "XVIII", "XIX", "XX",
		"XXI", "XXII", "XXIII", "XXIV", "XXV", "XXVI", "XXVII", "XXVIII", "XXIX", "XXX", "XXXI", "XXXII", "XXXIII", "XXXIV", "XXXV", "XXXVI", "XXXVII", "XXXVIII", "XXXIX", "XL",
		"XLI", "XLII", "XLIII", "XLIV", "XLV", "XLVI", "XLVII", "XLVIII", "XLIX", "L", "LI", "LII", "LIII", "LIV", "LV", "LVI", "LVII", "LVIII", "LIX", "LX",
		"LXI", "LXII", "LXIII", "LXIV", "LXV", "LXVI", "LXVII", "LXVIII", "LXIX", "LXX",
		"LXXI", "LXXII", "LXXIII", "LXXIV", "LXXV", "LXXVI", "LXXVII", "LXXVIII", "LXXIX", "LXXX",
		"LXXXI", "LXXXII", "LXXXIII", "LXXXIV", "LXXXV", "LXXXVI", "LXXXVII", "LXXXVIII", "LXXXIX", "XC",
		"XCI", "XCII", "XCIII", "XCIV", "XCV", "XCVI", "XCVII", "XCVIII", "XCIX", "C"}
	checking := ones[data]
	switch {
	case len(checking) < 1:
		fmt.Print("Roman numbers is out of range!")
		os.Exit(0)
	default:
		return ones[data]
	}
	return ""
}

func toArabic(data string) int {
	switch data {
	case "I":
		return 1
	case "II":
		return 2
	case "III":
		return 3
	case "IV":
		return 4
	case "V":
		return 5
	case "VI":
		return 6
	case "VII":
		return 7
	case "VIII":
		return 8
	case "IX":
		return 9
	case "X":
		return 10
	default:
		fmt.Print("Wrong data format or integer number must be in the range 1-10 included")
		os.Exit(0)
	}
	return 0
}

func calculation(num1 int, num2 int, char string) int {

	var result int

	switch char {
	case "+":
		result := num1 + num2
		return result
	case "-":
		result := num1 - num2
		return result
	case "*":
		result := num1 * num2
		return result
	case "/":
		if num2 == 0 {
			fmt.Print("Division by zero error")
			os.Exit(0)
		} else {
			result := float64(num1 / num2)
			return int(math.Floor(result))
		}
	default:
		fmt.Print("Invalid operator.")
	}
	return result
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter your statement using a spase after each element of statement: ")
	scanner.Scan()
	str := scanner.Text()
	split := strings.Split(str, " ")
	if len(split) > 3 || len(split) < 3 {
		fmt.Print("The format of the mathematical operation does not satisfy the task!")
		os.Exit(0)
	}
	a, b := split[0], split[2]
	oper := split[1]
	if _, err := strconv.Atoi(a); err == nil {
		number1, _ := strconv.Atoi(a)
		if _, err := strconv.Atoi(b); err == nil {
			number2, _ := strconv.Atoi(b)
			if number1 > 10 || number2 > 10 || number1 < 1 || number2 < 1 {
				fmt.Print("Numbers must be in the range 1-10 included")
				os.Exit(0)
			}
			fmt.Print(calculation(number1, number2, oper))
		} else {
			fmt.Print("Different format of the numbers!")
			os.Exit(0)
		}
	} else {
		number1 := toArabic(a)
		number2 := toArabic(b)
		if calculation(number1, number2, oper) < 1 {
			fmt.Print("Result of operation is not positive!")
			os.Exit(0)
		}
		fmt.Print(toRoman(calculation(number1, number2, oper)))
	}
}
