package evaluator

import (
	"fmt"
	"math"
	"strconv"
)

//storing key-value pairs
var Variables = make(map[string]float64)

func Evaluate(tokens []string) (float64, error) {
	return evalExpression(tokens)
}

func evalExpression(tokens []string) (float64, error) {
	var stack []float64  //to hold operands
	var ops []string  //to hold operators

	for i := 0; i < len(tokens); i++ {
		token := tokens[i]

		switch token {
		case "(":
			ops = append(ops, token)
		case")":
			for len(ops) > 0 && ops[len(ops)-1] != "(" {
				if err:= applyOp(&stack, &ops); err != nil {
					return 0, err
				}
			}
			if len(ops) == 0 {
				return 0, fmt.Errorf("mismatched parentheses")
			}

			ops = ops[:len(ops)-1]

		case "+", "-", "*", "/", "^":
			for len(ops) > 0 && precedence(ops[len(ops)-1]) >= precedence(token) {
				if err := applyOp(&stack, &ops); err != nil {
                    return 0, err
                }
			}

			ops = append(ops, token)
		
		default:
            if val, ok := Variables[token]; ok {
                stack = append(stack, val)
            } else if val, err := strconv.ParseFloat(token, 64); err == nil {
                stack = append(stack, val)
            } else if fn, ok := mathFuncs[token]; ok {
                if i+1 >= len(tokens) || tokens[i+1] != "(" {
                    return 0, fmt.Errorf("expected '(' after function %s", token)
                }
                j := i + 2
                parenCount := 1
                for ; j < len(tokens) && parenCount > 0; j++ {
                    if tokens[j] == "(" {
                        parenCount++
                    } else if tokens[j] == ")" {
                        parenCount--
                    }
                }
                if parenCount != 0 {
                    return 0, fmt.Errorf("mismatched parentheses in function call")
                }
                arg, err := evalExpression(tokens[i+2 : j-1])
                if err != nil {
                    return 0, err
                }
                stack = append(stack, fn(arg))
                i = j - 1  // Skip to end of function call
            } else {
                return 0, fmt.Errorf("unknown token: %s", token)
            }
	}
 }
 	for len(ops) > 0 {
		if err := applyOp(&stack, &ops); err != nil {
			return 0, err
		}
	}

	if len(stack) != 1 {
        return 0, fmt.Errorf("invalid expression")
    }
	return stack[0], nil
}

var mathFuncs = map[string]func(float64) float64 {
	"sin":   math.Sin,
    "cos":   math.Cos,
    "tan":   math.Tan,
    "sqrt":  math.Sqrt,
    "log":   math.Log,
    "exp":   math.Exp,
    "abs":   math.Abs,
    "floor": math.Floor,
    "ceil":  math.Ceil,
}