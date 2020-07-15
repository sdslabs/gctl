package cmd

import (
	"context"
	_context "context"
	"fmt"
	_nethttp "net/http"

	"github.com/howeyc/gopass"
	openapi "github.com/sdslabs/gctl/client"
	"github.com/spf13/cobra"
)

type UserAPIService interface {
	DeleteUser(ctx _context.Context) (openapi.InlineResponse2009, *_nethttp.Response, error)
	FetchUser(ctx _context.Context) (openapi.InlineResponse2008, *_nethttp.Response, error)
	UpdatePassword(ctx _context.Context, inlineObject openapi.InlineObject) (openapi.InlineResponse20010, *_nethttp.Response, error)
}

var object openapi.InlineObject
var userAPIService openapi.UserAPI = client.UserApi

func init() {
	fetchCmd.AddCommand(FetchUserCmd(userAPIService))
	deleteCmd.AddCommand(DeleteUserCmd(userAPIService))
	updateCmd.AddCommand(UpdateUserPasswdCmd(userAPIService))
}

//FetchUserCmd returns command to fetch details of user
func FetchUserCmd(userAPIService openapi.UserAPI) *cobra.Command {
	var fetchUserCmd = &cobra.Command{
		Use:   "user [BEARER_TOKEN]",
		Short: "View user details",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			token := args[0]
			auth := context.WithValue(context.Background(), openapi.ContextAccessToken, token)
			res, _, err := userAPIService.FetchUser(auth)
			if res.Success {
				cmd.Print("Username: " + res.Username + "\n" + "Email: " + res.Email + "\n")
			} else {
				cmd.Print("Error:", err)
			}
		},
	}
	return fetchUserCmd
}

//DeleteUserCmd returns command to delete a user
func DeleteUserCmd(userAPIService openapi.UserAPI) *cobra.Command {
	var deleteUserCmd = &cobra.Command{
		Use:   "user [BEARER_TOKEN]",
		Short: "Delete user",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			token := args[0]
			auth := context.WithValue(context.Background(), openapi.ContextAccessToken, token)
			res, _, err := userAPIService.DeleteUser(auth)
			if res.Success {
				cmd.Print(res.Message)
			} else {
				cmd.Print("Error:", err)
			}
		},
	}
	return deleteUserCmd
}

//UpdateUserPasswdCmd returns command to update password of an user
func UpdateUserPasswdCmd(userAPIService openapi.UserAPI) *cobra.Command {
	var updateUserPasswd = &cobra.Command{
		Use:   "user [BEARER_TOKEN]",
		Short: "Update the password of the logged in user",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			token := args[0]
			object.OldPassword, _ = cmd.Flags().GetString("oldpass")
			object.NewPassword, _ = cmd.Flags().GetString("newpass")
			if object.NewPassword == "" && object.OldPassword == "" {
				fmt.Printf("Old Password: ")
				maskedOldPasswd, _ := gopass.GetPasswdMasked()
				fmt.Printf("New Password: ")
				maskedNewPasswd, _ := gopass.GetPasswdMasked()
				object.OldPassword = string(maskedOldPasswd)
				object.NewPassword = string(maskedNewPasswd)
			}
			auth := context.WithValue(context.Background(), openapi.ContextAccessToken, token)
			res, _, err := userAPIService.UpdatePassword(auth, object)
			if res.Success {
				cmd.Print(res.Message)
			} else {
				cmd.Print("Error:", err)
			}
		},
	}
	updateUserPasswd.Flags().StringVarP(&object.OldPassword, "oldpass", "o", "", "Old Password")
	updateUserPasswd.Flags().StringVarP(&object.NewPassword, "newpass", "n", "", "New Password")
	return updateUserPasswd
}
