package main

import (
	"fmt"
	"strconv"
)

func main() {
	str1, str2, operation, err := scan()
	if err != nil {
		fmt.Println("произошла ошибка:", err)
		return
	}

	num1, num2, isRome, err := parse(str1, str2)
	if err != nil {
		fmt.Println("произошла ошибка:", err)
		return
	}

	err = validate(num1, num2)
	if err != nil {
		fmt.Println("произошла ошибка:", err)
		return
	}

	result, err := calculate(num1, num2, operation)
	if err != nil {
		fmt.Println("произошла ошибка:", err)
		return
	}

	if isRome == true {
		resultRome, err := resultInRome(result)
		if err != nil {
			fmt.Println("произошла ошибка:", err)
			return
		}
		fmt.Println(resultRome)
		return
	}

	fmt.Println(result)

}

func scan() (a, b, c string, err error) {
	_, err = fmt.Scanf("%s %s %s\n", &a, &c, &b)
	return a, b, c, err
}

func parse(str1, str2 string) (num1, num2 int, isRome bool, err error) {
	num1, err = strconv.Atoi(str1)
	if err != nil {
		return parseRome(str1, str2)
	}

	num2, err = strconv.Atoi(str2)
	if err != nil {
		return num1, num2, false, err
	}

	return num1, num2, false, nil
}

func parseRome(str1, str2 string) (num1, num2 int, isRome bool, err error) {
	num1, err = parseOneRome(str1)
	if err != nil {
		return num1, num2, true, err
	}

	num2, err = parseOneRome(str2)
	if err != nil {
		return num1, num2, true, err
	}

	return num1, num2, true, nil
}

func parseOneRome(str string) (num int, err error) {
	switch str {
	case "I":
		return 1, nil
	case "II":
		return 2, nil
	case "III":
		return 3, nil
	case "IV":
		return 4, nil
	case "V":
		return 5, nil
	case "VI":
		return 6, nil
	case "VII":
		return 7, nil
	case "VIII":
		return 8, nil
	case "IX":
		return 9, nil
	case "X":
		return 10, nil
	default:
		return 0, fmt.Errorf("не удалось распознать число")
	}
}

func validate(num1, num2 int) (err error) {
	if num1 < 0 || num1 > 10 {
		return fmt.Errorf("недопустимое число")
	}

	if num2 < 0 || num2 > 10 {
		return fmt.Errorf("недопустимое число")
	}

	return nil
}

func calculate(num1, num2 int, operation string) (result int, err error) {
	switch operation {
	case "+":
		return num1 + num2, nil
	case "-":
		return num1 - num2, nil
	case "*":
		return num1 * num2, nil
	case "/":
		return num1 / num2, nil
	default:
		return 0, fmt.Errorf("недопустимая операция")
	}
}

func resultInRome(result int) (resultRome string, err error) {
	if result < 0 {
		return "", fmt.Errorf("число отрицательное")
	}

	ones := []string{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"}
	tens := []string{"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"}
	hunds := []string{"", "C"}

	h := hunds[result/100%10]
	te := tens[result/10%10]
	o := ones[result%10]

	return h + te + o, nil
}
