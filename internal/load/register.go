package load

import (
	"errors"

	"github.com/uk-devops/yaq/internal/pipeline"
)

type loadFunc func(string) (pipeline.StructuredData, error)
type loadFuncRegister map[string]loadFunc

var FunctionRegister = loadFuncRegister{}

func Register(name string, loadFunction loadFunc) {
	FunctionRegister[name] = loadFunction
}

func Unmarshal(inputString string) (pipeline.StructuredData, error) {
	var d pipeline.StructuredData
	var err error

	d, err = FunctionRegister["json"](inputString)
	if err != nil {
		d, err = FunctionRegister["yaml"](inputString)
	}

	if err != nil {
		return d, errors.New("Invalid json or yaml:\n" + err.Error())
	}

	return d, err
}
