package cmd

import (
	"context"
	"fmt"

	"github.com/howeyc/gopass"
	openapi "github.com/sdslabs/gctl/client"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(loginCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of Gasper",
	//TODO RUN
}

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Login to get a bearer token ",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Email: ")
		var email string
		fmt.Scanln(&email)
		fmt.Printf("Password: ")
		maskedPasswd, err := gopass.GetPasswdMasked()
		fmt.Println(string(maskedPasswd))
		if err != nil {
			fmt.Println(err)
		}
		l := openapi.Login{Email: email, Password: string(maskedPasswd)}
		cfg := openapi.NewConfiguration()
		client := openapi.NewAPIClient(cfg)
		fmt.Println(client.AuthApi.Login(context.Background(), l))
	},
}
