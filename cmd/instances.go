package cmd

import (
	"context"
	"fmt"

	openapi "github.com/sdslabs/gctl/client"
	"github.com/spf13/cobra"
)

func init() {
	fetchCmd.AddCommand(getInstancesCmd)
}

var getInstancesCmd = &cobra.Command{
	Use:   "instances [BEARER_TOKEN]",
	Short: "Fetch all instances owned by a user",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		token := args[0]
		auth := context.WithValue(context.Background(), openapi.ContextAccessToken, token)
		fmt.Println(client.InstancesApi.FetchIntancesByUser(auth))
	},
}
