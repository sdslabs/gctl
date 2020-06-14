package cmd

import (
	"bufio"
	"context"
	"fmt"
	"os"

	"github.com/antihax/optional"
	openapi "github.com/sdslabs/gctl/client"
	"github.com/spf13/cobra"
)

var dbName string

func init() {
	fetchDbCmd.Flags().StringVarP(&dbName, "name", "n", "", "Fetch specific database")
}

var dbmakerCmd = &cobra.Command{
	Use:   "db",
	Short: "Create a database",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		var database openapi.Database
		scanner := bufio.NewScanner(os.Stdin)
		cmd.Printf("Token: ")
		scanner.Scan()
		token := scanner.Text()
		cmd.Printf("Database Name: ")
		scanner.Scan()
		database.Name = scanner.Text()
		cmd.Printf("Database Password: ")
		scanner.Scan()
		database.Password = scanner.Text()
		cmd.Printf("Database Type: ")
		scanner.Scan()
		dbtype := scanner.Text()
		localVarOptional := &openapi.CreateDBOpts{
			Database: optional.NewInterface(database),
		}
		auth := context.WithValue(context.Background(), openapi.ContextAccessToken, token)
		fmt.Print(localVarOptional.Database.Value())
		fmt.Print(client.DbsApi.CreateDB(auth, dbtype, localVarOptional))
	},
}

var fetchDbCmd = &cobra.Command{
	Use:   "db [BEARER_TOKEN]",
	Short: "Fetch database owned by a user",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		dbName, _ := cmd.Flags().GetString("name")
		token := args[0]
		auth := context.WithValue(context.Background(), openapi.ContextAccessToken, token)
		if dbName != "" {
			fmt.Println(client.DbsApi.FetchDbByUser(auth, dbName))
		} else {
			fmt.Println(client.DbsApi.FetchDbsByUser(auth))
		}
	},
}

var deleteDbCmd = &cobra.Command{
	Use:   "db [DB_NAME] [BEARER_TOKEN]",
	Short: "Delete a single database owned by a user",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		dbName := args[0]
		token := args[1]
		auth := context.WithValue(context.Background(), openapi.ContextAccessToken, token)
		fmt.Println(client.DbsApi.DeleteDbByUser(auth, dbName))
	},
}
