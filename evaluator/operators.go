package evaluator

import (
	"fmt"
	"math"
)

func applyOp(stack *[]float64, ops *[]string) error {
	if len(*stack) < 2 {
		return fmt.Errorf("not enough operands")
	}
	op := (*ops)[len(*ops)-1]
	*ops = (*ops)[:len(*ops)-1]
	b, a := (*stack)[len(*stack)-1], (*stack)[len(*stack)-2]
	*stack = (*stack)[:len(*stack)-2]
	var result float64
	switch op {
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "*":
		result = a * b
	case "/":
		if b == 0 {
			return fmt.Errorf("division by zero")
		}
		result = a / b
	case "^":
		result = math.Pow(a, b)
	}
	*stack = append(*stack, result)
	return nil
}

func precedence(op string) int {
	switch op {
	case "+", "-":
		return 1
	case "*", "/":
		return 2
	case "^":
		return 3
	}
	return 0
}