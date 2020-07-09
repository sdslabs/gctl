package cmd

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"path/filepath"
	"testing"

	openapi "github.com/sdslabs/gctl/client"
)

func Test_FetchUserCmd(t *testing.T) {
	newFetchUser := FetchUserCmd(client)
	b := bytes.NewBufferString("")
	newFetchUser.SetOut(b)
	newFetchUser.SetArgs([]string{generatedToken})
	newFetchUser.Execute()
	out, err := ioutil.ReadAll(b)
	if err != nil {
		t.Fatal("Error in reading output")
	}
	if bytes.Contains(out, []byte("Error")) {
		t.Fatalf("User details cannot be fetched")
	}
}

func Test_UpdateUserCmd(t *testing.T) {
	var updatePassDetails openapi.InlineObject
	g, _ := ioutil.ReadFile(filepath.Join("testdata", "updatepassdata.json"))
	if err := json.Unmarshal(g, &updatePassDetails); err != nil {
		t.Fatal("Error in reading update password details from json", err)
	}
	newUpdateCmd := UpdateUserPasswdCmd(client)
	b := bytes.NewBufferString("")
	newUpdateCmd.SetOut(b)
	newUpdateCmd.SetArgs([]string{generatedToken, "-o", updatePassDetails.OldPassword, "-n", updatePassDetails.NewPassword})
	newUpdateCmd.Execute()
	out, err := ioutil.ReadAll(b)
	if err != nil {
		t.Fatal("Error in reading output")
	}
	if bytes.Contains(out, []byte("Error")) {
		t.Fatalf("User password cannot be updated")
	}
}

func Test_DeleteUserCmd(t *testing.T) {
	newDeleteUserCmd := DeleteUserCmd(client)
	b := bytes.NewBufferString("")
	newDeleteUserCmd.SetOut(b)
	newDeleteUserCmd.SetArgs([]string{generatedToken})
	newDeleteUserCmd.Execute()
	out, err := ioutil.ReadAll(b)
	if err != nil {
		t.Fatal("Error in reading output")
	}
	if bytes.Contains(out, []byte("Error")) {
		t.Fatalf("User cannot be deleted")
	}
}
