package pkg

import (
	"encoding/json"
)

type FiberError struct {
	Error       error
	Description string
}

func (fe FiberError) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Error       string `json:"error"`
		Description string `json:"description"`
	}{
		Error:       fe.Error.Error(),
		Description: fe.Description,
	})
}
