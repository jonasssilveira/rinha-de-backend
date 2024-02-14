package errors

import (
	"encoding/json"
)

type ClientNotFound struct {
	errorMessage    string
	errorStatusCode int
}

func (n *ClientNotFound) Error() string {
	marshal, err := json.Marshal(n)

	if err != nil {
		return err.Error()
	}

	return string(marshal)
}
