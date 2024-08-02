package repl

import (
	"Gocalc/evaluator"
	"Gocalc/lexer"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func help(){
	fmt.Println("\n\033[1mGoCalc v1.0 - Command Line Calculator\033[0m")
    fmt.Println("\033[1mUSAGE:\033[0m")
    fmt.Println("  command [expression]")

    fmt.Println("\n\033[1mAVAILABLE COMMANDS:\033[0m")
    commands := []struct {
        name        string
        description string
        example     string
    }{
        {"calc", "Evaluate a mathematical expression", "calc 2 + 3 * (4 - 1)"},
        {"var", "Set or recall a variable", "var x = 10 or var x"},
        {"sin/cos/tan", "Trigonometric functions (in radians)", "sin(pi/2)"},
        {"log/exp", "Logarithmic and exponential functions", "log(100) or exp(2)"},
        {"help", "Display this help message", "help"},
        {"quit/exit", "Exit the calculator", "quit"},
    }

    for _, cmd := range commands {
        fmt.Printf("  \033[1m%-12s\033[0m %s\n", cmd.name, cmd.description)
        fmt.Printf("    Example: %s\n", cmd.example)
    }

    fmt.Println("\n\033[1mNOTE:\033[0m")
    fmt.Println("  - Use parentheses () for grouping expressions")
    fmt.Println("  - Variable names are case-sensitive")
    fmt.Println("  - Built-in constants: pi, e")
    fmt.Println("\nFor more detailed information, visit: https://gocalc.example.com/docs")
}

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
			case "help":
				help()
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