package cmd

import (
	"context"
	_context "context"
	_nethttp "net/http"

	openapi "github.com/sdslabs/gctl/client"
	"github.com/sdslabs/gctl/cmd/middlewares"
	"github.com/spf13/cobra"
)

type AuthAPIService interface {
	Login(ctx _context.Context, login openapi.Login) (openapi.LoginResponse, *_nethttp.Response, error)
	Refresh(ctx _context.Context, authorization string) (openapi.LoginResponse, *_nethttp.Response, error)
	Register(ctx _context.Context, user openapi.User) (openapi.InlineResponse200, *_nethttp.Response, error)
}

var authAPISservice openapi.AuthAPI = client.AuthApi
var loginCreds openapi.Login
var registerCreds openapi.User

func init() {
	rootCmd.AddCommand(LoginCmd(authAPISservice))
	rootCmd.AddCommand(RegisterCmd(authAPISservice))
	rootCmd.AddCommand(RefreshCmd(authAPISservice))
}

//LoginCmd returns command for login
func LoginCmd(authAPIService openapi.AuthAPI) *cobra.Command {
	var loginCmd = &cobra.Command{
		Use:   "login",
		Short: "Login to get a bearer token ",
		Run: func(cmd *cobra.Command, args []string) {
			loginCreds.Email, _ = cmd.Flags().GetString("email")
			loginCreds.Password, _ = cmd.Flags().GetString("password")
			if loginCreds.Email == "" && loginCreds.Password == "" {
				loginCreds = middlewares.LoginForm()
			}
			res, _, err := authAPIService.Login(context.Background(), loginCreds)
			if res.Token != "" {
				cmd.Println("Token: ", res.Token, "\n", "Expires at: ", res.Expire)
			} else {
				cmd.Println("Error:", err)
			}
		},
	}
	loginCmd.Flags().StringVarP(&loginCreds.Email, "email", "e", "", "Email of the user")
	loginCmd.Flags().StringVarP(&loginCreds.Password, "password", "p", "", "Password")
	return loginCmd
}

//RegisterCmd returns command to register a user to gasper
func RegisterCmd(authAPIService openapi.AuthAPI) *cobra.Command {
	var registerCmd = &cobra.Command{
		Use:   "register",
		Short: "Register a user",
		Run: func(cmd *cobra.Command, args []string) {
			registerCreds.Email, _ = cmd.Flags().GetString("email")
			registerCreds.Password, _ = cmd.Flags().GetString("password")
			registerCreds.Username, _ = cmd.Flags().GetString("username")
			if registerCreds.Email == "" && registerCreds.Password == "" && registerCreds.Username == "" {
				registerCreds = middlewares.RegisterForm()
			}
			res, _, err := authAPIService.Register(context.Background(), registerCreds)
			if res.Success {
				cmd.Println(res.Message)
			} else {
				cmd.Println(err)
			}
		},
	}
	registerCmd.Flags().StringVarP(&registerCreds.Email, "email", "e", "", "Email of the user")
	registerCmd.Flags().StringVarP(&registerCreds.Username, "username", "u", "", "Username")
	registerCmd.Flags().StringVarP(&registerCreds.Password, "password", "p", "", "Password")
	return registerCmd
}

//RefreshCmd returns command to refresh existing token
func RefreshCmd(authAPIService openapi.AuthAPI) *cobra.Command {
	var refreshCmd = &cobra.Command{
		Use:   "refresh [BEARER_TOKEN]",
		Short: "Refresh JWT token using existing token",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			token := args[0]
			res, _, err := client.AuthApi.Refresh(context.Background(), "Bearer "+token)
			if res.Code == 200 {
				cmd.Println("Token: ", res.Token, "\n", "Expires at: ", res.Expire)
			} else {
				if err != nil {
					cmd.Println("Error:", err)
				}
			}
		},
	}
	return refreshCmd
}
