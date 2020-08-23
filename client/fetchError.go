package openapi

import (
	"encoding/json"
)

func fetchError(b []byte) string {
	var res InlineResponse400
	if err := json.Unmarshal(b, &res); err != nil {
		return "Error in fetching errors"
	}
	return res.Error
}
