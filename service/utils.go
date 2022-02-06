package service

import "encoding/json"

func Map[T, V any](arr []V, mapFn func(V) (T, bool)) []T {
	newArr := make([]T, 0, len(arr))
	for _, v := range arr {
		if mapped, ok := mapFn(v); ok {
			newArr = append(newArr, mapped)
		}
	}
	return newArr
}

func autoCreateMap[T any, S any](src S) (T, error) {
	target := new(T)
	sourceJson, err := json.Marshal(src)
	if err != nil {
		return *target, err
	}
	err = json.Unmarshal(sourceJson, target)
	if err != nil {
		return *target, err
	}
	return *target, nil
}
