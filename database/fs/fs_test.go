package fs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type MyData struct {
	Name   string `db:"index"`
	Age    int
	Gender *string
}

func TestSaveData(t *testing.T) {
	expect := MyData{
		Name:   "Test",
		Age:    69,
		Gender: nil,
	}

	err := insert(expect)
	assert.NoError(t, err)

	data, err := findOne[MyData]("Test")
	assert.NoError(t, err)
	assert.Equal(t, expect, data)
}

func BenchmarkSaveData(b *testing.B) {
	expect := MyData{
		Name:   "Test",
		Age:    69,
		Gender: nil,
	}

	err := insert(expect)
	_, err = findOne[MyData]("Test")
	if err != nil {
		b.FailNow()
	}
	b.ReportAllocs()
}
