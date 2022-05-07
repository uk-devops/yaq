package load

import (
	"github.com/saliceti/yaq/pipeline"
	"gopkg.in/yaml.v3"
)

func init() {
	Register("yaml", loadFromYAML)
}

func loadFromYAML(inString string) (pipeline.GenericMap, error) {
	out := make(pipeline.GenericMap)
	err := yaml.Unmarshal([]byte(inString), &out)

	return out, err
}
