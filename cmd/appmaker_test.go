package cmd

import (
	"bytes"
	"io/ioutil"
	"testing"
)

var token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhZG1pbiI6ZmFsc2UsImVtYWlsIjoiZ21haGFrMUBnbWFpbC5jb20iLCJleHAiOjE1OTM" +
	"3MTI3MTAsIm9yaWdfaWF0IjoxNTkzNzA5MTEwLCJ1c2VybmFtZSI6Im1haGFrIn0.9nViWnWUabyGJb_o09DovT3GiI2S-PNb8eZuCYpF3EY"

func Test_CreateAppCmd(t *testing.T) {
	newAppCmd := CreateAppCmd(*client)
	b := bytes.NewBufferString("")
	newAppCmd.SetOut(b)
	newAppCmd.SetArgs([]string{token, "apptest.json", "nodejs"})
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
	fetchSingleApp.SetArgs([]string{token, "-n", "gctltestapp"})
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
	fetchAllApps.SetArgs([]string{token})
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
	newDeleteApp.SetArgs([]string{"gctltestapp", token})
	newDeleteApp.Execute()
	out, err := ioutil.ReadAll(b)
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Contains(out, []byte("App deleted successfully")) {
		t.Fatalf("App cannot be deleted.")
	}
}
