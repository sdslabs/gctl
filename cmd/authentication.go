package cmd

import (
	"context"
	_context "context"
	_nethttp "net/http"
	"os"

	openapi "github.com/sdslabs/gctl/client"
	"github.com/sdslabs/gctl/cmd/middlewares"
	"github.com/spf13/cobra"
)

type AuthAPIService interface {
	Login(ctx _context.Context, email openapi.Email) (openapi.InlineResponse2002, *_nethttp.Response, error)
	Refresh(ctx _context.Context, authorization string) (openapi.LoginResponse, *_nethttp.Response, error)
}

var authAPISservice AuthAPIService = client.AuthApi

func init() {
	rootCmd.AddCommand(RefreshCmd(authAPISservice))
	rootCmd.AddCommand(LoginCmd(authAPISservice))
	rootCmd.AddCommand(LogoutCmd())
}

//LoginCmd returns a command to login in gctl
func LoginCmd(authAPIService AuthAPIService) *cobra.Command {
	var email, tempToken string
	var loginCmd = &cobra.Command{
		Use:   "login",
		Short: "Login using personal access token and email id",
		Args:  cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			var input openapi.Email
			if !middlewares.ValidateEmail(email) {
				cmd.Print("Invalid email id")
				return
			}
			if tempToken == "" {
				cmd.Print("Token not provided")
				return
			}
			auth := context.WithValue(context.Background(), openapi.ContextAccessToken, tempToken)
			input.Email = email
			res, _, err := authAPIService.Login(auth, input)
			if res.Success {
				var file *os.File
				_, err := os.Stat("gctltoken.txt")
				if os.IsNotExist(err) {
					file, err = os.Create("gctltoken.txt")
					if err != nil {
						cmd.Print("system error")
						return
					}
				} else {
					file, err = os.OpenFile("gctltoken.txt", os.O_RDWR, 0644)
					if err != nil {
						cmd.Print("system error")
						return
					}
				}
				if _, err = file.WriteString(tempToken); err != nil {
					cmd.Print("system error")
					return
				}
				err = file.Sync()
				if err != nil {
					cmd.Print("system error")
					return
				}
				cmd.Println("Logged in successfully")
			} else {
				cmd.Print(err.Error())
			}
		},
	}
	loginCmd.Flags().StringVarP(&email, "email", "e", "", "email id")
	loginCmd.Flags().StringVarP(&tempToken, "token", "t", "", "personal access token")
	return loginCmd
}

//RefreshCmd returns command to refresh existing token
func RefreshCmd(authAPIService AuthAPIService) *cobra.Command {
	var refreshCmd = &cobra.Command{
		Use:   "refresh",
		Short: "Refresh JWT token using existing token",
		Args:  cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			res, _, err := authAPIService.Refresh(context.Background(), "gctlToken "+gctltoken)
			if res.Code == 200 {
				cmd.Println("Token: ", res.Token, "\n", "Expires at: ", res.Expire)
			} else {
				cmd.Println("Error:", err)
			}
		},
	}
	return refreshCmd
}

//LogoutCmd returns a comamnd to log out from gctl
func LogoutCmd() *cobra.Command {
	var logoutCmd = &cobra.Command{
		Use:   "logout",
		Short: "Logout from gctl",
		Args:  cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			_, err := os.Stat("gctltoken.txt")
			if !os.IsNotExist(err) {
				err := os.Remove("gctltoken.txt")
				if err != nil {
					cmd.Print("system error in logout")
					return
				}
			}
		},
	}
	return logoutCmd
}
