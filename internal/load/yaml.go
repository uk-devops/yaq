package load

import (
	"github.com/saliceti/yaq/internal/pipeline"
	"gopkg.in/yaml.v3"
)

func init() {
	Register("yaml", UnmarshalYAML)
}

func unmarshalMapFromYAML(inString string) (pipeline.GenericMap, error) {
	var out pipeline.GenericMap
	err := yaml.Unmarshal([]byte(inString), &out)

	return out, err
}

func unmarshalArrayFromYAML(inString string) (pipeline.GenericArray, error) {
	var out pipeline.GenericArray
	err := yaml.Unmarshal([]byte(inString), &out)

	return out, err
}

func UnmarshalYAML(inString string) (pipeline.StructuredData, error) {
	var d pipeline.StructuredData
	var err error

	d, err = unmarshalMapFromYAML(inString)
	if err != nil {
		d, err = unmarshalArrayFromYAML(inString)
	}

	if err != nil {
		return d, err
	}

	return d, nil
}
