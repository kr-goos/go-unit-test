package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Add(a, b int) int {
	return a + b
}

func TestAdd(t *testing.T) {
	sum := Add(2, 3)
	assert.Equal(t, 5, sum)
	assert.NotEqual(t, 4, sum)
	assert.True(t, sum > 0)
}
