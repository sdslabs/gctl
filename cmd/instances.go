package cmd

import (
	"context"
	_context "context"
	_nethttp "net/http"

	openapi "github.com/sdslabs/gctl/client"
	"github.com/spf13/cobra"
)

type InstancesAPIService interface {
	FetchIntancesByUser(ctx _context.Context) (openapi.InlineResponse2001, *_nethttp.Response, error)
}

var instancesAPIService openapi.InstancesAPI = client.InstancesApi

func init() {
	fetchCmd.AddCommand(GetInstancesCmd(instancesAPIService))
}

//GetInstancesCmd returns a command to fetch all instances of a user
func GetInstancesCmd(instancesAPIService openapi.InstancesAPI) *cobra.Command {
	var getInstancesCmd = &cobra.Command{
		Use:   "instances [BEARER_TOKEN]",
		Short: "Fetch all instances owned by a user",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			token := args[0]
			auth := context.WithValue(context.Background(), openapi.ContextAccessToken, token)
			res, _, err := instancesAPIService.FetchIntancesByUser(auth)
			if res.Success && len(res.Data) != 0 {
				for i := 0; i < len(res.Data); i++ {
					cmd.Println("Id: "+res.Data[i].Id+"\t", "Name: "+res.Data[i].Name+"\t",
						"Instance Type: "+res.Data[i].InstanceType+"\t", "Language: "+res.Data[i].Language)
				}
			} else if res.Success && len(res.Data) == 0 {
				cmd.Println("No instances found")
			} else {
				cmd.Println("Error: ", err)
			}
		},
	}
	return getInstancesCmd
}
