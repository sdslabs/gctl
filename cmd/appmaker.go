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

var appName string

func init() {
	createCmd.AddCommand(appmakerCmd)
	fetchCmd.AddCommand(fetchAppCmd)
	deleteCmd.AddCommand(deleteAppCmd)
	fetchAppCmd.Flags().StringVarP(&appName, "name", "n", "", "show specific app")
}

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
		res, _, err := client.AppsApi.CreateApp(auth, language, application)
		if res.Success {
			res, _, err := client.AppsApi.FetchAppByUser(auth, application.Name)
			if res.Success {
				for i := 0; i < len(res.Data); i++ {
					cmd.Println("App created successfully"+"\n"+"Container Id: "+res.Data[i].ContainerId, "Container Port: "+string(res.Data[i].ContainerPort),
						"Docker Image: "+res.Data[i].DockerImage, "App Url: "+res.Data[i].AppUrl, "Host Ip: "+res.Data[i].HostIp,
						"Name Servers: ", res.Data[i].NameServers, "Instance Type: "+res.Data[i].InstanceType, "Language: "+res.Data[i].Language,
						"Owner: "+res.Data[i].Owner, "Ssh Cmd: "+res.Data[i].SshCmd, "Id: "+res.Data[i].Id)
				}
			} else {
				cmd.Println(err)
			}
		} else {
			cmd.Println(err)
		}
	},
}

var fetchAppCmd = &cobra.Command{
	Use:   "app [BEARER_TOKEN]",
	Short: "fetch apps",
	Run: func(cmd *cobra.Command, args []string) {
		appName, _ := cmd.Flags().GetString("name")
		token := args[0]
		auth := context.WithValue(context.Background(), openapi.ContextAccessToken, token)
		if appName != "" {
			res, _, err := client.AppsApi.FetchAppByUser(auth, appName)
			if res.Success {
				if len(res.Data) != 0 {
					for i := 0; i < len(res.Data); i++ {
						cmd.Println("Container Id: "+res.Data[i].ContainerId, "Container Port: "+string(res.Data[i].ContainerPort),
							"Docker Image: "+res.Data[i].DockerImage, "App Url: "+res.Data[i].AppUrl, "Host Ip: "+res.Data[i].HostIp,
							"Name Servers: ", res.Data[i].NameServers, "Instance Type: "+res.Data[i].InstanceType, "Language: "+res.Data[i].Language,
							"Owner: "+res.Data[i].Owner, "Ssh Cmd: "+res.Data[i].SshCmd, "Id: "+res.Data[i].Id)
					}
				} else {
					cmd.Println("No such app found")
				}
			} else {
				cmd.Println(err)
			}
		} else {
			res, _, err := client.AppsApi.FetchAppsByUser(auth)
			if res.Success {
				if len(res.Data) != 0 {
					for i := 0; i < len(res.Data); i++ {
						cmd.Println("Container Id: "+res.Data[i].ContainerId, "Container Port: "+string(res.Data[i].ContainerPort),
							"Docker Image: "+res.Data[i].DockerImage, "App Url: "+res.Data[i].AppUrl, "Host Ip: "+res.Data[i].HostIp,
							"Name Servers: ", res.Data[i].NameServers, "Instance Type: "+res.Data[i].InstanceType, "Language: "+res.Data[i].Language,
							"Owner: "+res.Data[i].Owner, "Ssh Cmd: "+res.Data[i].SshCmd, "Id: "+res.Data[i].Id)
					}
				} else {
					cmd.Println("No app found")
				}
			} else {
				cmd.Println(err)
			}
		}
	},
}
var deleteAppCmd = &cobra.Command{
	Use:   "app [APP_NAME] [BEARER TOKEN]",
	Short: "delete an app",
	Run: func(cmd *cobra.Command, args []string) {
		token := args[1]
		appName := args[0]
		auth := context.WithValue(context.Background(), openapi.ContextAccessToken, token)
		res, _, err := client.AppsApi.DeleteAppByUser(auth, appName)
		if res.Success {
			cmd.Println("App deleted successfully")
		} else {
			cmd.Println(err)
		}
	},
}
