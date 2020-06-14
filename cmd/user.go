package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/howeyc/gopass"
	openapi "github.com/sdslabs/gctl/client"
	"github.com/spf13/cobra"
)

func init() {
	fetchCmd.AddCommand(fetchUserCmd)
	deleteCmd.AddCommand(deleteUserCmd)
	updateCmd.AddCommand(updateUserPasswd)
}

var fetchUserCmd = &cobra.Command{
	Use:   "user [BEARER_TOKEN]",
	Short: "View user details",
	Run: func(cmd *cobra.Command, args []string) {
		token := args[0]
		auth := context.WithValue(context.Background(), openapi.ContextAccessToken, token)
		res, _, err := client.UserApi.FetchUser(auth)
		if err != nil {
			cmd.Print(err)
			os.Exit(1)
		}
		cmd.Print("Username: " + res.Username + "\n" + "Email: " + res.Email + "\n")
	},
}

var deleteUserCmd = &cobra.Command{
	Use:   "user [BEARER_TOKEN]",
	Short: "Delete user",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		token := args[0]
		auth := context.WithValue(context.Background(), openapi.ContextAccessToken, token)
		fmt.Println(client.UserApi.DeleteUser(auth))
	},
}

var updateUserPasswd = &cobra.Command{
	Use:   "user [BEARER_TOKEN]",
	Short: "Update the password of the logged in user",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		token := args[0]
		var object openapi.InlineObject
		fmt.Printf("Old Password: ")
		maskedOldPasswd, _ := gopass.GetPasswdMasked()
		fmt.Printf("New Password: ")
		maskedNewPasswd, _ := gopass.GetPasswdMasked()
		object.OldPassword = string(maskedOldPasswd)
		object.NewPassword = string(maskedNewPasswd)
		auth := context.WithValue(context.Background(), openapi.ContextAccessToken, token)
		fmt.Println(client.UserApi.UpdatePassword(auth, object))
	},
}
