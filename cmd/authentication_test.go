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

func Test_RefreshCmd(t *testing.T) {
	var httpRes *_nethttp.Response
	var loginres openapi.LoginResponse
	ctx := context.Background()
	token, _ := ioutil.ReadFile(filepath.Join("testdata", "token.txt"))
	g, _ := ioutil.ReadFile(filepath.Join("testdata", "loginresponse.json"))
	if err := json.Unmarshal(g, &loginres); err != nil {
		t.Fatal("Error in reading user details from json", err)
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockAuthAPI := testmocks.NewMockAuthAPIService(ctrl)
	loginres2 := openapi.LoginResponse{Code: 400, Token: ""}
	mockAuthAPI.EXPECT().Refresh(ctx, "gctlToken "+string(token)).Return(loginres, httpRes, nil)
	mockAuthAPI.EXPECT().Refresh(ctx, "gctlToken ").Return(loginres2, httpRes, nil)

	refreshCmd := RefreshCmd(mockAuthAPI)
	b := bytes.NewBufferString("")
	refreshCmd.SetOut(b)

	refreshCmd.SetArgs([]string{string(token)})
	refreshCmd.Execute()
	out, err := ioutil.ReadAll(b)
	if err != nil {
		t.Fatal("Error in reading output")
	}
	if bytes.Contains(out, []byte("Error")) {
		t.Fatal("Token cannot be refreshed.")
	}

	refreshCmd.SetArgs([]string{""})
	refreshCmd.Execute()
	out, err = ioutil.ReadAll(b)
	if err != nil {
		t.Fatal("Error in reading output")
	}
	if !bytes.Contains(out, []byte("Error")) {
		t.Fatal("Token can be refreshed without old token.")
	}
}
