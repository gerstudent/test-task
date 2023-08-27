package main

import (
	"fmt"
	"testing"
)

var (
	tests = []struct {
		s    string
		want string
	}{
		{"{()}{}()", "Скобки сбалансированы"},
		{"{()}(])", "Скобки несбалансированы"},
		{"()", "Скобки сбалансированы"},
		{"()[]{}", "Скобки сбалансированы"},
		{"([)][]", "Скобки несбалансированы"},
		{"{[]}{}", "Скобки сбалансированы"},
		{")", "Скобки несбалансированы"},
		{"(]", "Скобки несбалансированы"},
	}
	functions = []func(s string) string{isValid}
)

func TestIsValid(t *testing.T) {
	for _, fun := range functions {
		for _, test := range tests {
			test := test

			t.Run(fmt.Sprint(test.s, test.want), func(t *testing.T) {
				t.Parallel()
				if have := fun(test.s); have != test.want {
					t.Errorf(`
					Ожидаемый результат: %+v
					Полученный результат: %+v`, test.want, have)
				}
			})

		}
	}
}
