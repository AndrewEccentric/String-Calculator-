package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"bufio"
)

// +
func concatenateStrings(a, b string) (string, error) {
	if len(a) > 10 || len(b) > 10 {
		return "", errors.New("Строка слишком длинная")
	}
	return a + b, nil
}

// -
func subtractStrings(a, b string) (string, error) {
	if len(a) > 10 || len(b) > 10 {
		return "", errors.New("Строка слишком длинная")
	}
	if strings.Contains(a, b) {
		return strings.Replace(a, b, "", 1), nil
	} else {
		return a, nil
	}
}

// строка на число
func multiplyString(a string, b int) (string, error) {
	if len(a) > 10 {
		return "", errors.New("Строка слишком длинная")
	}
	if len(b) > 10 {
		return "", errors.New("Число слишком длинное")
	}
	if b <= 0 {
		return "", errors.New("Число должно быть положительным")
	}
    result := strings.Repeat(a, b)
		} else {
	        return result, nil
		}

// : строки на число
func divideString(a string, b int) (string, error) {
	if len(a) > 10 {
		return "", errors.New("Строка слишком длинная")
	}
	if len(b) > 10 {
		return "", errors.New("Число слишком длинное")
	}
	if b <= 0 {
		return "", errors.New("Число должно быть положительным")
	}
    result := a[:len(a)/b]
	} else {
	return result, nil
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

func truncateString(s string) string {
	if len(s) > 40 {
		return s[:40] + "..."
	}
	return s 
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Введите выражение: ")
	expression, _ := reader.ReadString('\n')
	expression = strings.TrimSpace(expression)

	result, err := evaluateExpression(expression)
	if err != nil {
		fmt.Println("Ошибка", err)
		return
	}

	fmt.Println("Результат:", result)
}
