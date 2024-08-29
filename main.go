

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

func readChoice() int {
    var choice int
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

// Функция для вычисления инфиксного выражения
func calculateInfix(expression string) int {
    values := make([]int, 0) // Стек для операндов
    operators := make([]string, 0) // Стек для операторов
    tokens := strings.Fields(expression)

    for _, token := range tokens {
        if isOperand(token) {
            operand, _ := strconv.Atoi(token)
            values = append(values, operand)
        } else if token == "(" {
            operators = append(operators, token)
        } else if token == ")" {
            for len(operators) > 0 && operators[len(operators)-1] != "(" {
                values = applyOperator(values, operators)
            }
            if len(operators) > 0 {
                operators = operators[:len(operators)-1] // Удаляем '('
            }
        } else { // Оператор
            for len(operators) > 0 && precedence(operators[len(operators)-1]) >= precedence(token) {
                values = applyOperator(values, operators)
            }
            operators = append(operators, token)
        }
    }

    // Применяем оставшиеся операторы
    for len(operators) > 0 {
        values = applyOperator(values, operators)
    }

    if len(values) == 0 {
        fmt.Println("Ошибка: пустой стек после вычисления инфиксного выражения")
        return 0
    }

    return values[0]
}

// Проверка на операнд
func isOperand(token string) bool {
    _, err := strconv.Atoi(token)
    return err == nil
}

// Определение приоритета операции
func precedence(op string) int {
    switch op {
    case "+", "-":
        return 1
    case "*", "/":
        return 2
    }
    return 0
}

// Применение оператора
func applyOperator(values []int, operators []string) []int {
    if len(values) < 2 || len(operators) == 0 {
        fmt.Println("Ошибка: недостаточно операндов для операции")
        return values
    }

    operand2 := values[len(values)-1]
    values = values[:len(values)-1]
    operand1 := values[len(values)-1]
    values = values[:len(values)-1]
    operator := operators[len(operators)-1]
    operators = operators[:len(operators)-1]

    result := performOperation(operator, operand1, operand2)
    values = append(values, result)
    return values
}

// Выполнение операции
func performOperation(operator string, operand1, operand2 int) int {
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
func calculatePostfix(expression string) int {
    stack := make([]int, 0)
    tokens := strings.Fields(expression)

    for _, token := range tokens {
        if isOperand(token) {
            operand, _ := strconv.Atoi(token)
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
func calculatePrefix(expression string) int {
    tokens := strings.Fields(expression)
    stack := make([]int, 0)

    // Обратный порядок для вычисления
    for i := len(tokens) - 1; i >= 0; i-- {
        token := tokens[i]
        if isOperand(token) {
            operand, _ := strconv.Atoi(token)
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

