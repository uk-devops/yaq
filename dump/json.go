package dump

import (
	"encoding/json"

	"github.com/saliceti/yaq/pipeline"
)

func init() {
	Register("json", DumpToJSON)
}

const jsonPrefix = ""
const jsonIndent = "  "

func DumpToJSON(inputMap pipeline.GenericMap) (string, error) {
	j, err := json.MarshalIndent(inputMap, jsonPrefix, jsonIndent)
	if err != nil {
		return "", err
	}
	return string(j), nil
}
