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

var dbmakerCmd = &cobra.Command{
	Use:   "db",
	Short: "Create a database",
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
