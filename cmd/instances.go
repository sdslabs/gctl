package cmd

import (
	"context"
	_context "context"
	_nethttp "net/http"

	openapi "github.com/sdslabs/gctl/client"
	"github.com/sdslabs/gctl/cmd/middlewares"
	"github.com/spf13/cobra"
)

//InstancesAPIService is interface for all client functions of interfaces
type InstancesAPIService interface {
	FetchIntancesByUser(ctx _context.Context) (openapi.InlineResponse2001, *_nethttp.Response, error)
}

var instancesAPIService InstancesAPIService = client.InstancesAPI

func init() {
	fetchCmd.AddCommand(GetInstancesCmd(instancesAPIService))
}

//GetInstancesCmd returns a command to fetch all instances of a user
func GetInstancesCmd(instancesAPIService InstancesAPIService) *cobra.Command {
	var getInstancesCmd = &cobra.Command{
		Use:   "instances",
		Short: "Fetch all instances owned by a user",
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

			auth := context.WithValue(context.Background(), openapi.ContextAccessToken, gctltoken)
			res, _, err := instancesAPIService.FetchIntancesByUser(auth)

			if res.Success && len(res.Data) != 0 {
				for i := 0; i < len(res.Data); i++ {
					cmd.Println("Id: "+res.Data[i].Id,
						"Name: "+res.Data[i].Name,
						"Instance Type: "+res.Data[i].InstanceType,
						"Language: "+res.Data[i].Language)
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
