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

func Test_GetInstancesCmd(t *testing.T) {
	var httpres *_nethttp.Response
	token, _ := ioutil.ReadFile(filepath.Join("testdata", "token.txt"))
	gctltoken = string(token)
	defer LogoutCmd().Execute()
	auth1 := context.WithValue(context.Background(), openapi.ContextAccessToken, string(token))

	ctrl := gomock.NewController(t)
	mockInstance := testmocks.NewMockInstancesAPIService(ctrl)
	instanceres1 := openapi.InlineResponse2001{Success: true, Data: []openapi.Instances{{Id: "1"}}}
	mockInstance.EXPECT().FetchIntancesByUser(auth1).Return(instanceres1, httpres, nil).AnyTimes()

	newGetInstancesCmd := GetInstancesCmd(mockInstance)
	b := bytes.NewBufferString("")
	newGetInstancesCmd.SetOut(b)

	newGetInstancesCmd.Execute()
	out, err := ioutil.ReadAll(b)
	if err != nil {
		t.Fatal("Error in reading output")
	}
	if bytes.Contains(out, []byte("Error")) {
		t.Fatalf("All instances cannot be fetched")
	}
}
