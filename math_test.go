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
}
