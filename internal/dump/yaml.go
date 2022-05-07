package dump

import (
	"github.com/uk-devops/yaq/internal/pipeline"
	"gopkg.in/yaml.v3"
)

func init() {
	Register("yaml", DumpToYAML)
}

func DumpToYAML(inputMap pipeline.StructuredData) (string, error) {
	j, err := yaml.Marshal(inputMap)
	if err != nil {
		return "", err
	}
	return string(j), nil
}
