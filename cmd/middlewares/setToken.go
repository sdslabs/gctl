package middlewares

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"time"

	openapi "github.com/sdslabs/gctl/client"
)

func SetToken(client *openapi.APIClient) string {
	var tokenres openapi.LoginResponse
	var gctltoken string
	g, err := ioutil.ReadFile(filepath.Join("/tmp", "gctltoken.json"))
	if err != nil {
		fmt.Print(err)
	}
	if err := json.Unmarshal(g, &tokenres); err != nil {
		fmt.Print(err)
	}
	if tokenres.Expire.Sub(time.Now()) < 0 {
		fmt.Print(tokenres.Expire, time.Now())
		res, _, err := client.AuthApi.Refresh(context.Background(), "gctlToken "+gctltoken)
		if res.Code == 200 {
			gctltoken = res.Token
		} else {
			fmt.Print(err)
		}
	} else {
		gctltoken = tokenres.Token
	}
	return gctltoken
}
