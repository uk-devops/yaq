package load

import (
	"encoding/json"
	"github.com/saliceti/yaq/pipeline"
)

func init() {
	Register("json", UnmarshalJSON)
}

func unmarshalMapFromJSON(inString string) (pipeline.GenericMap, error) {
	var out pipeline.GenericMap
	err := json.Unmarshal([]byte(inString), &out)

	return out, err
}

func unmarshalArrayFromJSON(inString string) (pipeline.GenericArray, error) {
	var out pipeline.GenericArray
	err := json.Unmarshal([]byte(inString), &out)

	return out, err
}

func UnmarshalJSON(inString string) (interface{}, error) {
	var d interface{}
	var err error

	d, err = unmarshalMapFromJSON(inString)
	if err != nil {
		d, err = unmarshalArrayFromJSON(inString)
	}

	if err != nil {
		return d, err
	}

	return d, nil
}
