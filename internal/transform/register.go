package transform

import (
	"errors"

	"github.com/saliceti/yaq/internal/pipeline"
)

type transformFunctionType func(pipeline.StructuredData, string) (pipeline.StructuredData, error)

type functionRegister map[string]transformFunctionType

var FunctionRegister = functionRegister{}

func RegisterTransformFunction(name string, processMapFunction transformFunctionType) {
	FunctionRegister[name] = processMapFunction
}

func TransformWith(transformArg string, inputMap pipeline.StructuredData) (pipeline.StructuredData, error) {
	transformName, parameter := pipeline.SplitArg(transformArg)

	f, err := lookupTransformFunction(transformName)
	if err != nil {
		return nil, err
	}

	return f(inputMap, parameter)
}

func lookupTransformFunction(transformName string) (transformFunctionType, error) {
	if f, ok := FunctionRegister[transformName]; ok {
		return f, nil
	}
	return nil, errors.New("Unknown transform: " + transformName)
}
