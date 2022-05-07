package dump

import (
	"encoding/json"
)

func init() {
	Register("json", DumpToJSON)
}

const jsonPrefix = ""
const jsonIndent = "  "

func DumpToJSON(inputMap interface{}) (string, error) {
	j, err := json.MarshalIndent(inputMap, jsonPrefix, jsonIndent)
	if err != nil {
		return "", err
	}
	return string(j), nil
}
