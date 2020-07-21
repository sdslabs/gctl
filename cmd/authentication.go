package cmd

import (
	"context"
	_context "context"
	_nethttp "net/http"

	openapi "github.com/sdslabs/gctl/client"
	"github.com/spf13/cobra"
)

type AuthAPIService interface {
	Refresh(ctx _context.Context, authorization string) (openapi.LoginResponse, *_nethttp.Response, error)
}

var authAPISservice AuthAPIService = client.AuthApi

func init() {
	rootCmd.AddCommand(RefreshCmd(authAPISservice))
}

//RefreshCmd returns command to refresh existing token
func RefreshCmd(authAPIService AuthAPIService) *cobra.Command {
	var refreshCmd = &cobra.Command{
		Use:   "refresh [BEARER_TOKEN]",
		Short: "Refresh JWT token using existing token",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			token := args[0]
			res, _, err := authAPIService.Refresh(context.Background(), "gctlToken "+token)
			if res.Code == 200 {
				cmd.Println("Token: ", res.Token, "\n", "Expires at: ", res.Expire)
			} else {
				cmd.Println("Error:", err)
			}
		},
	}
	return refreshCmd
}
