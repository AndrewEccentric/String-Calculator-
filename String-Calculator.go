package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// +
func concatenateStrings(a, b string) (string, error) {
	if len(a) > 10 || len(b) > 10 {
		return "", errors.New("строка слишком длинная")
	}
	return a + b, nil
}

// -
func subtractStrings(a, b string) (string, error) {
	if len(a) > 10 || len(b) > 10 {
		return "", errors.New("строка слишком длинная")
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
		return "", errors.New("строка слишком длинная")
	}
	if b > 10 {
		return "", errors.New("число слишком длинное")
	}
	if b <= 0 {
		return "", errors.New("число должно быть положительным")
	}
	result := strings.Repeat(a, b)
	return result, nil
}

// : строки на число
func divideString(a string, b int) (string, error) {
	if len(a) > 10 {
		return "", errors.New("строка слишком длинная")
	}
	if b > 10 {
		return "", errors.New("число слишком длинное")
	}
	if b <= 0 {
		return "", errors.New("число должно быть положительным")
	}
	result := a[:len(a)/b]
	return result, nil
}

// для обработки введенного выражения и выполнения операции
func evaluateExpression(expression string) (string, error) {
	parts := strings.Split(expression, " ")

	if len(parts) != 3 {
		return "", errors.New("некорректное выражение")
	}

	operand1 := parts[0]
	operator := parts[1]
	operand2 := parts[2]

	if operator != "+" && operator != "-" && operator != "*" && operator != "/" {
		return "", errors.New("некорректный оператор")
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
			return "", errors.New("неподдерживаемая операция над строками")
		}

	} else {

		switch operator {
		case "*":
			num, err := strconv.Atoi(operand2)
			if err != nil {
				return "", errors.New("некорректное число")
			}
			return multiplyString(operand1, num)

		case "/":
			num, err := strconv.Atoi(operand2)
			if err != nil {
				return "", errors.New("некорректное число")
			}
			if num == 0 {
				return "", errors.New("деление на ноль недопустимо")
			}
			return divideString(operand1, num)
		default:
			return "", errors.New("неподдерживаемая операция с числами")
		}
	}
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

	fmt.Println("Результат:", truncateString(result))
}
