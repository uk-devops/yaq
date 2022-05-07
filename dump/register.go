package dump

import (
	"errors"

	"github.com/saliceti/yaq/pipeline"
)

type dumpFunctionType func(pipeline.StructuredData) (string, error)
type dumpFunctionRegister map[string]dumpFunctionType

var FunctionRegister = dumpFunctionRegister{}

func Register(name string, dumpFunction dumpFunctionType) {
	FunctionRegister[name] = dumpFunction
}

func MapToString(format string, inputData pipeline.StructuredData) (string, error) {
	if f, ok := FunctionRegister[format]; ok {
		outputString, err := f(inputData)
		if err != nil {
			return "", err
		}
		return outputString, nil
	}
	return "", errors.New("Unknown dump format: " + format)
}
