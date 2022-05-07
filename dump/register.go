package dump

import (
	"errors"
)

type dumpFunctionType func(interface{}) (string, error)
type dumpFunctionRegister map[string]dumpFunctionType

var FunctionRegister = dumpFunctionRegister{}

func Register(name string, dumpFunction dumpFunctionType) {
	FunctionRegister[name] = dumpFunction
}

func MapToString(format string, inputData interface{}) (string, error) {
	if f, ok := FunctionRegister[format]; ok {
		outputString, err := f(inputData)
		if err != nil {
			return "", err
		}
		return outputString, nil
	}
	return "", errors.New("Unknown dump format: " + format)
}
