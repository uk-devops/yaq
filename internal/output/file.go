package output

import (
	"errors"
	"os"
)

func init() {
	RegisterStringFunction("file", WriteToFile)
}

func WriteToFile(inputString, path string) error {
	if path == "" {
		return errors.New("output file path cannot be empty")
	}
	err := os.WriteFile(path, []byte(inputString), 0755)

	if err != nil {
		return errors.New("can't write to file: " + err.Error())
	}

	return nil

}
