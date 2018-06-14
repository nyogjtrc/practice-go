package imgur

import (
	"encoding/json"
	"errors"
)

// APIError API error message
type APIError struct {
	Error   string `json:"error"`
	Request string `json:"request"`
	Method  string `json:"method"`
}

// ResponseError response message of API error
type ResponseError struct {
	Data    APIError `json:"data"`
	Success bool     `json:"success"`
	Status  int      `json:"status"`
}

func parseError(body []byte) (err error) {
	r := new(ResponseError)
	err = json.Unmarshal(body, r)
	if err != nil {
		return
	}
	return errors.New(r.Data.Error)
}
