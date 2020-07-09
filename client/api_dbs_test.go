package openapi

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"path/filepath"
	"testing"

	"github.com/antihax/optional"
)

func TestApiDb(t *testing.T) {
	var dbdata Database
	g, _ := ioutil.ReadFile(filepath.Join("testdata", "dbdata.json"))
	json.Unmarshal(g, &dbdata)
	auth := context.WithValue(context.Background(), ContextAccessToken, token)
	localVarOptional := CreateDBOpts{
		Database: optional.NewInterface(dbdata),
	}
	dbres, _, err := client.DbsApi.CreateDB(auth, "mysql", &localVarOptional)
	if !dbres.Success {
		t.Fatal("Error in creating database", err)
	}
	fetchDbres, _, err := client.DbsApi.FetchDbByUser(auth, dbdata.Name)
	if !fetchDbres.Success {
		t.Fatal("Error in fetching database", err)
	}
	fetchDbsres, _, err := client.DbsApi.FetchDbsByUser(auth)
	if !fetchDbsres.Success {
		t.Fatal("Error in fetching databases", err)
	}
	deleteDbres, _, err := client.DbsApi.DeleteDbByUser(auth, dbdata.Name)
	if !deleteDbres.Success {
		t.Fatal("Error in deleting database", err)
	}
}
