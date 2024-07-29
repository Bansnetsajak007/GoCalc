package repl

import (
	"Gocalc/evaluator"
	"Gocalc/lexer"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Start() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Advanced Arithmetic REPL")
    fmt.Println("Enter expressions, 'vars' to see variables, or 'exit' to quit")

	for {
		fmt.Print("> ")
		scanner.Scan()
		input := scanner.Text()

		switch input {
			case "exit":
				return
			case "vars":
				for k,v := range evaluator.Variables {
					fmt.Printf("%s = %v\n", k,v)
				}
				continue
		}

		if strings.Contains(input, "=") {
			parts := strings.SplitN(input, "=", 2)
			if len(parts) == 2 {
				varName := strings.TrimSpace(parts[0])
				tokens := lexer.Tokenize(strings.TrimSpace(parts[1]))
				result, err := evaluator.Evaluate(tokens)

				if err != nil {
					fmt.Println("Error: ", err)
				} else {
					evaluator.Variables[varName] = result
					fmt.Printf("%s = %v\n", varName, result)
				}
				continue
			}
		}
		tokens := lexer.Tokenize(input)
		result, err := evaluator.Evaluate(tokens)
		if err != nil {
            fmt.Println("Error:", err)
        } else {
            fmt.Println("Result:", result)
        }
	}
}