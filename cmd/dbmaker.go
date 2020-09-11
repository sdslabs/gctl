package cmd

import (
	"context"
	_context "context"
	_nethttp "net/http"

	"github.com/antihax/optional"
	openapi "github.com/sdslabs/gctl/client"
	"github.com/sdslabs/gctl/cmd/middlewares"
	"github.com/spf13/cobra"
)

//DbsAPIService is interface for all client functions of dbs
type DbsAPIService interface {
	CreateDb(ctx _context.Context, databaseType string, localVarOptionals *openapi.CreateDBOpts) (openapi.InlineResponse2002, *_nethttp.Response, error)
	DeleteDbByUser(ctx _context.Context, db string) (openapi.InlineResponse2002, *_nethttp.Response, error)
	FetchDbByUser(ctx _context.Context, db string) (openapi.InlineResponse2007, *_nethttp.Response, error)
	FetchDbsByUser(ctx _context.Context) (openapi.InlineResponse2007, *_nethttp.Response, error)
}

var dbName string
var dbtype string
var db openapi.Database
var dbsAPIService DbsAPIService = client.DbsAPI

func init() {
	createCmd.AddCommand(CreateDbCmd(dbsAPIService))
	fetchCmd.AddCommand(FetchDbCmd(dbsAPIService))
	deleteCmd.AddCommand(DeleteDbCmd(dbsAPIService))
}

//CreateDbCmd returns command to create a database
func CreateDbCmd(dbsAPIService DbsAPIService) *cobra.Command {
	var dbmakerCmd = &cobra.Command{
		Use:   "db",
		Short: "Create a database",
		Args:  cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			var err error

			if gctltoken == "" {
				gctltoken, err = middlewares.SetToken(client)
				if err != nil {
					cmd.Print(err)
					return
				}
			}

			dbtype, _ = cmd.Flags().GetString("dbtype")
			db.Name, _ = cmd.Flags().GetString("name")
			db.Password, _ = cmd.Flags().GetString("password")

			if dbtype == "" && db.Name == "" && db.Password == "" {
				dbtype, db = middlewares.DbForm()
			}

			localVarOptional := &openapi.CreateDBOpts{
				Database: optional.NewInterface(db),
			}

			auth := context.WithValue(context.Background(), openapi.ContextAccessToken, gctltoken)
			res, _, err := dbsAPIService.CreateDb(auth, dbtype, localVarOptional)

			if res.Success {
				cmd.Print("Database created")
			} else {
				cmd.Print(err)
			}
		},
	}

	dbmakerCmd.Flags().StringVarP(&db.Name, "name", "n", "", "Database name")
	dbmakerCmd.Flags().StringVarP(&db.Password, "password", "p", "", "Database password")
	dbmakerCmd.Flags().StringVarP(&dbtype, "dbtype", "t", "", "Database type")
	return dbmakerCmd
}

//FetchDbCmd returns command to fetch databases of a user
func FetchDbCmd(dbsAPIService DbsAPIService) *cobra.Command {
	var fetchDbCmd = &cobra.Command{
		Use:   "db",
		Short: "Fetch database owned by a user",
		Args:  cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			var err error

			if gctltoken == "" {
				gctltoken, err = middlewares.SetToken(client)
				if err != nil {
					cmd.Print(err)
					return
				}
			}

			dbName, _ := cmd.Flags().GetString("name")
			auth := context.WithValue(context.Background(), openapi.ContextAccessToken, gctltoken)

			if dbName != "" {
				res, _, err := dbsAPIService.FetchDbByUser(auth, dbName)
				if res.Success {
					if len(res.Data) != 0 {
						for i := 0; i < len(res.Data); i++ {
							cmd.Println("User: "+res.Data[i].User, "Owner: "+res.Data[i].Owner,
								"DbUrl: "+res.Data[i].DbUrl, "Port: ", res.Data[i].Port,
								"Host: "+res.Data[i].HostIp, "Language: "+res.Data[i].Language,
								"Instance Type: "+res.Data[i].InstanceType)
						}
					} else {
						cmd.Println("No such database")
					}
				} else {
					cmd.Println("Error in fetching the database", err)
				}
			} else {
				res, _, err := dbsAPIService.FetchDbsByUser(auth)
				if res.Success {
					if len(res.Data) != 0 {
						for i := 0; i < len(res.Data); i++ {
							cmd.Println("User: "+res.Data[i].User, "Owner: "+res.Data[i].Owner, "DbUrl: "+res.Data[i].DbUrl, "Port: ", res.Data[i].Port,
								"Host: "+res.Data[i].HostIp, "Language: "+res.Data[i].Language, "Instance Type: "+res.Data[i].InstanceType)
						}
					} else {
						cmd.Println("No database for the user")
					}
				} else {
					cmd.Println("Error in fetching the databases", err)
				}
			}
		},
	}

	fetchDbCmd.Flags().StringVarP(&dbName, "name", "n", "", "Fetch specific database")
	return fetchDbCmd
}

//DeleteDbCmd returns command to delete database owned by a user
func DeleteDbCmd(dbsAPIService DbsAPIService) *cobra.Command {
	var deleteDbCmd = &cobra.Command{
		Use:   "db [DB_NAME]",
		Short: "Delete a single database owned by a user",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			var err error

			if gctltoken == "" {
				gctltoken, err = middlewares.SetToken(client)
				if err != nil {
					cmd.Print(err)
					return
				}
			}

			dbName := args[0]
			auth := context.WithValue(context.Background(), openapi.ContextAccessToken, gctltoken)
			res, _, err := dbsAPIService.DeleteDbByUser(auth, dbName)
			if res.Success {
				cmd.Println("Database deleted successfully")
			} else {
				cmd.Println(err)
			}
		},
	}
	return deleteDbCmd
}
