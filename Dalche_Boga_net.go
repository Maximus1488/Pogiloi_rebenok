package main

import (
	"fmt"
	"strconv"
)

func main() {
	var sraw1, sraw2 bool
	var operator, numStr1, numStr2 string
	var result int
	romaMap := map[string]string{
		"I":    "1",
		"II":   "2",
		"III":  "3",
		"IV":   "4",
		"V":    "5",
		"VI":   "6",
		"VII":  "7",
		"VIII": "8",
		"IX":   "9",
		"X":    "10",
	} //Из римских в арабские
	arabMap := map[string]string{
		"1":   "I",
		"2":   "II",
		"3":   "III",
		"4":   "IV",
		"5":   "V",
		"6":   "VI",
		"7":   "VII",
		"8":   "VIII",
		"9":   "IX",
		"10":  "X",
		"11":  "XI",
		"12":  "XII",
		"13":  "XIII",
		"14":  "XIV",
		"15":  "XV",
		"16":  "XVI",
		"17":  "XVII",
		"18":  "XVIII",
		"19":  "XIX",
		"20":  "XX",
		"21":  "XXI",
		"22":  "XXII",
		"23":  "XXIII",
		"24":  "XXIV",
		"25":  "XXV",
		"26":  "XXVI",
		"27":  "XXVII",
		"28":  "XXVIII",
		"29":  "XXIX",
		"30":  "XXX",
		"31":  "XXXI",
		"32":  "XXXII",
		"33":  "XXXIII",
		"34":  "XXXIV",
		"35":  "XXXV",
		"36":  "XXXVI",
		"37":  "XXXVII",
		"38":  "XXXVIII",
		"39":  "XXXIX",
		"40":  "XL",
		"41":  "XLI",
		"42":  "XLII",
		"43":  "XLIII",
		"44":  "XLIV",
		"45":  "XLV",
		"46":  "XLVI",
		"47":  "XLVII",
		"48":  "XLVIII",
		"49":  "XLIX",
		"50":  "L",
		"51":  "LI",
		"52":  "LII",
		"53":  "LIII",
		"54":  "LIV",
		"55":  "LV",
		"56":  "LVI",
		"57":  "LVII",
		"58":  "LVIII",
		"59":  "LIX",
		"60":  "LX",
		"61":  "LXI",
		"62":  "LXII",
		"63":  "LXIII",
		"64":  "LXIV",
		"65":  "LXV",
		"66":  "LXVI",
		"67":  "LXVII",
		"68":  "LXVIII",
		"69":  "LXIX",
		"70":  "LXX",
		"71":  "LXXI",
		"72":  "LXXII",
		"73":  "LXXIII",
		"74":  "LXXIV",
		"75":  "LXXV",
		"76":  "LXXVI",
		"77":  "LXXVII",
		"78":  "LXXVIII",
		"79":  "LXXIX",
		"80":  "LXXX",
		"81":  "LXXXI",
		"82":  "LXXXII",
		"83":  "LXXXIII",
		"84":  "LXXXIV",
		"85":  "LXXXV",
		"86":  "LXXXVI",
		"87":  "LXXXVII",
		"88":  "LXXXVIII",
		"89":  "LXXXIX",
		"90":  "XC",
		"91":  "XCI",
		"92":  "XCII",
		"93":  "XCIII",
		"94":  "XCIV",
		"95":  "XCV",
		"96":  "XCVI",
		"97":  "XCVII",
		"98":  "XCVIII",
		"99":  "XCIX",
		"100": "C",
	} //Из арабских в римские
	fmt.Println("Введите выражение арабскими или римскими цифрами. Пример:[II * VI], [8 - 3]. Введённые числа должны быть целыми, одного типа и в интервале от 1 до 10.\n")
	fmt.Scan(&numStr1, &operator, &numStr2)

	if _, ok1 := romaMap[numStr1]; ok1 {
		sraw1 = true
	} else {
		sraw1 = false
	} //Проверка первого значения на римские цифры
	if _, ok1 := romaMap[numStr2]; ok1 {
		sraw2 = true
	} else {
		sraw2 = false
	} //Проверка второго значения на римские цифры

	if sraw1 != false && sraw2 != false {

		arabNum1 := romaMap[numStr1]
		arabNum2 := romaMap[numStr2]

		num1, _ := strconv.Atoi(arabNum1)
		num2, _ := strconv.Atoi(arabNum2)
		switch operator {
		case "+":
			if operator == "+" {
				result = (num1 + num2)
			}
		case "-":
			if operator == "-" {
				result = (num1 - num2)
			}
		case "*":
			if operator == "*" {
				result = (num1 * num2)
			}
		case "/":
			if num2 == 0 {
				fmt.Println("Нельзя делить на ноль")
			}

			if operator == "/" {
				result = (num1 / num2)
			}

		} //калькулятор

		if result < 1 {
			panic("В риских цифрах отсутствует ноль и отрицательные значения.\n")
		}
		romaNumS1 := strconv.Itoa(num1)
		romaNumS2 := strconv.Itoa(num2)
		romaResultS := strconv.Itoa(result)

		romaNum1 := arabMap[romaNumS1]
		romaNum2 := arabMap[romaNumS2]
		romaResult := arabMap[romaResultS]

		fmt.Println(romaNum1, operator, romaNum2, "=", romaResult)
	} else {
		numa1, _ := strconv.Atoi(numStr1)
		numa2, _ := strconv.Atoi(numStr2)

		if numa1 < 1 || numa1 > 10 || numa2 < 1 || numa2 > 10 {
			panic("Введённые числа должны быть целые, одного типа и в интервале от 1 до 10.\n")
		}

		switch operator {
		case "+":
			if operator == "+" {
				result = (numa1 + numa2)
			}
		case "-":
			if operator == "-" {
				result = (numa1 - numa2)
			}
		case "*":
			if operator == "*" {
				result = (numa1 * numa2)
			}
		case "/":
			if numa2 == 0 {
				fmt.Println("Нельзя делить на ноль")
			}

			if operator == "/" {
				result = (numa1 / numa2)
			}

		} //калькулятор

		arabNumS1 := strconv.Itoa(numa1)
		arabNumS2 := strconv.Itoa(numa2)
		arabResultS := strconv.Itoa(result)

		fmt.Println(arabNumS1, operator, arabNumS2, "=", arabResultS)
	}

}
