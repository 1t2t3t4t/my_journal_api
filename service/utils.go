package service

import "encoding/json"

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
