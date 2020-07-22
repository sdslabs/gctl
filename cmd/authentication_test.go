package cmd

import (
	"bytes"
	"context"
	"io/ioutil"
	_nethttp "net/http"
	"path/filepath"
	"testing"

	"github.com/golang/mock/gomock"
	openapi "github.com/sdslabs/gctl/client"
	"github.com/sdslabs/gctl/cmd/testmocks"
)

func Test_LoginCmd(t *testing.T) {
	var httpRes *_nethttp.Response
	token, _ := ioutil.ReadFile(filepath.Join("testdata", "token.txt"))
	email, _ := ioutil.ReadFile(filepath.Join("testdata", "email.txt"))
	auth1 := context.WithValue(context.Background(), openapi.ContextAccessToken, string(token))
	input1 := openapi.Email{Email: string(email)}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockAuthAPI := testmocks.NewMockAuthAPIService(ctrl)
	output1 := openapi.InlineResponse2004{Success: true}
	mockAuthAPI.EXPECT().Login(auth1, input1).Return(output1, httpRes, nil)

	logincmd := LoginCmd(mockAuthAPI)
	b := bytes.NewBufferString("")
	logincmd.SetOut(b)

	logincmd.SetArgs([]string{"-e", string(email), "-t", string(token)})
	logincmd.Execute()
	out, err := ioutil.ReadAll(b)
	if err != nil {
		t.Fatal("Error in reading output")
	}
	if !bytes.Contains(out, []byte("Logged in successfully")) {
		t.Fatal("Cannot log in", string(out))
	}

	logincmd.SetArgs([]string{"-e", "", "-t", string(token)})
	logincmd.Execute()
	out, err = ioutil.ReadAll(b)
	if err != nil {
		t.Fatal("Error in reading output")
	}
	if !bytes.Contains(out, []byte("Invalid email id")) {
		t.Fatal("Can log in without email")
	}

	logincmd.SetArgs([]string{"-e", string(email), "-t", ""})
	logincmd.Execute()
	out, err = ioutil.ReadAll(b)
	if err != nil {
		t.Fatal("Error in reading output")
	}
	if !bytes.Contains(out, []byte("Token not provided")) {
		t.Fatal("Can log in without personal access token")
	}
}

func Test_LogoutCmd(t *testing.T) {
	logoutCmd := LogoutCmd()
	b := bytes.NewBufferString("")
	logoutCmd.SetOut(b)

	logoutCmd.Execute()
	out, err := ioutil.ReadAll(b)
	if err != nil {
		t.Fatal("Error in reading output")
	}
	if bytes.Compare(out, []byte("")) != 0 {
		t.Fatal("Token cannot be refreshed.")
	}
}
