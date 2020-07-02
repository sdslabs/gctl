package cmd

import (
	"bytes"
	"io/ioutil"
	"strings"
	"testing"
)

var tokenTest string

func Test_GenerateToken(t *testing.T) {
	loginCmd := LoginCmd(client)
	b := bytes.NewBufferString("")
	loginCmd.SetOut(b)
	loginCmd.SetArgs([]string{"-e", "anish.mukherjee1996@gmail.com", "-p", "alphadose"})
	loginCmd.Execute()
	out, _ := ioutil.ReadAll(b)
	tokenTest = strings.Split(string(out), " ")[2]
}

func Test_CreateAppCmd(t *testing.T) {
	newAppCmd := CreateAppCmd(*client)
	b := bytes.NewBufferString("")
	newAppCmd.SetOut(b)
	newAppCmd.SetArgs([]string{tokenTest, "apptest.json", "nodejs"})
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
	fetchSingleApp.SetArgs([]string{tokenTest, "-n", "gctltestapp"})
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
	newDeleteApp.SetArgs([]string{"gctltestapp", tokenTest})
	newDeleteApp.Execute()
	out, err := ioutil.ReadAll(b)
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Contains(out, []byte("App deleted successfully")) {
		t.Fatalf("App cannot be deleted.")
	}
}
