package dump

import (
	"github.com/saliceti/yaq/pipeline"
	"gopkg.in/yaml.v3"
)

func init() {
	Register("yaml", DumpToYAML)
}

func DumpToYAML(inputMap pipeline.GenericMap) (string, error) {
	j, err := yaml.Marshal(inputMap)
	if err != nil {
		return "", err
	}
	return string(j), nil
}
