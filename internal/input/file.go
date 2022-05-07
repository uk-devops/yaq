package input

import (
	"io/ioutil"
)

func init() {
	RegisterStringFunction("file", ReadFromFile)
}

func ReadFromFile(path string) (string, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}

	return string(data), nil
}
