package fs

import (
	"fmt"
	"reflect"
)

const tagKeyword = "db"

type tag string

const (
	tagIndex tag = "index"
)

var ErrorInvalidIndex = fmt.Errorf("invalid index struct")

func getIndex(data any) (string, error) {
	val := reflect.ValueOf(data)
	if val.Kind() != reflect.Struct {
		return "", ErrorInvalidIndex
	}

	for i := 0; i < val.NumField(); i++ {
		if tag, ok := val.Type().Field(i).Tag.Lookup(tagKeyword); ok && tag == string(tagIndex) {
			return val.Field(i).String(), nil
		}
	}

	return "", ErrorInvalidIndex
}
