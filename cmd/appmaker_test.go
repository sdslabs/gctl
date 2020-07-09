package cmd

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"

	openapi "github.com/sdslabs/gctl/client"
)

var tokenTest string
var appdata openapi.Application

func TestMain(m *testing.M) {
	var loginCreds openapi.Login
	g, _ := ioutil.ReadFile(filepath.Join("testdata", "logincreds.json"))
	json.Unmarshal(g, &loginCreds)
	loginCmd := LoginCmd(client)
	b := bytes.NewBufferString("")
	loginCmd.SetOut(b)
	loginCmd.SetArgs([]string{"-e", loginCreds.Email, "-p", loginCreds.Password})
	loginCmd.Execute()
	out, _ := ioutil.ReadAll(b)
	tokenTest = strings.Split(string(out), " ")[2]
	os.Exit(m.Run())
}

func Test_CreateAppCmd(t *testing.T) {
	g, _ := ioutil.ReadFile(filepath.Join("testdata", "apptest.json"))
	if err := json.Unmarshal(g, &appdata); err != nil {
		t.Fatal("Error in reading app data from json file", err)
	}
	newAppCmd := CreateAppCmd(*client)
	b := bytes.NewBufferString("")
	newAppCmd.SetOut(b)
	newAppCmd.SetArgs([]string{tokenTest, "/testdata/apptest.json", "nodejs"})
	newAppCmd.Execute()
	out, err := ioutil.ReadAll(b)
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Contains(out, []byte("App created successfully")) {
		t.Fatalf("App cannot be created.")
	}
}

func Test_FetchAppCmd(t *testing.T) {
	fetchSingleApp := FetchAppCmd(*client)
	b := bytes.NewBufferString("")
	fetchSingleApp.SetOut(b)
	fetchSingleApp.SetArgs([]string{tokenTest, "-n", appdata.Name})
	fetchSingleApp.Execute()
	out, err := ioutil.ReadAll(b)
	if err != nil {
		t.Fatal(err)
	}
	if bytes.Contains(out, []byte("Error in fetching the app.")) {
		t.Fatalf("Single app cannot be fetched")
	}
	fetchAllApps := FetchAppCmd(*client)
	bnew := bytes.NewBufferString("")
	fetchAllApps.SetOut(bnew)
	fetchAllApps.SetArgs([]string{tokenTest})
	fetchAllApps.Execute()
	out, err = ioutil.ReadAll(b)
	if err != nil {
		t.Fatal(err)
	}
	if bytes.Contains(out, []byte("Error in fetching the apps.")) {
		t.Fatalf("All apps cannot be fetched")
	}
}

func Test_DeleteAppCmd(t *testing.T) {
	newDeleteApp := DeleteAppCmd(*client)
	b := bytes.NewBufferString("")
	newDeleteApp.SetOut(b)
	newDeleteApp.SetArgs([]string{appdata.Name, tokenTest})
	newDeleteApp.Execute()
	out, err := ioutil.ReadAll(b)
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Contains(out, []byte("App deleted successfully")) {
		t.Fatalf("App cannot be deleted.")
	}
}
