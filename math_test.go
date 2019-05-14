package badmath

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	assert := assert.New(t)

	result := Add(3, 4)
	assert.Equal(7, result)

	result = Add(0, 5)
	assert.Equal(5, result)

	nums := []int{2, 4, 8}
	result = Add(nums...)
	assert.Equal(14, result)
}

func TestMultiply(t *testing.T) {
	assert := assert.New(t)

	result := Multiply(2, 11)
	assert.Equal(22, result)

	result = Multiply(0, 10)
	assert.Equal(0, result)
}
