package cmd

import (
	"context"
	_context "context"
	"encoding/json"
	_nethttp "net/http"
	"os"
	"path/filepath"

	openapi "github.com/sdslabs/gctl/client"
	"github.com/sdslabs/gctl/cmd/middlewares"
	"github.com/spf13/cobra"
)

type AuthAPIService interface {
	Login(ctx _context.Context, email openapi.Email) (openapi.InlineResponse2004, *_nethttp.Response, error)
}

var authAPISservice AuthAPIService = client.AuthApi

func init() {
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
			var (
				input openapi.Email
				file  *os.File
			)

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
				data := openapi.LoginResponse{Token: tempToken, Expire: res.Expire}
				jsonBytes, _ := json.Marshal(data)
				_, err := os.Stat(filepath.Join("/tmp", "gctltoken.json"))

				if os.IsNotExist(err) {
					file, err = os.Create(filepath.Join("/tmp", "gctltoken.json"))
					if err != nil {
						cmd.Print(err)
						return
					}
				} else {
					file, err = os.OpenFile(filepath.Join("/tmp", "gctltoken.json"), os.O_RDWR, 0644)
					if err != nil {
						cmd.Print(err)
						return
					}
				}

				if _, err = file.Write(jsonBytes); err != nil {
					cmd.Print(err)
					return
				}

				if err = file.Sync(); err != nil {
					cmd.Print(err)
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

//LogoutCmd returns a comamnd to log out from gctl
func LogoutCmd() *cobra.Command {
	var logoutCmd = &cobra.Command{
		Use:   "logout",
		Short: "Logout from gctl",
		Args:  cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			_, err := os.Stat(filepath.Join("/tmp", "gctltoken.json"))

			if !os.IsNotExist(err) {
				err := os.Remove(filepath.Join("/tmp", "gctltoken.json"))
				if err != nil {
					cmd.Print(err)
					return
				}
			}
		},
	}
	return logoutCmd
}
