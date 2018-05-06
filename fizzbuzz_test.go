package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFizzBuzz(t *testing.T) {
	assert.Equal(t, configurableFizzbuzz(3), "fizz")
	assert.Equal(t, configurableFizzbuzz(6), "fizz")
	assert.Equal(t, configurableFizzbuzz(18), "fizz")

	assert.Equal(t, configurableFizzbuzz(5), "buzz")
	assert.Equal(t, configurableFizzbuzz(25), "buzz")
	assert.Equal(t, configurableFizzbuzz(50), "buzz")

	assert.Equal(t, configurableFizzbuzz(7), "pop")
	assert.Equal(t, configurableFizzbuzz(14), "pop")
	assert.Equal(t, configurableFizzbuzz(28), "pop")

	assert.Equal(t, configurableFizzbuzz(15), "fizz buzz")
	assert.Equal(t, configurableFizzbuzz(30), "fizz buzz")
	assert.Equal(t, configurableFizzbuzz(75), "fizz buzz")

	assert.Equal(t, configurableFizzbuzz(21), "fizz pop")
	assert.Equal(t, configurableFizzbuzz(42), "fizz pop")
	assert.Equal(t, configurableFizzbuzz(84), "fizz pop")

	assert.Equal(t, configurableFizzbuzz(35), "buzz pop")
	assert.Equal(t, configurableFizzbuzz(70), "buzz pop")
	assert.Equal(t, configurableFizzbuzz(140), "buzz pop")

	assert.Equal(t, configurableFizzbuzz(105), "fizz buzz pop")
	assert.Equal(t, configurableFizzbuzz(210), "fizz buzz pop")
	assert.Equal(t, configurableFizzbuzz(420), "fizz buzz pop")

	assert.Equal(t, configurableFizzbuzz(2), "2")
	assert.Equal(t, configurableFizzbuzz(11), "11")
	assert.Equal(t, configurableFizzbuzz(103), "103")
}
