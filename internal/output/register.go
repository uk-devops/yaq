package output

import (
	"errors"

	"github.com/saliceti/yaq/internal/pipeline"
)

type stringOutputFunctionType func(string, string) error
type mapOutputFunctionType func(pipeline.StructuredData, string, []string) error

type outputFunc struct {
	stringOutputFunction stringOutputFunctionType
	mapOutputFunction    mapOutputFunctionType
}
type outputFuncRegister map[string]outputFunc

var register = outputFuncRegister{}

func (f outputFunc) isMapFunc() bool {
	return f.mapOutputFunction != nil
}

func RegisterStringFunction(name string, outputFunction stringOutputFunctionType) {
	register[name] = outputFunc{stringOutputFunction: outputFunction}
}
func RegisterMapFunction(name string, outputFunction mapOutputFunctionType) {
	register[name] = outputFunc{mapOutputFunction: outputFunction}
}

func PushString(outputArg string, outputString string) error {
	outputName, parameter := pipeline.SplitArg(outputArg)

	f, err := lookupOutputFunction(outputName)
	if err != nil {
		return err
	}

	return f.stringOutputFunction(outputString, parameter)
}

func PushMap(outputArg string, outputData pipeline.StructuredData, extra []string) error {
	outputName, parameter := pipeline.SplitArg(outputArg)

	f, err := lookupOutputFunction(outputName)
	if err != nil {
		return err
	}

	return f.mapOutputFunction(outputData, parameter, extra)
}

func RequiresMap(outputArg string) (bool, error) {
	outputName, _ := pipeline.SplitArg(outputArg)

	f, err := lookupOutputFunction(outputName)
	if err != nil {
		return false, err
	}

	return f.isMapFunc(), nil
}

func lookupOutputFunction(outputName string) (*outputFunc, error) {
	if f, ok := register[outputName]; ok {
		return &f, nil
	}
	return nil, errors.New("Unknown output: " + outputName)
}
