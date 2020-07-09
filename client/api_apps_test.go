package openapi

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"path/filepath"
	"testing"
)

func TestAppsApi(t *testing.T) {
	var appdata Application
	auth := context.WithValue(context.Background(), ContextAccessToken, token)
	g, _ := ioutil.ReadFile(filepath.Join("testdata", "appApiTest.json"))
	json.Unmarshal(g, &appdata)
	response, _, err := client.AppsApi.CreateApp(auth, "nodejs", appdata)
	if !response.Success {
		t.Fatal("Error in creating app", err)
	}
	fetchedApp, _, err := client.AppsApi.FetchAppByUser(auth, appdata.Name)
	if !fetchedApp.Success {
		t.Fatal("Error in fetching app", err)
	}
	fetchedApps, _, err := client.AppsApi.FetchAppsByUser(auth)
	if !fetchedApps.Success {
		t.Fatal("Error in fetching apps", err)
	}
	deleteapp, _, err := client.AppsApi.DeleteAppByUser(auth, appdata.Name)
	if !deleteapp.Success {
		t.Fatal("Error in deleting app", err)
	}
}
