package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var re = regexp.MustCompile(`(\d+)([-+*/])(\d+)`)

// calculateExpression вычисляет результат математического выражения
func calculateExpression(expression string) (string, error) {
	match := re.FindStringSubmatch(expression)

	if len(match) != 4 {
		return "", fmt.Errorf("invalid expression format: %s", expression)
	}

	num1, err := strconv.ParseFloat(match[1], 64)
	if err != nil {
		return "", fmt.Errorf("invalid number: %s", match[1])
	}
	num2, err := strconv.ParseFloat(match[3], 64)
	if err != nil {
		return "", fmt.Errorf("invalid number: %s", match[3])
	}
	operator := match[2]

	var result float64
	switch operator {
	case "+":
		result = num1 + num2
	case "-":
		result = num1 - num2
	case "*":
		result = num1 * num2
	case "/":
		if num2 == 0 {
			return "", fmt.Errorf("division by zero")
		}
		result = num1 / num2
	default:
		return "", fmt.Errorf("invalid operator: %s", operator)
	}
	return fmt.Sprintf("%s=%.2f", expression, result), nil
}

func main() {
	if len(os.Args) != 3 {
		fmt.Println("usage: go run main.go input_file output_file")
		return
	}

	inputFile := os.Args[1]
	outputFile := os.Args[2]

	content, err := os.ReadFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(content), "\n")

	file, err := os.Create(outputFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)

	for _, line := range lines {
		line = strings.TrimSpace(strings.ReplaceAll(line, "?", ""))

		if line == "" {
			continue
		}

		result, err := calculateExpression(line)
		if err != nil {
			continue
		}

		_, err = writer.WriteString(result + "\n")
		if err != nil {
			log.Fatal(err)
		}
	}

	err = writer.Flush()
	if err != nil {
		log.Fatal(err)
	}
}
