package openapi

import (
	"context"
	"testing"
)

func TestInstancesApi(t *testing.T) {
	auth := context.WithValue(context.Background(), ContextAccessToken, token)
	response, _, err := client.InstancesApi.FetchIntancesByUser(auth)
	if !response.Success {
		t.Fatal("Error in fetching instances", err)
	}
}
