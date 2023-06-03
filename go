package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// +
func concatenateStrings(a, b string) (string, error) {
	return a + b, nil
}

// -
func subtractStrings(a, b string) (string, error) {
	if strings.Contains(a, b) {
		return strings.Replace(a, b, "", 1), nil
	} else {
		return a, nil
	}
}

// строка на число
func multiplyString(a string, b int) (string, error) {
	if b <= 0 {
		return "", errors.New("Число должно быть положительным")
	}

	result := ""
	for i := 0; i < b; i++ {
		result += a
		if len(result) > 10 {
			return "", errors.New("Результат слишком длинный")
		}
	}
	return result, nil
}

// : строки на число
func divideString(a string, b int) (string, error) {
	if b <= 0 {
		return "", errors.New("Число должно быть положительным")
	}

	if len(a)/b > 10 {
		return "", errors.New("Результат слишком длинный")
	}
	return a[:len(a)/b], nil
}

// для обработки введенного выражения и выполнения операции
func evaluateExpression(expression string) (string, error) {
	parts := strings.Split(expression, " ")

	if len(parts) != 3 {
		return "", errors.New("Некорректное выражение")
	}

	operand1 := parts[0]
	operator := parts[1]
	operand2 := parts[2]

	if operator != "+" && operator != "-" && operator != "*" && operator != "/" {
		return "", errors.New("Некорректный оператор")
	}

	if strings.HasPrefix(operand1, "\"") && strings.HasSuffix(operand1, "\"") &&
		strings.HasPrefix(operand2, "\"") && strings.HasSuffix(operand2, "\"") {
		operand1 = operand1[1 : len(operand1)-1]
		operand2 = operand2[1 : len(operand2)-1]

		switch operator {
		case "+":
			return concatenateStrings(operand1, operand2)
		case "-":
			return subtractStrings(operand1, operand2)
		default:
			return "", errors.New("Неподдерживаемая операция над строками")
		}
	} else if num1, err := strconv.Atoi(operand1); err == nil {
		switch operator {
		case "*":
			num2, err := strconv.Atoi(operand2)
			if err != nil {
				return "", errors.New("Некорректное число")
			}
			return multiplyString(operand1, num2)
		case "/":
			num2, err := strconv.Atoi(operand2)
			if err != nil {
				return "", errors.New("Некорректное число")
			}
			return divideString(operand1, num2)
		default:
			return "", errors.New("Неподдерживаемая операция с числами")
		}
	}
	return "", errors.New("Некорректные операнды")
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Некорректные аргументы")
		return
	}

	expression := os.Args[1]
	result, err := evaluateExpression(expression)
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}

	fmt.Println(result)
}
