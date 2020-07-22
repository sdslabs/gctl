package openapi

import (
	"encoding/json"
	"fmt"
)

func fetchError(b []byte) string {
	var res InlineResponse400
	if err := json.Unmarshal(b, &res); err != nil {
		fmt.Print(err)
	}
	return res.Error
}
