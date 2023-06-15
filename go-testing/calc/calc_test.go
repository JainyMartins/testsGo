package calc_test

import (
	"testing"

	"github.com/JainyMartins/testsGo/go-testing/calc"
    "github.com/stretchr/testify/assert"
)

func TestSum(t *testing.T) {
    num1 := 3
    num2 := 5
    expectedResult := 8

    result := calc.Sum(num1, num2)

    assert.Equal(t, expectedResult, result, "devem ser iguais")
}

func TestDivideError(t *testing.T) {
    num := 10
    den := 0

    _, err := calc.Divide(num, den)

    assert.EqualError(t, err, "The denominator cannot be 0.")
}

