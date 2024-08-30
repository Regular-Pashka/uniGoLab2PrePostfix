

package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
)

// Основная функция
func main() {
    var expression string
    for {
        printMenu()
        choice := readChoice()

        switch choice {
        case 1:
            expression = readExpression()
            fmt.Println("Введенное выражение:", expression)
        case 2:
            result := calculateInfix(expression)
            fmt.Println("Выражение в инфиксной форме:", expression)
            fmt.Println("Результат в инфиксной форме:", result)
        case 3:
            postfix := convertToPostfix(expression)
            result := calculatePostfix(postfix)
            fmt.Println("Выражение в постфиксной форме:", postfix)
            fmt.Println("Результат в постфиксной форме:", result)
        case 4:
            prefix := convertToPrefix(expression)
            result := calculatePrefix(prefix)
            fmt.Println("Выражение в префиксной форме:", prefix)
            fmt.Println("Результат в префиксной форме:", result)
        case 5:
            fmt.Println("Выход из программы.")
            return
        default:
            fmt.Println("Неверный выбор. Попробуйте снова.")
        }
    }
}

// Функции для меню и ввода
func printMenu() {
    fmt.Println("Меню")
    fmt.Println("1. Ввести выражение")
    fmt.Println("2. Посчитать выражение в инфиксной форме")
    fmt.Println("3. Посчитать выражение в постфиксной форме")
    fmt.Println("4. Посчитать выражение в префиксной форме")
    fmt.Println("5. Выход")
}

func readChoice() float64 {
    var choice float64
    fmt.Print("Ваш выбор: ")
    fmt.Scanln(&choice)
    return choice
}

func readExpression() string {
    reader := bufio.NewReader(os.Stdin)
    fmt.Print("Введите выражение: ")
    expression, _ := reader.ReadString('\n')
    return strings.TrimSpace(expression)
}



func calc(operator string, operandStack *[]int) int {
    right := (*operandStack)[len(*operandStack)-1]
    *operandStack = (*operandStack)[:len(*operandStack)-1]
    left := (*operandStack)[len(*operandStack)-1]
    *operandStack = (*operandStack)[:len(*operandStack)-1]

    switch operator {
    case "+":
        return left + right
    case "-":
        return left - right
    case "*":
        return left * right
    case "/":
        return left / right
    }
    return 0
}

func calculateInfix(exp string) int {
    operandStack := []int{}
    operatorStack := []string{}
    tokens := strings.Split(exp, "")

    for i := 0; i < len(tokens); i++ {
        if tokens[i] == " " {
            continue
        }
        if isDigit(tokens[i]) {
            num := 0
            for i < len(tokens) && isDigit(tokens[i]) {
                digit, _ := strconv.Atoi(tokens[i])
                num = num*10 + digit
                i++
            }
            i-- // Вернуться к последнему числу
            operandStack = append(operandStack, num)
        } else if tokens[i] == "(" {
            operatorStack = append(operatorStack, tokens[i])
        } else if tokens[i] == ")" {
            for operatorStack[len(operatorStack)-1] != "(" {
                ans := calc(operatorStack[len(operatorStack)-1], &operandStack)
                operatorStack = operatorStack[:len(operatorStack)-1]
                operandStack = append(operandStack, ans)
            }
            operatorStack = operatorStack[:len(operatorStack)-1] // Удаляем '('
        } else {
            operator := tokens[i]
            for len(operatorStack) > 0 && precedence(operator) <= precedence(operatorStack[len(operatorStack)-1]) {
                ans := calc(operatorStack[len(operatorStack)-1], &operandStack)
                operatorStack = operatorStack[:len(operatorStack)-1]
                operandStack = append(operandStack, ans)
            }
            operatorStack = append(operatorStack, operator)
        }
    }

    for len(operatorStack) > 0 {
        ans := calc(operatorStack[len(operatorStack)-1], &operandStack)
        operatorStack = operatorStack[:len(operatorStack)-1]
        operandStack = append(operandStack, ans)
    }
    return operandStack[len(operandStack)-1]
}

func isDigit(s string) bool {
    _, err := strconv.Atoi(s)
    return err == nil
}

// Проверка на операнд
func isOperand(token string) bool {
    _, err := strconv.ParseFloat(token, 64)
    return err == nil
}

// Определение приоритета операции
func precedence(op string) float64 {
    switch op {
    case "+", "-":
        return 1
    case "*", "/":
        return 2
    }
    return 0
}

// Применение оператора
func applyOperator(values []float64, operators []string) []float64 {
    if len(values) < 2 || len(operators) == 0 {
        fmt.Println("Ошибка: недостаточно операндов для операции")
        return values
    }

    operand2 := values[len(values) - 1]
    values = values[:len(values) - 1]
    operand1 := values[len(values) - 1]
    values = values[:len(values) - 1]
    operator := operators[len(operators) - 1]
    operators = operators[:len(operators) - 1]

    result := performOperation(operator, operand1, operand2)
    values = append(values, result)
    return values
}

