package ordering_test

import (
	"testing"

	"github.com/JainyMartins/testsGo/go-testing/ordering"
	"github.com/stretchr/testify/assert"
)

func TestOrder(t *testing.T){
	numbers := []int{9,1,2,7,4,8,10}

	expectedResult := []int{1,2,4,7,8,9,10}
	result := ordering.OrderArray(numbers)

	assert.Equal(t, expectedResult, result, "devem ser iguais")
}