package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMap(t *testing.T) {
	stub := []int{1, 2, 3, 4, 5}

	res := Map(stub, func(i int) (int, bool) {
		return i * 2, true
	})

	assert.Equal(t, res, []int{2, 4, 6, 8, 10})
}
