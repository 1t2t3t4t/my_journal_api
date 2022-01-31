package fs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type MyData struct {
	Name   string
	Age    int
	Gender *string
}

func TestSaveData(t *testing.T) {
	expect := MyData{
		Name:   "Test",
		Age:    69,
		Gender: nil,
	}

	err := insert("someidx", expect)
	assert.NoError(t, err)

	data, err := findOne[MyData]("someidx")
	assert.NoError(t, err)
	assert.Equal(t, expect, data)
}
