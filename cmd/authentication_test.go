package cmd

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"path/filepath"
	"strings"
	"testing"

	openapi "github.com/sdslabs/gctl/client"
)

var generatedToken string
var userdata openapi.User

func Test_RegisterCmd(t *testing.T) {
	g, _ := ioutil.ReadFile(filepath.Join("testdata", "userdata.json"))
	if err := json.Unmarshal(g, &userdata); err != nil {
		t.Fatal("Error in reading user details from json", err)
	}
	newRegisterCmd := RegisterCmd(client)
	b := bytes.NewBufferString("")
	newRegisterCmd.SetOut(b)
	newRegisterCmd.SetArgs([]string{"-u", userdata.Username, "-e", userdata.Email, "-p", userdata.Password})
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
	newLoginCmd.SetArgs([]string{"-e", userdata.Email, "-p", userdata.Password})
	newLoginCmd.Execute()
	out, err := ioutil.ReadAll(b)
	if err != nil {
		t.Fatal("Error in reading output")
	}
	if bytes.Contains(out, []byte("Error")) {
		t.Fatalf("User cannot be logged in.")
	} else {
		generatedToken = strings.Split(string(out), " ")[2]
	}
}
