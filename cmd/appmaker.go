package cmd

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/howeyc/gopass"
	openapi "github.com/sdslabs/gctl/client"
	"github.com/spf13/cobra"
)

var appmakerCmd = &cobra.Command{
	Use:   "app",
	Short: "Create an application",
	Run: func(cmd *cobra.Command, args []string) {
		var application openapi.Application
		scanner := bufio.NewScanner(os.Stdin)
		cmd.Printf("*Token:")
		scanner.Scan()
		token := scanner.Text()
		cmd.Printf("*App Name: ")
		scanner.Scan()
		application.Name = scanner.Text()
		cmd.Printf("*Language: ")
		scanner.Scan()
		language := scanner.Text()
		cmd.Printf("*Application Password: ")
		maskedPasswd, _ := gopass.GetPasswdMasked()
		application.Password = string(maskedPasswd)
		cmd.Printf("*Git URL: ")
		scanner.Scan()
		application.Git.RepoUrl = scanner.Text()
		cmd.Printf("Is this repo private? [yes/no]: ")
		scanner.Scan()
		if scanner.Text() == "yes" {
			cmd.Printf("*Git Access Token: ")
			scanner.Scan()
			application.Git.AccessToken = scanner.Text()
		}
		cmd.Printf("Branch: ")
		scanner.Scan()
		application.Git.Branch = scanner.Text()
		cmd.Printf("*Index: ")
		scanner.Scan()
		application.Context.Index = scanner.Text()
		cmd.Printf("Port: ")
		fmt.Scanln(&application.Context.Port)
		cmd.Printf("Does this repo contain Gasperfile.txt? [yes/no]: ")
		scanner.Scan()
		if scanner.Text() == "no" {
			cmd.Printf("Build Commands: ")
			scanner.Scan()
			application.Context.Build = strings.Split(scanner.Text(), ",")
			cmd.Printf("Run Commands: ")
			scanner.Scan()
			application.Context.Run = strings.Split(scanner.Text(), ",")
		} else {
			application.Context.RcFile = true
		}
		cmd.Printf("Environment Variables(key:value): ")
		scanner.Scan()
		if scanner.Text() != "" {
			m := make(map[string]string)
			vars := strings.Split(scanner.Text(), ",")
			for v := 0; v < len(vars); v++ {
				key := strings.Split(vars[v], ":")[0]
				value := strings.Split(vars[v], ":")[1]
				m[key] = value
			}
			application.Env = m
		}
		auth := context.WithValue(context.Background(), openapi.ContextAccessToken, token)
		client.AppsApi.CreateApp(auth, language, application)
		fmt.Println(client.AppsApi.FetchAppByUser(auth, application.Name))
	},
}
