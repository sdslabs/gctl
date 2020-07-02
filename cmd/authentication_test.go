package cmd

import (
	"bytes"
	"io/ioutil"
	"testing"
)

func Test_RegisterCmd(t *testing.T) {
	newRegisterCmd := RegisterCmd(client)
	b := bytes.NewBufferString("")
	newRegisterCmd.SetOut(b)
	newRegisterCmd.SetArgs([]string{"-u", "gctltest", "-e", "gctltest@test.com", "-p", "gctltest"})
	newRegisterCmd.Execute()
	out, err := ioutil.ReadAll(b)
	if err != nil {
		t.Fatal("Error in reading output")
	}
	if !bytes.Contains(out, []byte("user created")) {
		t.Fatalf("User cannot be created.")
	}
}

func Test_LoginCmd(t *testing.T) {
	newLoginCmd := LoginCmd(client)
	b := bytes.NewBufferString("")
	newLoginCmd.SetOut(b)
	newLoginCmd.SetArgs([]string{"-e", "gctltest@test.com", "-p", "gctltest"})
	newLoginCmd.Execute()
	out, err := ioutil.ReadAll(b)
	if err != nil {
		t.Fatal("Error in reading output")
	}
	if bytes.Contains(out, []byte("Error")) {
		t.Fatalf(string(out), "User cannot be logged in.")
	}
}
