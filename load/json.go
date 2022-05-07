package load

import (
	"encoding/json"

	"github.com/saliceti/yaq/pipeline"
)

func init() {
	Register("json", LoadFromJSON)
}

func LoadFromJSON(inString string) (pipeline.GenericMap, error) {
	out := make(pipeline.GenericMap)
	err := json.Unmarshal([]byte(inString), &out)

	return out, err
}
