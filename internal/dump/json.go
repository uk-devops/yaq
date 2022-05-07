package dump

import (
	"encoding/json"

	"github.com/saliceti/yaq/internal/pipeline"
)

func init() {
	Register("json", DumpToJSON)
}

const jsonPrefix = ""
const jsonIndent = "  "

func DumpToJSON(inputMap pipeline.StructuredData) (string, error) {
	j, err := json.MarshalIndent(inputMap, jsonPrefix, jsonIndent)
	if err != nil {
		return "", err
	}
	return string(j), nil
}
