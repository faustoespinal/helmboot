package utils

import (
	"bytes"
	"encoding/json"
)

const (
	empty = ""
	tab   = "    "
)

// PrettyJSON returns a string of nicely formatted json.
func PrettyJSON(data interface{}) (string, error) {
	buffer := new(bytes.Buffer)
	encoder := json.NewEncoder(buffer)
	encoder.SetIndent(empty, tab)

	err := encoder.Encode(data)
	if err != nil {
		return empty, err
	}
	return buffer.String(), nil
}
