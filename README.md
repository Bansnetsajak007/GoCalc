# GoCalc: Interactive command-line calculator

GoCalc is an advanced, interactive command-line calculator written in Go. It combines the simplicity of a basic calculator with the power of a scientific computing tool.

## Features

- Interactive REPL (Read-Eval-Print Loop) interface
- Complex mathematical expressions
- Basic arithmetic operations (+, -, *, /, ^)
- Trigonometric functions (sin, cos, tan)
- Logarithmic and exponential functions
- Variable storage and recall
- Parentheses for expression grouping

## Installation

1. Ensure you have Go installed on your system. If not, download and install it from [golang.org](https://golang.org/).

2. Clone the repository:
   ```
   git clone https://github.com/Bansnetsajak007/GoCalc.git
   cd gocalc
   ```

3. Build the project:
   ```
   go build
   ```

## Usage

Run the program:
```
./gocalc
```

Once started, you can enter mathematical expressions. For example:

```
> 2 + 3 * 4
Result: 14

> sin(pi/2)
Result: 1

> x = 5
x = 5

> y = x^2 + 3
y = 28

> sqrt(y)
Result: 5.291502622129181
```

Special commands:
- `vars`: Display all stored variables
- `exit`: Quit the program

## Project Structure

- `main.go`: Entry point of the application
- `evaluator/`: Package for expression evaluation
- `lexer/`: Package for tokenizing input
- `repl/`: Package for the interactive interface
