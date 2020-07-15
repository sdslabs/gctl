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

func Test_FetchUserCmd(t *testing.T) {
	var httpres *_nethttp.Response
	token, _ := ioutil.ReadFile(filepath.Join("testdata", "token.txt"))
	auth1 := context.WithValue(context.Background(), openapi.ContextAccessToken, string(token))
	auth2 := context.WithValue(context.Background(), openapi.ContextAccessToken, string(""))

	ctrl := gomock.NewController(t)
	mockUser := testmocks.NewMockUserAPIService(ctrl)
	fetchres1 := openapi.InlineResponse2008{Success: true}
	fetchres2 := openapi.InlineResponse2008{Success: false}
	mockUser.EXPECT().FetchUser(auth1).Return(fetchres1, httpres, nil).AnyTimes()
	mockUser.EXPECT().FetchUser(auth2).Return(fetchres2, httpres, nil)

	newFetchUser := FetchUserCmd(mockUser)
	b := bytes.NewBufferString("")
	newFetchUser.SetOut(b)

	newFetchUser.SetArgs([]string{string(token)})
	newFetchUser.Execute()
	out, err := ioutil.ReadAll(b)
	if err != nil {
		t.Fatal("Error in reading output")
	}
	if bytes.Contains(out, []byte("Error")) {
		t.Fatalf("User details cannot be fetched")
	}

	newFetchUser.SetArgs([]string{string("")})
	newFetchUser.Execute()
	out, err = ioutil.ReadAll(b)
	if err != nil {
		t.Fatal("Error in reading output")
	}
	if !bytes.Contains(out, []byte("Error")) {
		t.Fatalf("User details can be fetched without token")
	}
}

func Test_UpdateUserCmd(t *testing.T) {
	var httpres *_nethttp.Response
	var updatePassDetails openapi.InlineObject
	token, _ := ioutil.ReadFile(filepath.Join("testdata", "token.txt"))
	auth1 := context.WithValue(context.Background(), openapi.ContextAccessToken, string(token))
	auth2 := context.WithValue(context.Background(), openapi.ContextAccessToken, string(""))
	g, _ := ioutil.ReadFile(filepath.Join("testdata", "updatepassdata.json"))
	if err := json.Unmarshal(g, &updatePassDetails); err != nil {
		t.Fatal("Error in reading update password details from json", err)
	}
	updatePassDetails2 := openapi.InlineObject{OldPassword: updatePassDetails.OldPassword, NewPassword: ""}

	ctrl := gomock.NewController(t)
	mockUser := testmocks.NewMockUserAPIService(ctrl)
	updateres1 := openapi.InlineResponse20010{Success: true}
	updateres2 := openapi.InlineResponse20010{Success: false}
	mockUser.EXPECT().UpdatePassword(auth1, updatePassDetails).Return(updateres1, httpres, nil).AnyTimes()
	mockUser.EXPECT().UpdatePassword(auth2, updatePassDetails2).Return(updateres2, httpres, nil).AnyTimes()

	newUpdateCmd := UpdateUserPasswdCmd(mockUser)
	b := bytes.NewBufferString("")
	newUpdateCmd.SetOut(b)

	newUpdateCmd.SetArgs([]string{string(token), "-o", updatePassDetails.OldPassword, "-n", updatePassDetails.NewPassword})
	newUpdateCmd.Execute()
	out, err := ioutil.ReadAll(b)
	if err != nil {
		t.Fatal("Error in reading output")
	}
	if bytes.Contains(out, []byte("Error")) {
		t.Fatalf("User password cannot be updated")
	}

	newUpdateCmd.SetArgs([]string{"", "-o", updatePassDetails.OldPassword, "-n", ""})
	newUpdateCmd.Execute()
	out, err = ioutil.ReadAll(b)
	if err != nil {
		t.Fatal("Error in reading output")
	}
	if !bytes.Contains(out, []byte("Error")) {
		t.Fatalf("User password can be updated without new password")
	}
}

func Test_DeleteUserCmd(t *testing.T) {
	var httpres *_nethttp.Response
	token, _ := ioutil.ReadFile(filepath.Join("testdata", "token.txt"))
	auth1 := context.WithValue(context.Background(), openapi.ContextAccessToken, string(token))
	auth2 := context.WithValue(context.Background(), openapi.ContextAccessToken, string(""))

	ctrl := gomock.NewController(t)
	mockUser := testmocks.NewMockUserAPIService(ctrl)
	deleteres1 := openapi.InlineResponse2009{Success: true}
	deleteres2 := openapi.InlineResponse2009{Success: false}
	mockUser.EXPECT().DeleteUser(auth1).Return(deleteres1, httpres, nil).AnyTimes()
	mockUser.EXPECT().DeleteUser(auth2).Return(deleteres2, httpres, nil).AnyTimes()

	newDeleteUserCmd := DeleteUserCmd(mockUser)
	b := bytes.NewBufferString("")
	newDeleteUserCmd.SetOut(b)

	newDeleteUserCmd.SetArgs([]string{string(token)})
	newDeleteUserCmd.Execute()
	out, err := ioutil.ReadAll(b)
	if err != nil {
		t.Fatal("Error in reading output")
	}
	if bytes.Contains(out, []byte("Error")) {
		t.Fatalf("User cannot be deleted")
	}

	newDeleteUserCmd.SetArgs([]string{""})
	newDeleteUserCmd.Execute()
	out, err = ioutil.ReadAll(b)
	if err != nil {
		t.Fatal("Error in reading output")
	}
	if !bytes.Contains(out, []byte("Error")) {
		t.Fatalf("User cannot be deleted")
	}
}
