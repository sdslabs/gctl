package cmd

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	_nethttp "net/http"
	"path/filepath"
	"testing"

	"github.com/golang/mock/gomock"
	openapi "github.com/sdslabs/gctl/client"
	"github.com/sdslabs/gctl/cmd/testmocks"
)

var appdata openapi.Application

func Test_CreateAppCmd(t *testing.T) {
	var httpres *_nethttp.Response
	token, _ := ioutil.ReadFile(filepath.Join("testdata", "token.txt"))
	gctltoken = string(token)
	defer LogoutCmd().Execute()
	auth1 := context.WithValue(context.Background(), openapi.ContextAccessToken, string(token))
	g, _ := ioutil.ReadFile(filepath.Join("testdata", "apptest.json"))
	if err := json.Unmarshal(g, &appdata); err != nil {
		t.Fatal("Error in reading app data from json file", err)
	}

	ctrl := gomock.NewController(t)
	mockApp := testmocks.NewMockAppsAPIService(ctrl)
	output1 := openapi.InlineResponse2002{Success: true}
	output2 := openapi.InlineResponse2002{Success: false}
	output3 := openapi.InlineResponse2003{Success: true, Data: []openapi.CreatedApplication{{Id: "1"}}}
	mockApp.EXPECT().CreateApp(auth1, "nodejs", appdata).Return(output1, httpres, nil).AnyTimes()
	mockApp.EXPECT().CreateApp(auth1, "", appdata).Return(output2, httpres, nil).AnyTimes()
	mockApp.EXPECT().FetchAppByUser(auth1, appdata.Name).Return(output3, httpres, nil).AnyTimes()

	newAppCmd := CreateAppCmd(mockApp)
	b := bytes.NewBufferString("")
	newAppCmd.SetOut(b)

	newAppCmd.SetArgs([]string{"/testdata/apptest.json", "nodejs"})
	newAppCmd.Execute()
	out, err := ioutil.ReadAll(b)
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Contains(out, []byte("App created successfully")) {
		t.Fatal("App cannot be created.")
	}

	newAppCmd.SetArgs([]string{"/testdata/apptest.json", ""})
	newAppCmd.Execute()
	out, err = ioutil.ReadAll(b)
	if err != nil {
		t.Fatal(err)
	}
	if bytes.Contains(out, []byte("App created successfully")) {
		t.Fatalf("App created without app language.")
	}
}

func Test_FetchAppCmd(t *testing.T) {
	var httpres *_nethttp.Response
	token, _ := ioutil.ReadFile(filepath.Join("testdata", "token.txt"))
	gctltoken = string(token)
	defer LogoutCmd().Execute()
	auth1 := context.WithValue(context.Background(), openapi.ContextAccessToken, string(token))

	ctrl := gomock.NewController(t)
	mockApp := testmocks.NewMockAppsAPIService(ctrl)
	output1 := openapi.InlineResponse2003{Success: true}
	mockApp.EXPECT().FetchAppByUser(auth1, appdata.Name).Return(output1, httpres, nil).AnyTimes()
	mockApp.EXPECT().FetchAppsByUser(auth1).Return(output1, httpres, nil).AnyTimes()

	fetchSingleApp := FetchAppCmd(mockApp)
	b := bytes.NewBufferString("")
	fetchSingleApp.SetOut(b)

	fetchSingleApp.SetArgs([]string{"-n", appdata.Name})
	fetchSingleApp.Execute()
	out, err := ioutil.ReadAll(b)
	if err != nil {
		t.Fatal(err)
	}
	if bytes.Contains(out, []byte("Error in fetching the app.")) {
		t.Fatalf("Single app cannot be fetched")
	}

	fetchAllApps := FetchAppCmd(mockApp)
	bnew := bytes.NewBufferString("")
	fetchAllApps.SetOut(bnew)

	fetchAllApps.Execute()
	out, err = ioutil.ReadAll(bnew)
	if err != nil {
		t.Fatal(err)
	}
	if bytes.Contains(out, []byte("Error in fetching the apps.")) {
		t.Fatalf("All apps cannot be fetched")
	}
}

func Test_DeleteAppCmd(t *testing.T) {
	var httpres *_nethttp.Response
	token, _ := ioutil.ReadFile(filepath.Join("testdata", "token.txt"))
	gctltoken = string(token)
	defer LogoutCmd()
	auth1 := context.WithValue(context.Background(), openapi.ContextAccessToken, string(token))

	ctrl := gomock.NewController(t)
	mockApp := testmocks.NewMockAppsAPIService(ctrl)
	output1 := openapi.InlineResponse2002{Success: true}
	output2 := openapi.InlineResponse2002{Success: false}
	mockApp.EXPECT().DeleteAppByUser(auth1, appdata.Name).Return(output1, httpres, nil).AnyTimes()
	mockApp.EXPECT().DeleteAppByUser(auth1, "").Return(output2, httpres, nil).AnyTimes()

	newDeleteApp := DeleteAppCmd(mockApp)
	b := bytes.NewBufferString("")
	newDeleteApp.SetOut(b)

	newDeleteApp.SetArgs([]string{appdata.Name})
	newDeleteApp.Execute()
	out, err := ioutil.ReadAll(b)
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Contains(out, []byte("App deleted successfully")) {
		t.Fatalf("App cannot be deleted.")
	}

	newDeleteApp.SetArgs([]string{""})
	newDeleteApp.Execute()
	out, err = ioutil.ReadAll(b)
	if err != nil {
		t.Fatal(err)
	}
	if bytes.Contains(out, []byte("App deleted successfully")) {
		t.Fatalf("App can be deleted without app name.")
	}
}

