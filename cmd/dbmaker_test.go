package cmd

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"path/filepath"
	"testing"

	openapi "github.com/sdslabs/gctl/client"
)

var dbdata openapi.Database

func Test_CreateDbCmd(t *testing.T) {
	g, _ := ioutil.ReadFile(filepath.Join("testdata", "dbdata.json"))
	if err := json.Unmarshal(g, &dbdata); err != nil {
		t.Fatal("Error in reading db data from json", err)
	}
	newDbCmd := CreateDbCmd(client)
	b := bytes.NewBufferString("")
	newDbCmd.SetOut(b)
	newDbCmd.SetArgs([]string{tokenTest, "-n", dbdata.Name, "-p", dbdata.Password, "-t", "mysql"})
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
	fetchSingleDb.SetArgs([]string{tokenTest, "-n", dbdata.Name})
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
	deleteDbCmd.SetArgs([]string{dbdata.Password, tokenTest})
	deleteDbCmd.Execute()
	out, err := ioutil.ReadAll(b)
	if err != nil {
		t.Fatal("Error in reading output")
	}
	if !bytes.Contains(out, []byte("Database deleted successfully")) {
		t.Fatalf("Error in deleting database")
	}
}
