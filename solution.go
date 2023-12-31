package main

import (
	"fmt"
	"strings"
)

const Balanced = "Скобки сбалансированы"
const NotBalanced = "Скобки несбалансированы"

var brackets = map[rune]rune{
	'{': '}',
	'(': ')',
	'[': ']',
}

func main() {
	var str string
	fmt.Scan(&str)
	fmt.Println(isValid(str))
}

func isValid(s string) string {
	stack := make([]rune, 0)

	for _, elem := range s {
		if !strings.Contains("{}[]()", string(elem)) {
			return "Ошибка: строка должна содержать только скобки"
		}

		matchToOpen, isOpen := brackets[elem]
		if isOpen {
			stack = append(stack, matchToOpen)
			continue
		}

		ln := len(stack)
		if ln == 0 || stack[ln-1] != elem {
			return NotBalanced
		}

		stack = stack[:ln-1]
	}

	if len(stack) == 0 {
		return Balanced
	}
	return NotBalanced
}
