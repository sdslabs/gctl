package middlewares

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
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
	gctltoken = tokenres.Token
	if tokenres.Expire.Sub(time.Now()) < 0 {
		res, _, err := client.AuthApi.Refresh(context.Background(), "gctlToken "+tokenres.Token)
		if res.Code == 200 {
			jsonBytes, _ := json.Marshal(res)
			file, err := os.OpenFile(filepath.Join("/tmp", "gctltoken.json"), os.O_RDWR, 0644)
			if err != nil {
				fmt.Print("system error2")
			}
			if _, err = file.Write(jsonBytes); err != nil {
				fmt.Print("system error3")
			}
			err = file.Sync()
			if err != nil {
				fmt.Print("system error4")
			}
			gctltoken = res.Token
		} else {
			fmt.Print(err)
		}
	}
	return gctltoken
}
