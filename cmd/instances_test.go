package cmd

import (
	"bytes"
	"io/ioutil"
	"testing"
)

func Test_GetInstancesCmd(t *testing.T) {
	newGetInstancesCmd := GetInstancesCmd(client)
	b := bytes.NewBufferString("")
	newGetInstancesCmd.SetOut(b)
	newGetInstancesCmd.SetArgs([]string{tokenTest})
	newGetInstancesCmd.Execute()
	out, err := ioutil.ReadAll(b)
	if err != nil {
		t.Fatal("Error in reading output")
	}
	if bytes.Contains(out, []byte("Error")) {
		t.Fatalf(string(out), "All instances cannot be fetched")
	}
}
