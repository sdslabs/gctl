package cmd

import (
	"bytes"
	"io/ioutil"
	"testing"
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
	newUpdateCmd := UpdateUserPasswdCmd(client)
	b := bytes.NewBufferString("")
	newUpdateCmd.SetOut(b)
	newUpdateCmd.SetArgs([]string{generatedToken, "-o", "gctltest", "-n", "gctltest"})
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