// Выполнение операции
func performOperation(operator string, operand1, operand2 float64) float64 {
    switch operator {
    case "+":
        return operand1 + operand2
    case "-":
        return operand1 - operand2
    case "*":
        return operand1 * operand2
    case "/":
        return operand1 / operand2
    default:
        return 0
    }
}

// Преобразование в постфиксную форму
func convertToPostfix(expression string) string {
    precedence := map[string]int{
        "+": 1,
        "-": 1,
        "*": 2,
        "/": 2,
    }

    var output []string
    var stack []string
    tokens := strings.Fields(expression)

    for _, token := range tokens {
        if isOperand(token) {
            output = append(output, token)
        } else if token == "(" {
            stack = append(stack, token)
        } else if token == ")" {
            for len(stack) > 0 && stack[len(stack)-1] != "(" {
                output = append(output, stack[len(stack)-1])
                stack = stack[:len(stack)-1]
            }
            if len(stack) > 0 {
                stack = stack[:len(stack)-1] // Удаляем '('
            }
        } else { // Оператор
            for len(stack) > 0 && precedence[stack[len(stack)-1]] >= precedence[token] {
                output = append(output, stack[len(stack)-1])
                stack = stack[:len(stack)-1]
            }
            stack = append(stack, token)
        }
    }

    for len(stack) > 0 {
        output = append(output, stack[len(stack)-1])
        stack = stack[:len(stack)-1]
    }

    return strings.Join(output, " ")
}

// Вычисление постфиксного выражения
func calculatePostfix(expression string) float64 {
    stack := make([]float64, 0)
    tokens := strings.Fields(expression)

    for _, token := range tokens {
        if isOperand(token) {
            operand, _ := strconv.ParseFloat(token, 64)
            stack = append(stack, operand)
        } else {
            if len(stack) < 2 {
                fmt.Println("Ошибка: недостаточно операндов для операции", token)
                return 0
            }
            operand2 := stack[len(stack)-1]
            stack = stack[:len(stack)-1]
            operand1 := stack[len(stack)-1]
            stack = stack[:len(stack)-1]

            result := performOperation(token, operand1, operand2)
            stack = append(stack, result)
        }
    }

    return stack[0]
}

// Преобразование в префиксную форму
func convertToPrefix(expression string) string {
    precedence := map[string]int{
        "+": 1,
        "-": 1,
        "*": 2,
        "/": 2,
    }

    var output []string
    var stack []string
    tokens := strings.Fields(expression)

    // Обратный порядок для префиксной записи
    for i := len(tokens) - 1; i >= 0; i-- {
        token := tokens[i]
        if isOperand(token) {
            output = append(output, token)
        } else if token == ")" {
            stack = append(stack, token)
        } else if token == "(" {
            for len(stack) > 0 && stack[len(stack)-1] != ")" {
                output = append(output, stack[len(stack)-1])
                stack = stack[:len(stack)-1]
            }
            if len(stack) > 0 {
                stack = stack[:len(stack)-1] // Удаляем ')'
            }
        } else { // Оператор
            for len(stack) > 0 && precedence[stack[len(stack)-1]] > precedence[token] {
                output = append(output, stack[len(stack)-1])
                stack = stack[:len(stack)-1]
            }
            stack = append(stack, token)
        }
    }

    for len(stack) > 0 {
        output = append(output, stack[len(stack)-1])
        stack = stack[:len(stack)-1]
    }

    // Реверсируем результат для префиксной формы
    for i, j := 0, len(output)-1; i < j; i, j = i+1, j-1 {
        output[i], output[j] = output[j], output[i]
    }

    return strings.Join(output, " ")
}

// Вычисление префиксного выражения
func calculatePrefix(expression string) float64 {
    tokens := strings.Fields(expression)
    stack := make([]float64, 0)

    // Обратный порядок для вычисления
    for i := len(tokens) - 1; i >= 0; i-- {
        token := tokens[i]
        if isOperand(token) {
            operand, _ := strconv.ParseFloat(token, 64)
            stack = append(stack, operand)
        } else {
            if len(stack) < 2 {
                fmt.Println("Ошибка: недостаточно операндов для операции", token)
                return 0
            }
            operand1 := stack[len(stack)-1]
            stack = stack[:len(stack)-1]
            operand2 := stack[len(stack)-1]
            stack = stack[:len(stack)-1]

            result := performOperation(token, operand1, operand2)
            stack = append(stack, result)
        }
    }

    if len(stack) == 0 {
        fmt.Println("Ошибка: пустой стек после вычисления префиксного выражения")
        return 0
    }

    return stack[0]
}
