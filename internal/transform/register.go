package transform

import (
	"errors"
	"strings"

	"github.com/saliceti/yaq/internal/pipeline"
)

type transformFunctionType func(pipeline.StructuredData, string) (pipeline.StructuredData, error)

type functionRegister map[string]transformFunctionType

var FunctionRegister = functionRegister{}

func RegisterTransformFunction(name string, processMapFunction transformFunctionType) {
	FunctionRegister[name] = processMapFunction
}

func TransformWith(transformArg string, inputMap pipeline.StructuredData) (pipeline.StructuredData, error) {
	transformArgArray := strings.Split(transformArg, ":")
	if f, ok := FunctionRegister[transformArgArray[0]]; ok {
		var parameter string
		if len(transformArgArray) == 1 {
			parameter = ""
		} else {
			parameter = transformArgArray[1]
		}
		outputMap, err := f(inputMap, parameter)

		if err != nil {
			return nil, err
		}
		return outputMap, nil
	}
	return nil, errors.New("Unknown transformation: " + transformArg)
}
