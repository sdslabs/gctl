package cmd

import (
	"context"

	"github.com/antihax/optional"
	openapi "github.com/sdslabs/gctl/client"
	"github.com/sdslabs/gctl/cmd/middlewares"
	"github.com/spf13/cobra"
)

var dbName string
var dbtype string
var db openapi.Database

func init() {
	createCmd.AddCommand(CreateDbCmd(client))
	fetchCmd.AddCommand(FetchDbCmd(client))
	deleteCmd.AddCommand(DeleteDbCmd(client))
}

//CreateDbCmd returns command to create a database
func CreateDbCmd(client *openapi.APIClient) *cobra.Command {
	var dbmakerCmd = &cobra.Command{
		Use:   "db [BEARER_TOKEN]",
		Short: "Create a database",
		Args:  cobra.MaximumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			token := args[0]
			dbtype, _ = cmd.Flags().GetString("dbtype")
			db.Name, _ = cmd.Flags().GetString("name")
			db.Password, _ = cmd.Flags().GetString("password")
			if token == "" && dbtype == "" && db.Name == "" && db.Password == "" {
				token, dbtype, db = middlewares.DbForm()
			}
			localVarOptional := &openapi.CreateDBOpts{
				Database: optional.NewInterface(db),
			}
			auth := context.WithValue(context.Background(), openapi.ContextAccessToken, token)
			res, _, err := client.DbsApi.CreateDB(auth, dbtype, localVarOptional)
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
func FetchDbCmd(client *openapi.APIClient) *cobra.Command {
	var fetchDbCmd = &cobra.Command{
		Use:   "db [BEARER_TOKEN]",
		Short: "Fetch database owned by a user",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			dbName, _ := cmd.Flags().GetString("name")
			token := args[0]
			auth := context.WithValue(context.Background(), openapi.ContextAccessToken, token)
			if dbName != "" {
				res, _, err := client.DbsApi.FetchDbByUser(auth, dbName)
				if res.Success {
					if len(res.Data) != 0 {
						for i := 0; i < len(res.Data); i++ {
							cmd.Println("User: "+res.Data[i].User, "Owner: "+res.Data[i].Owner, "DbUrl: "+res.Data[i].DbUrl, "Port: ", res.Data[i].Port,
								"Host: "+res.Data[i].HostIp, "Language: "+res.Data[i].Language, "Instance Type: "+res.Data[i].InstanceType)
						}
					} else {
						cmd.Println("No such database")
					}
				} else {
					cmd.Println("Error in fetching the database", err)
				}
			} else {
				res, _, err := client.DbsApi.FetchDbsByUser(auth)
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
func DeleteDbCmd(client *openapi.APIClient) *cobra.Command {
	var deleteDbCmd = &cobra.Command{
		Use:   "db [DB_NAME] [BEARER_TOKEN]",
		Short: "Delete a single database owned by a user",
		Args:  cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			dbName := args[0]
			token := args[1]
			auth := context.WithValue(context.Background(), openapi.ContextAccessToken, token)
			res, _, err := client.DbsApi.DeleteDbByUser(auth, dbName)
			if res.Success {
				cmd.Println("Database deleted successfully")
			} else {
				cmd.Println(err)
			}
		},
	}
	return deleteDbCmd
}
