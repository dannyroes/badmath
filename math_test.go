package badmath

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	assert := assert.New(t)

	result, err := Add(3, 4)
	assert.Nil(err)
	assert.Equal(7, result)

	result, err = Add(0, 5)
	assert.Equal(5, result)
	assert.Nil(err)

	nums := []int{2, 4, 8}
	result, err = Add(nums...)
	assert.Equal(14, result)
	assert.Nil(err)
}

func TestMultiply(t *testing.T) {
	assert := assert.New(t)

	result := Multiply(2, 11)
	assert.Equal(22, result)

	result = Multiply(0, 10)
	assert.Equal(0, result)
}

func TestDivide(t *testing.T) {
	assert := assert.New(t)

	result, err := Divide(25, 10)
	assert.Equal(2.5, result)
	assert.Nil(err)

	result, err = Divide(20, 5)
	assert.Equal(4.0, result)
	assert.Nil(err)
}
