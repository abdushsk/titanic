package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRead(t *testing.T) {
	v := readHumans("assets/values.csv")
	assert.NotEqual(t, 0, len(v))
}
