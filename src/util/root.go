package util

import (
	"bytes"
	"encoding/json"
)

func PrettyJson(str string) string {
	var buf bytes.Buffer
	if err := json.Indent(&buf, []byte(str), "", "    "); err != nil {
		return ""
	}
	return buf.String()
}

func PrettyJsonBody(str string) string {
	var buf bytes.Buffer
	if err := json.Indent(&buf, []byte(str), "", "    "); err != nil {
		return ""
	}
	return buf.String()
}
