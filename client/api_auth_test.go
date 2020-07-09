package openapi

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"path/filepath"
	"testing"
)

var testToken string

func TestAuthApi(t *testing.T) {
	var userdata User
	var loginCreds Login
	g, _ := ioutil.ReadFile(filepath.Join("testdata", "userdata.json"))
	json.Unmarshal(g, &userdata)
	response, _, err := client.AuthApi.Register(context.Background(), userdata)
	if !response.Success {
		t.Fatal("Error in registering user", err)
	}
	loginCreds.Email = userdata.Email
	loginCreds.Password = userdata.Password
	loginres, _, err := client.AuthApi.Login(context.Background(), loginCreds)
	if loginres.Code != 200 {
		t.Fatal("Error in logging in", err)
	}
	refreshRes, _, err := client.AuthApi.Refresh(context.Background(), "Bearer "+loginres.Token)
	if refreshRes.Code != 200 {
		t.Fatal("Error in refreshing the token", err)
	} else {
		testToken = refreshRes.Token
	}
}
