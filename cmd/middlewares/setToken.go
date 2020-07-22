package middlewares

import (
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"

	openapi "github.com/sdslabs/gctl/client"
)

//SetToken sets gctltoken
func SetToken(client *openapi.APIClient) (string, error) {
	var tokenres openapi.LoginResponse
	var gctltoken string
	_, err := os.Stat(filepath.Join("/tmp", "gctltoken.json"))
	if os.IsNotExist(err) {
		return "", errors.New("you are not logged in. Log in using command 'gctl login'")
	}
	g, err := ioutil.ReadFile(filepath.Join("/tmp", "gctltoken.json"))
	if err != nil {
		return "", err
	}
	if err := json.Unmarshal(g, &tokenres); err != nil {
		return "", err
	}
	gctltoken = tokenres.Token
	if tokenres.Expire.Sub(time.Now()) < 0 {
		res, _, err := client.AuthApi.Refresh(context.Background(), "gctlToken "+tokenres.Token)
		if res.Code == 200 {
			jsonBytes, _ := json.Marshal(res)
			file, err := os.OpenFile(filepath.Join("/tmp", "gctltoken.json"), os.O_RDWR, 0644)
			if err != nil {
				return "", err
			}
			if _, err = file.Write(jsonBytes); err != nil {
				return "", err
			}
			err = file.Sync()
			if err != nil {
				return "", err
			}
			gctltoken = res.Token
		} else {
			return "", err
		}
	}
	return gctltoken, nil
}
