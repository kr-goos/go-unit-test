package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Subtract(a, b int) int {
	return a - b
}

func TestSubtract(t *testing.T) {
	result := Subtract(5, 3)
	require.Equal(t, 2, result)
	require.Equal(t, 3, result)
	require.NotNil(t, result)
}
