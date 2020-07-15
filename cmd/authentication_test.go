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

var userdata openapi.User

func Test_RegisterCmd(t *testing.T) {
	var httpRes *_nethttp.Response
	ctx := context.Background()
	g, _ := ioutil.ReadFile(filepath.Join("testdata", "userdata.json"))
	if err := json.Unmarshal(g, &userdata); err != nil {
		t.Fatal("Error in reading user details from json", err)
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockAuthAPI := testmocks.NewMockAuthAPIService(ctrl)
	newRegisterCmd := RegisterCmd(mockAuthAPI)
	userdata2 := openapi.User{Email: "", Username: userdata.Username, Password: userdata.Password}
	registerRes := openapi.InlineResponse200{Success: true, Message: "user created"}
	registerRes2 := openapi.InlineResponse200{Success: false, Message: "failed"}
	mockAuthAPI.EXPECT().Register(ctx, userdata).Return(registerRes, httpRes, nil).AnyTimes()
	mockAuthAPI.EXPECT().Register(ctx, userdata2).Return(registerRes2, httpRes, nil)

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

	newRegisterCmd.SetArgs([]string{"-u", userdata.Username, "-e", "", "-p", userdata.Password})
	newRegisterCmd.Execute()
	out, err = ioutil.ReadAll(b)
	if err != nil {
		t.Fatal("Error in reading output")
	}
	if bytes.Contains(out, []byte("user created")) {
		t.Fatalf("User created without email.")
	}
}

func Test_LoginCmd(t *testing.T) {
	var httpRes *_nethttp.Response
	var loginres openapi.LoginResponse
	ctx := context.Background()
	g, _ := ioutil.ReadFile(filepath.Join("testdata", "loginresponse.json"))
	if err := json.Unmarshal(g, &loginres); err != nil {
		t.Fatal("Error in reading user details from json", err)
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockAuthAPI := testmocks.NewMockAuthAPIService(ctrl)
	loginCreds := openapi.Login{Email: userdata.Email, Password: userdata.Password}
	loginCreds2 := openapi.Login{Email: userdata.Email, Password: ""}
	loginres2 := openapi.LoginResponse{Code: 400, Token: ""}
	mockAuthAPI.EXPECT().Login(ctx, loginCreds).Return(loginres, httpRes, nil).AnyTimes()
	mockAuthAPI.EXPECT().Login(ctx, loginCreds2).Return(loginres2, httpRes, nil)

	newLoginCmd := LoginCmd(mockAuthAPI)
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
	}

	newLoginCmd.SetArgs([]string{"-e", userdata.Email, "-p", ""})
	newLoginCmd.Execute()
	out, err = ioutil.ReadAll(b)
	if err != nil {
		t.Fatal("Error in reading output")
	}
	if !bytes.Contains(out, []byte("Error")) {
		t.Fatalf("User can be logged in without password")
	}
}
