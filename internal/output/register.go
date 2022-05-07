package output

import (
	"errors"
	"strings"

	"github.com/saliceti/yaq/internal/pipeline"
)

type stringOutputFunctionType func(string, string) error
type mapOutputFunctionType func(pipeline.StructuredData, []string) error

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
	outputArgArray := strings.Split(outputArg, ":")

	if f, ok := register[outputArgArray[0]]; ok {
		var parameter string
		if len(outputArgArray) == 1 {
			parameter = ""
		} else {
			parameter = outputArgArray[1]
		}

		err := f.stringOutputFunction(outputString, parameter)
		return err
	}
	return errors.New("Unknown output: " + outputArg)
}

func PushMap(outputArg string, outputData pipeline.StructuredData, extra []string) error {
	if f, ok := register[outputArg]; ok {
		err := f.mapOutputFunction(outputData, extra)
		return err
	}
	return errors.New("Unknown output: " + outputArg)
}

func RequiresMap(outputArg string) bool {
	return register[outputArg].isMapFunc()
}