func Test_RebuildAppCmd(t *testing.T) {
	var httpres *_nethttp.Response
	token, _ := ioutil.ReadFile(filepath.Join("testdata", "token.txt"))
	gctltoken = string(token)
	defer LogoutCmd().Execute()
	auth1 := context.WithValue(context.Background(), openapi.ContextAccessToken, string(token))

	ctrl := gomock.NewController(t)
	mockApp := testmocks.NewMockAppsAPIService(ctrl)
	output1 := openapi.InlineResponse2002{Success: true}
	output2 := openapi.InlineResponse2002{Success: false}
	mockApp.EXPECT().RebuildAppByUser(auth1, appdata.Name).Return(output1, httpres, nil).AnyTimes()
	mockApp.EXPECT().RebuildAppByUser(auth1, "").Return(output2, httpres, nil).AnyTimes()

	newRebuildApp := RebuildAppCmd(mockApp)
	b := bytes.NewBufferString("")
	newRebuildApp.SetOut(b)

	newRebuildApp.SetArgs([]string{appdata.Name})
	newRebuildApp.Execute()
	out, err := ioutil.ReadAll(b)
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Contains(out, []byte("App rebuilt successfully")) {
		t.Fatalf("App cannot be rebuilt.")
	}

	newRebuildApp.SetArgs([]string{""})
	newRebuildApp.Execute()
	out, err = ioutil.ReadAll(b)
	if err != nil {
		t.Fatal(err)
	}
	if bytes.Contains(out, []byte("App rebuilt successfully")) {
		t.Fatalf("App can be rebuilt without app name.")
	}
}

func Test_UpdateAppCmd(t *testing.T) {
	var httpres *_nethttp.Response
	token, _ := ioutil.ReadFile(filepath.Join("testdata", "token.txt"))
	gctltoken = string(token)
	defer LogoutCmd().Execute()
	auth1 := context.WithValue(context.Background(), openapi.ContextAccessToken, string(token))
	g, _ := ioutil.ReadFile(filepath.Join("testdata", "apptest.json"))
	if err := json.Unmarshal(g, &appdata); err != nil {
		t.Fatal("Error in reading app data from json file", err)
	}

	ctrl := gomock.NewController(t)
	mockApp := testmocks.NewMockAppsAPIService(ctrl)
	output1 := openapi.InlineResponse2002{Success: true}
	output2 := openapi.InlineResponse2002{Success: false}
	output3 := openapi.InlineResponse2003{Success: true, Data: []openapi.CreatedApplication{{Id: "1"}}}
	mockApp.EXPECT().UpdateAppByUser(auth1, appdata.Name, appdata).Return(output1, httpres, nil).AnyTimes()
	mockApp.EXPECT().UpdateAppByUser(auth1, "", appdata).Return(output2, httpres, nil).AnyTimes()
	mockApp.EXPECT().FetchAppByUser(auth1, appdata.Name).Return(output3, httpres, nil).AnyTimes()

	newUpdateAppCmd := UpdateAppCmd(mockApp)
	b := bytes.NewBufferString("")
	newUpdateAppCmd.SetOut(b)

	newUpdateAppCmd.SetArgs([]string{appdata.Name, "/testdata/apptest.json"})
	newUpdateAppCmd.Execute()
	out, err := ioutil.ReadAll(b)
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Contains(out, []byte("App updated successfully")) {
		t.Fatal("App cannot be updated.")
	}

	newUpdateAppCmd.SetArgs([]string{"", "/testdata/apptest.json"})
	newUpdateAppCmd.Execute()
	out, err = ioutil.ReadAll(b)
	if err != nil {
		t.Fatal(err)
	}
	if bytes.Contains(out, []byte("App updated successfully")) {
		t.Fatalf("App updated without app language.")
	}
}

func Test_FetchLogsCmd(t *testing.T) {
	var localVarOptional openapi.FetchLogsByUserOpts
	var httpres *_nethttp.Response
	token, _ := ioutil.ReadFile(filepath.Join("testdata", "token.txt"))
	gctltoken = string(token)
	defer LogoutCmd().Execute()
	auth1 := context.WithValue(context.Background(), openapi.ContextAccessToken, string(token))

	ctrl := gomock.NewController(t)
	mockApp := testmocks.NewMockAppsAPIService(ctrl)
	output1 := openapi.InlineResponse2005{Success: true}
	mockApp.EXPECT().FetchLogsByUser(auth1, appdata.Name, &localVarOptional).Return(output1, httpres, nil).AnyTimes()

	fetchLogs := FetchLogsCmd(mockApp)
	b := bytes.NewBufferString("")
	fetchLogs.SetOut(b)

	fetchLogs.SetArgs([]string{appdata.Name})
	fetchLogs.Execute()
	out, err := ioutil.ReadAll(b)
	if err != nil {
		t.Fatal(err)
	}
	if bytes.Contains(out, []byte("Error in fetching the logs.")) {
		t.Fatalf("Logs cannot be fetched")
	}
}
