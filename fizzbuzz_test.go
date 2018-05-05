package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFizzBuzz(t *testing.T) {
	assert.Equal(t, fizzbuzz(3), "fizz")
	assert.Equal(t, fizzbuzz(6), "fizz")
	assert.Equal(t, fizzbuzz(18), "fizz")

	assert.Equal(t, fizzbuzz(5), "buzz")
	assert.Equal(t, fizzbuzz(25), "buzz")
	assert.Equal(t, fizzbuzz(50), "buzz")

	assert.Equal(t, fizzbuzz(15), "fizz-buzz")
	assert.Equal(t, fizzbuzz(30), "fizz-buzz")
	assert.Equal(t, fizzbuzz(75), "fizz-buzz")

	assert.Equal(t, fizzbuzz(2), "2")
	assert.Equal(t, fizzbuzz(11), "11")
	assert.Equal(t, fizzbuzz(103), "103")
}
