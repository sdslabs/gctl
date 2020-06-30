package cmd

import (
	"context"
	"strings"

	openapi "github.com/sdslabs/gctl/client"
	"github.com/sdslabs/gctl/cmd/middlewares"
	"github.com/spf13/cobra"
)

var appName string

func init() {
	createCmd.AddCommand(CreateAppCmd(*client))
	fetchCmd.AddCommand(FetchAppCmd(*client))
	deleteCmd.AddCommand(DeleteAppCmd(*client))
}

//CreateAppCmd is command to create an app
func CreateAppCmd(client openapi.APIClient) *cobra.Command {
	var appmakerCmd = &cobra.Command{
		Use:   "app [TOKEN] [FILENAME] [LANGUAGE]",
		Short: "Create an application",
		Run: func(cmd *cobra.Command, args []string) {
			var token, language string
			var application openapi.Application
			if len(args) == 3 {
				token = args[0]
				filename := strings.Split(args[1], ".")[0]
				language = args[2]
				application = middlewares.ReadAppJSON(filename)
			} else {
				token, language, application = middlewares.AppForm()
			}
			auth := context.WithValue(context.Background(), openapi.ContextAccessToken, token)
			res, _, err := client.AppsApi.CreateApp(auth, language, application)
			if res.Success {
				res, _, err := client.AppsApi.FetchAppByUser(auth, application.Name)
				if res.Success {
					for i := 0; i < len(res.Data); i++ {
						cmd.Println("App created successfully"+"\n"+"Container Id: "+res.Data[i].ContainerId, "Container Port: ", res.Data[i].ContainerPort,
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
	return appmakerCmd
}

//FetchAppCmd returns command to fetch apps of a user
func FetchAppCmd(client openapi.APIClient) *cobra.Command {
	var fetchAppCmd = &cobra.Command{
		Use:   "app [BEARER_TOKEN]",
		Short: "fetch apps",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			appName, _ := cmd.Flags().GetString("name")
			token := args[0]
			auth := context.WithValue(context.Background(), openapi.ContextAccessToken, token)
			if appName != "" {
				res, _, err := client.AppsApi.FetchAppByUser(auth, appName)
				if res.Success {
					if len(res.Data) != 0 {
						for i := 0; i < len(res.Data); i++ {
							cmd.Println("Container Id: "+res.Data[i].ContainerId, "Container Port: ", res.Data[i].ContainerPort,
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
							cmd.Println("Container Id: "+res.Data[i].ContainerId, "Container Port: ", res.Data[i].ContainerPort,
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
	fetchAppCmd.Flags().StringVarP(&appName, "name", "n", "", "show specific app")
	return fetchAppCmd
}

//DeleteAppCmd returns command to delete app owned by a user
func DeleteAppCmd(client openapi.APIClient) *cobra.Command {
	var deleteAppCmd = &cobra.Command{
		Use:   "app [APP_NAME] [BEARER TOKEN]",
		Short: "delete an app",
		Args:  cobra.ExactArgs(2),
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
	return deleteAppCmd
}
