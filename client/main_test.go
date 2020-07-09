package openapi

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

var (
	cfg    *Configuration
	client *APIClient
	token  string
)

func TestMain(m *testing.M) {
	var loginCreds Login
	cfg = NewConfiguration()
	client = NewAPIClient(cfg)
	g, _ := ioutil.ReadFile(filepath.Join("testdata", "logincreds.json"))
	json.Unmarshal(g, &loginCreds)
	response, _, _ := client.AuthApi.Login(context.Background(), loginCreds)
	token = response.Token
	os.Exit(m.Run())
}
