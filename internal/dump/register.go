package dump

import (
	"errors"

	"github.com/saliceti/yaq/internal/pipeline"
)

type dumpFunc func(pipeline.StructuredData) (string, error)
type dumpFuncRegister map[string]dumpFunc

var FunctionRegister = dumpFuncRegister{}

func Register(name string, dumpFunction dumpFunc) {
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
