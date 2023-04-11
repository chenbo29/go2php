package util

import (
	"bytes"
	"encoding/json"
)

// PrettyJsonStr pretty json string
func PrettyJsonStr(jsonString string) (string, error) {
	var prettyJSON bytes.Buffer
	err := json.Indent(&prettyJSON, []byte(jsonString), "", "  ")
	if err != nil {
		return "", err
	}
	return string(prettyJSON.Bytes()), nil
}
