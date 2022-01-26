package main

import (
	"io/ioutil"
)

func loadFile(path string) (string, error) {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}
	str := string(b)
	return str, nil
}

func loadSchema() (string, error) {
	return loadFile("schema.graphql")
}

func loadPlayground() (string, error) {
	htmlFile, err := loadFile("graphql_playground.template")
	if err != nil {
		return "", err
	}
	return string(htmlFile), nil
}
