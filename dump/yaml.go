package dump

import (
	"gopkg.in/yaml.v3"
)

func init() {
	Register("yaml", DumpToYAML)
}

func DumpToYAML(inputMap interface{}) (string, error) {
	j, err := yaml.Marshal(inputMap)
	if err != nil {
		return "", err
	}
	return string(j), nil
}
