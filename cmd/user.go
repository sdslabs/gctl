package cmd

import (
	"context"
	"fmt"

	"github.com/howeyc/gopass"
	openapi "github.com/sdslabs/gctl/client"
	"github.com/spf13/cobra"
)

var object openapi.InlineObject

func init() {
	fetchCmd.AddCommand(FetchUserCmd(client))
	deleteCmd.AddCommand(DeleteUserCmd(client))
	updateCmd.AddCommand(UpdateUserPasswdCmd(client))
}

//FetchUserCmd returns command to fetch details of user
func FetchUserCmd(client *openapi.APIClient) *cobra.Command {
	var fetchUserCmd = &cobra.Command{
		Use:   "user [BEARER_TOKEN]",
		Short: "View user details",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			token := args[0]
			auth := context.WithValue(context.Background(), openapi.ContextAccessToken, token)
			res, _, err := client.UserApi.FetchUser(auth)
			if res.Success {
				cmd.Print("Username: " + res.Username + "\n" + "Email: " + res.Email + "\n")
			} else {
				if err != nil {
					cmd.Print("Error:", err)
				}
			}
		},
	}
	return fetchUserCmd
}

//DeleteUserCmd returns command to delete a user
func DeleteUserCmd(client *openapi.APIClient) *cobra.Command {
	var deleteUserCmd = &cobra.Command{
		Use:   "user [BEARER_TOKEN]",
		Short: "Delete user",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			token := args[0]
			auth := context.WithValue(context.Background(), openapi.ContextAccessToken, token)
			res, _, err := client.UserApi.DeleteUser(auth)
			if res.Success {
				cmd.Print(res.Message)
			} else {
				if err != nil {
					cmd.Print("Error:", err)
				}
			}
		},
	}
	return deleteUserCmd
}

//UpdateUserPasswdCmd returns command to update password of an user
func UpdateUserPasswdCmd(client *openapi.APIClient) *cobra.Command {
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
			client.UserApi.UpdatePassword(auth, object)
			res, _, err := client.UserApi.UpdatePassword(auth, object)
			if res.Success {
				cmd.Print(res.Message)
			} else {
				if err != nil {
					cmd.Print("Error:", err)
				}
			}
		},
	}
	updateUserPasswd.Flags().StringVarP(&object.OldPassword, "oldpass", "o", "", "Old Password")
	updateUserPasswd.Flags().StringVarP(&object.NewPassword, "newpass", "n", "", "New Password")
	return updateUserPasswd
}
