package calc

import (
    "errors"
)

func Sum(num1, num2 int) int {
    return num1 + num2
}

// Função que recebe dois inteiros e retorna a diferença 
// ou diferença resultante
func Sub(num1, num2 int) int {
    return num1 - num2
}

func Divide(num, den int) (*int, error) {
    if den == 0 {
        return nil, errors.New("The denominator cannot be 0.")
    }

    result := num / den

    return &result, nil
}