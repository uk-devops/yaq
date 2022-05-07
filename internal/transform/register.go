package transform

import (
	"errors"

	"github.com/saliceti/yaq/internal/pipeline"
)

type transformFunc func(pipeline.StructuredData, string) (pipeline.StructuredData, error)

type funcRegister map[string]transformFunc

var FunctionRegister = funcRegister{}

func RegisterTransformFunction(name string, processMapFunction transformFunc) {
	FunctionRegister[name] = processMapFunction
}

func TransformWith(transformArg string, inputMap pipeline.StructuredData) (pipeline.StructuredData, error) {
	transformName, parameter := pipeline.SplitArg(transformArg)

	f, err := lookupTransformFunc(transformName)
	if err != nil {
		return nil, err
	}

	return f(inputMap, parameter)
}

func lookupTransformFunc(transformName string) (transformFunc, error) {
	if f, ok := FunctionRegister[transformName]; ok {
		return f, nil
	}
	return nil, errors.New("Unknown transform: " + transformName)
}
