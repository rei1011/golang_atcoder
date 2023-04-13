package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSum(t *testing.T) {
	actual := 8
	expected := 8

	assert.Equal(t, expected, actual)
}
