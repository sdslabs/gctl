package cmd

import (
	"bytes"
	"io/ioutil"
	"testing"
)

func Test_CreateDbCmd(t *testing.T) {
	newDbCmd := CreateDbCmd(client)
	b := bytes.NewBufferString("")
	newDbCmd.SetOut(b)
	newDbCmd.SetArgs([]string{tokenTest, "-n", "gctltestdb", "-p", "gctltestdb", "-t", "mysql"})
	newDbCmd.Execute()
	out, err := ioutil.ReadAll(b)
	if err != nil {
		t.Fatal("Error in reading output")
	}
	if !bytes.Contains(out, []byte("Database created")) {
		t.Fatalf("Error in creating database")
	}
}

func Test_FetchDbCmd(t *testing.T) {
	fetchSingleDb := FetchDbCmd(client)
	b := bytes.NewBufferString("")
	fetchSingleDb.SetOut(b)
	fetchSingleDb.SetArgs([]string{tokenTest, "-n", "gctltestdb"})
	fetchSingleDb.Execute()
	out, err := ioutil.ReadAll(b)
	if err != nil {
		t.Fatal("Error in reading output")
	}
	if bytes.Contains(out, []byte("Error in fetching the database")) {
		t.Fatalf("Error in fetching single database")
	}
	fetchAllDb := FetchDbCmd(client)
	bnew := bytes.NewBufferString("")
	fetchAllDb.SetOut(bnew)
	fetchAllDb.SetArgs([]string{tokenTest})
	fetchAllDb.Execute()
	out, err = ioutil.ReadAll(b)
	if err != nil {
		t.Fatal("Error in reading output")
	}
	if bytes.Contains(out, []byte("Error in fetching the databases")) {
		t.Fatalf("Error in fetching all databases")
	}
}

func Test_DeleteDbCmd(t *testing.T) {
	deleteDbCmd := DeleteDbCmd(client)
	b := bytes.NewBufferString("")
	deleteDbCmd.SetOut(b)
	deleteDbCmd.SetArgs([]string{"gctltestdb", tokenTest})
	deleteDbCmd.Execute()
	out, err := ioutil.ReadAll(b)
	if err != nil {
		t.Fatal("Error in reading output")
	}
	if !bytes.Contains(out, []byte("Database deleted successfully")) {
		t.Fatalf("Error in deleting database")
	}
}
