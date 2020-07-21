package cmd

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	_nethttp "net/http"
	"path/filepath"
	"testing"

	"github.com/antihax/optional"
	"github.com/golang/mock/gomock"
	openapi "github.com/sdslabs/gctl/client"
	"github.com/sdslabs/gctl/cmd/testmocks"
)

var dbdata openapi.Database

func Test_CreateDbCmd(t *testing.T) {
	var httpres *_nethttp.Response
	g, _ := ioutil.ReadFile(filepath.Join("testdata", "dbdata.json"))
	if err := json.Unmarshal(g, &dbdata); err != nil {
		t.Fatal("Error in reading db data from json", err)
	}
	input := &openapi.CreateDBOpts{
		Database: optional.NewInterface(dbdata),
	}
	token, _ := ioutil.ReadFile(filepath.Join("testdata", "token.txt"))
	gctltoken = string(token)
	defer LogoutCmd().Execute()
	auth1 := context.WithValue(context.Background(), openapi.ContextAccessToken, string(token))

	ctrl := gomock.NewController(t)
	mockDb := testmocks.NewMockDbsAPIService(ctrl)
	createDbRes1 := openapi.InlineResponse2002{Success: true}
	createDbRes2 := openapi.InlineResponse2002{Success: false}
	mockDb.EXPECT().CreateDB(auth1, "mysql", input).Return(createDbRes1, httpres, nil).AnyTimes()
	mockDb.EXPECT().CreateDB(auth1, "", input).Return(createDbRes2, httpres, nil).AnyTimes()

	newDbCmd := CreateDbCmd(mockDb)
	b := bytes.NewBufferString("")
	newDbCmd.SetOut(b)

	newDbCmd.SetArgs([]string{"-n", dbdata.Name, "-p", dbdata.Password, "-t", "mysql"})
	newDbCmd.Execute()
	out, err := ioutil.ReadAll(b)
	if err != nil {
		t.Fatal("Error in reading output")
	}
	if !bytes.Contains(out, []byte("Database created")) {
		t.Fatalf("Error in creating database")
	}

	b = bytes.NewBufferString("")
	newDbCmd.SetArgs([]string{"-n", dbdata.Name, "-p", dbdata.Password})
	newDbCmd.Execute()
	out, err = ioutil.ReadAll(b)
	if err != nil {
		t.Fatal("Error in reading output")
	}
	if bytes.Contains(out, []byte("Database created")) {
		t.Fatalf("Database created without db type")
	}
}

func Test_FetchDbCmd(t *testing.T) {
	var httpres *_nethttp.Response
	token, _ := ioutil.ReadFile(filepath.Join("testdata", "token.txt"))
	gctltoken = string(token)
	defer LogoutCmd().Execute()
	auth1 := context.WithValue(context.Background(), openapi.ContextAccessToken, string(token))

	ctrl := gomock.NewController(t)
	mockDb := testmocks.NewMockDbsAPIService(ctrl)
	output1 := openapi.InlineResponse2007{Success: true}
	mockDb.EXPECT().FetchDbByUser(auth1, dbdata.Name).Return(output1, httpres, nil).AnyTimes()
	mockDb.EXPECT().FetchDbsByUser(auth1).Return(output1, httpres, nil).AnyTimes()

	fetchSingleDb := FetchDbCmd(mockDb)
	b := bytes.NewBufferString("")
	fetchSingleDb.SetOut(b)

	fetchSingleDb.SetArgs([]string{"-n", dbdata.Name})
	fetchSingleDb.Execute()
	out, err := ioutil.ReadAll(b)
	if err != nil {
		t.Fatal("Error in reading output")
	}
	if bytes.Contains(out, []byte("Error in fetching the database")) {
		t.Fatalf("Error in fetching single database")
	}

	fetchAllDb := FetchDbCmd(mockDb)
	bnew := bytes.NewBufferString("")
	fetchAllDb.SetOut(bnew)

	fetchAllDb.Execute()
	out, err = ioutil.ReadAll(bnew)
	if err != nil {
		t.Fatal("Error in reading output")
	}
	if bytes.Contains(out, []byte("Error in fetching the databases")) {
		t.Fatalf("Error in fetching all databases")
	}

}

func Test_DeleteDbCmd(t *testing.T) {
	var httpres *_nethttp.Response
	token, _ := ioutil.ReadFile(filepath.Join("testdata", "token.txt"))
	gctltoken = string(token)
	defer LogoutCmd().Execute()
	auth1 := context.WithValue(context.Background(), openapi.ContextAccessToken, string(token))

	ctrl := gomock.NewController(t)
	mockDb := testmocks.NewMockDbsAPIService(ctrl)
	output1 := openapi.InlineResponse2002{Success: true}
	output2 := openapi.InlineResponse2002{Success: false}
	mockDb.EXPECT().DeleteDbByUser(auth1, dbdata.Name).Return(output1, httpres, nil).AnyTimes()
	mockDb.EXPECT().DeleteDbByUser(auth1, "").Return(output2, httpres, nil).AnyTimes()

	deleteDbCmd := DeleteDbCmd(mockDb)
	b := bytes.NewBufferString("")
	deleteDbCmd.SetOut(b)

	deleteDbCmd.SetArgs([]string{dbdata.Name})
	deleteDbCmd.Execute()
	out, err := ioutil.ReadAll(b)
	if err != nil {
		t.Fatal("Error in reading output")
	}
	if !bytes.Contains(out, []byte("Database deleted successfully")) {
		t.Fatalf("Error in deleting database")
	}

	deleteDbCmd.SetArgs([]string{""})
	deleteDbCmd.Execute()
	out, err = ioutil.ReadAll(b)
	if err != nil {
		t.Fatal("Error in reading output")
	}
	if bytes.Contains(out, []byte("Database deleted successfully")) {
		t.Fatalf("Database deleted without db name")
	}
}
