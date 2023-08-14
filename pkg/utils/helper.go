package utils

import (
	"encoding/json"
	"io"
	"strings"
)

func ConvertObjToStringReader[T any](obj T) io.Reader {
	result, _ := json.Marshal(&obj)
	return strings.NewReader(string(result))
}
