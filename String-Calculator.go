package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func validateInput(a string, b int) error {
	if len(a) > 10 {
		return errors.New("строка слишком длинная")
	}
	if b > 10 {
		return errors.New("число слишком длинное")
	}
	if b <= 0 {
		return errors.New("число должно быть положительным")
	}
	return nil
}
func concatenateStrings(a, b string) (string, error) {
	if len(a) > 10 || len(b) > 10 {
		return "", errors.New("строка слишком длинная")
	}
	return a + b, nil
}

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

func multiplyString(a string, b int) (string, error) {
	if err := validateInput(a, b); err != nil {
		return "", err
	}
	result := strings.Repeat(a, b)
	return result, nil
}
func divideString(a string, b int) (string, error) {
	if err := validateInput(a, b); err != nil {
		return "", err
	}
	result := a[:len(a)/b]
	return result, nil
}

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

	if !strings.HasPrefix(operand1, "\"") || !strings.HasSuffix(operand1, "\"") ||
		!strings.HasPrefix(operand2, "\"") || !strings.HasSuffix(operand2, "\"") {
		return "", errors.New("некорректное выражение")
	}

	operand1 = operand1[1 : len(operand1)-1]
	operand2 = operand2[1 : len(operand2)-1]

	switch operator {
	case "+":
		return concatenateStrings(operand1, operand2)
	case "-":
		return subtractStrings(operand1, operand2)
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
		return "", errors.New("неподдерживаемая операция")
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
	expression, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Ошибка при чтении строки:", err)
		return
	}
	expression = strings.TrimSpace(expression)

	result, err := evaluateExpression(expression)
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}

	fmt.Println("Результат:", truncateString(result))
}
