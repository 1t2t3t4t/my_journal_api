package fs

import (
	"bytes"
	"compress/zlib"
	"encoding/json"
	"fmt"
	"os"
	"path"
	"reflect"
	"strings"
)

var dbPath string = "./db"

func SetDbPath(path string) {
	dbPath = path
}

func getPureTypeName[T any]() string {
	val := new(T)
	typeName := reflect.TypeOf(val)
	splitName := strings.Split(typeName.String(), ".")
	if len(splitName) > 1 {
		return splitName[1]
	} else {
		return splitName[0]
	}
}

func decode[T any](source []byte) (T, error) {
	var res T
	r, err := zlib.NewReader(bytes.NewReader(source))
	if err != nil {
		return res, err
	}
	deZip := new(bytes.Buffer)
	_, err = deZip.ReadFrom(r)
	if err != nil {
		return res, err
	}
	err = json.Unmarshal(deZip.Bytes(), &res)
	if err != nil {
		return res, err
	}
	return res, nil
}

func encode[T any](data T) ([]byte, error) {
	dataJson, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	var enc bytes.Buffer
	w := zlib.NewWriter(&enc)
	_, err = w.Write(dataJson)
	if err != nil {
		return nil, err
	}
	err = w.Close()
	if err != nil {
		return nil, err
	}
	return enc.Bytes(), nil
}

func dataPath[T any](index string) (string, error) {
	fileName := getPureTypeName[T]()
	dataDirPath := path.Join(dbPath, fileName, index)
	file, err := os.Stat(dataDirPath)
	if err != nil {
		err = os.MkdirAll(dataDirPath, os.ModePerm)
		return dataDirPath, err
	} else if !file.IsDir() {
		return "", fmt.Errorf("invalid database path %v: is not a dir", dataDirPath)
	}
	return dataDirPath, nil
}

func findOne[T any](index string) (T, error) {
	var out T
	dataDirPath, err := dataPath[T](index)
	if err != nil {
		return out, err
	}
	filePath := path.Join(dataDirPath, "ent")
	content, err := os.ReadFile(filePath)
	if err != nil {
		return out, err
	}
	return decode[T](content)
}

func insert[T any](data T) error {
	index, err := getIndex(data)
	if err != nil {
		return err
	}
	dataDirPath, err := dataPath[T](index)
	if err != nil {
		return err
	}
	filePath := path.Join(dataDirPath, "ent")
	encData, err := encode(data)
	if err != nil {
		return err
	}
	return os.WriteFile(filePath, encData, os.ModePerm)
}
