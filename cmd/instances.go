package cmd

import (
	"context"

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
		res, _, err := client.InstancesApi.FetchIntancesByUser(auth)
		if res.Success && len(res.Data) != 0 {
			for i := 0; i < len(res.Data); i++ {
				cmd.Println("Id: "+res.Data[i].Id+"\t", "Name: "+res.Data[i].Name+"\t",
					"Instance Type: "+res.Data[i].InstanceType+"\t", "Language: "+res.Data[i].Language)
			}
		} else {
			cmd.Println("Error: " + err.Error())
		}
	},
}
