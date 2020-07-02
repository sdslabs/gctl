package cmd

import (
	"bytes"
	"io/ioutil"
	"testing"
)

var tokenTest = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhZG1pbiI6ZmFsc2UsImVtYWlsIjoiZ21haGFrMUBnbWFpbC5jb20iLCJleHAiOjE1OTM3MjI4NDUsIm9yaWdfaWF0IjoxNTkzNzE5MjQ1LCJ1c2VybmFtZSI6Im1haGFrIn0.Rb5TxokNX73WGgXErhyHlKEfdREwnp1snxkMXMmMn24"

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
