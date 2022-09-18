package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSum(t *testing.T) {
	intToSum := []int{2, 3, 4, 9}
	sumExpected := 7
	res := FindSum(intToSum, sumExpected)
	assert.True(t, res)
}
