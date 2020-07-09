package openapi

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"path/filepath"
	"testing"
)

func TestUserApi(t *testing.T) {
	var passdetails InlineObject
	g, err := ioutil.ReadFile(filepath.Join("testdata", "updatepassdata.json"))
	json.Unmarshal(g, &passdetails)
	auth := context.WithValue(context.Background(), ContextAccessToken, testToken)
	fetchRes, _, err := client.UserApi.FetchUser(auth)
	if !fetchRes.Success {
		t.Fatal("Error in fetching user details", err)
	}
	updatePass, _, err := client.UserApi.UpdatePassword(auth, passdetails)
	if !updatePass.Success {
		t.Fatal("Error in updating password", err)
	}
	deleteUserRes, _, err := client.UserApi.DeleteUser(auth)
	if !deleteUserRes.Success {
		t.Fatal("Error in deleting the user", err)
	}
}
