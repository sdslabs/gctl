package cmd

import (
	"context"
	_context "context"
	_nethttp "net/http"
	"strconv"
	"strings"

	"github.com/antihax/optional"
	openapi "github.com/sdslabs/gctl/client"
	"github.com/sdslabs/gctl/cmd/middlewares"
	"github.com/spf13/cobra"
)

type AppsAPIService interface {
	CreateApp(ctx _context.Context, language string, application openapi.Application) (openapi.InlineResponse2002, *_nethttp.Response, error)
	DeleteAppByUser(ctx _context.Context, app string) (openapi.InlineResponse2002, *_nethttp.Response, error)
	FetchAppByUser(ctx _context.Context, app string) (openapi.InlineResponse2003, *_nethttp.Response, error)
	FetchAppsByUser(ctx _context.Context) (openapi.InlineResponse2003, *_nethttp.Response, error)
	FetchLogsByUser(ctx _context.Context, app string, localVarOptionals *openapi.FetchLogsByUserOpts) (openapi.InlineResponse2005, *_nethttp.Response, error)
	RebuildAppByUser(ctx _context.Context, app string) (openapi.InlineResponse2002, *_nethttp.Response, error)
	UpdateAppByUser(ctx _context.Context, app string, application openapi.Application) (openapi.InlineResponse2002, *_nethttp.Response, error)
}

var appName string
var appsAPIService AppsAPIService = client.AppsApi

func init() {
	createCmd.AddCommand(CreateAppCmd(appsAPIService))
	fetchCmd.AddCommand(FetchAppCmd(appsAPIService))
	deleteCmd.AddCommand(DeleteAppCmd(appsAPIService))
	rootCmd.AddCommand(RebuildAppCmd(appsAPIService))
	updateCmd.AddCommand(UpdateAppCmd(appsAPIService))
	fetchCmd.AddCommand(FetchLogsCmd(appsAPIService))
}

//CreateAppCmd is command to create an app
func CreateAppCmd(appsAPIService AppsAPIService) *cobra.Command {
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
			res, _, err := appsAPIService.CreateApp(auth, language, application)
			if res.Success {
				res, _, err := appsAPIService.FetchAppByUser(auth, application.Name)
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
func FetchAppCmd(appsAPIService AppsAPIService) *cobra.Command {
	var fetchAppCmd = &cobra.Command{
		Use:   "app [BEARER_TOKEN]",
		Short: "fetch apps",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			appName, _ := cmd.Flags().GetString("name")
			token := args[0]
			auth := context.WithValue(context.Background(), openapi.ContextAccessToken, token)
			if appName != "" {
				res, _, err := appsAPIService.FetchAppByUser(auth, appName)
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
					cmd.Println("Error in fetching the app.", err)
				}
			} else {
				res, _, err := appsAPIService.FetchAppsByUser(auth)
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
					cmd.Println("Error in fetching the apps.", err)
				}
			}
		},
	}
	fetchAppCmd.Flags().StringVarP(&appName, "name", "n", "", "show specific app")
	return fetchAppCmd
}

//DeleteAppCmd returns command to delete app owned by a user
func DeleteAppCmd(appsAPIService AppsAPIService) *cobra.Command {
	var deleteAppCmd = &cobra.Command{
		Use:   "app [APP_NAME] [BEARER TOKEN]",
		Short: "delete an app",
		Args:  cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			token := args[1]
			appName := args[0]
			auth := context.WithValue(context.Background(), openapi.ContextAccessToken, token)
			res, _, err := appsAPIService.DeleteAppByUser(auth, appName)
			if res.Success {
				cmd.Println("App deleted successfully")
			} else {
				cmd.Println(err)
			}
		},
	}
	return deleteAppCmd
}

//RebuildAppCmd returns a command to rebuild an app
func RebuildAppCmd(appsAPIService AppsAPIService) *cobra.Command {
	var rebuildAppCmd = &cobra.Command{
		Use:   "rebuild [APP_NAME] [TOKEN]",
		Short: "rebuild an app",
		Args:  cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			token := args[1]
			appName := args[0]
			auth := context.WithValue(context.Background(), openapi.ContextAccessToken, token)
			res, _, err := appsAPIService.RebuildAppByUser(auth, appName)
			if res.Success {
				cmd.Println("App rebuilt successfully")
			} else {
				cmd.Println(err)
			}
		},
	}
	return rebuildAppCmd
}

//UpdateAppCmd returns a command to update an app
func UpdateAppCmd(appsAPIService AppsAPIService) *cobra.Command {
	var updateAppCmd = &cobra.Command{
		Use:   "app [APP_NAME] [FILE_NAME] [TOKEN]",
		Short: "update an app",
		Args:  cobra.MaximumNArgs(3),
		Run: func(cmd *cobra.Command, args []string) {
			var token string
			var application openapi.Application
			if len(args) == 3 {
				token = args[2]
				filename := strings.Split(args[1], ".")[0]
				appName = args[0]
				application = middlewares.ReadAppJSON(filename)
			} else {
				token, _, application = middlewares.AppForm()
				appName = application.Name
			}
			auth := context.WithValue(context.Background(), openapi.ContextAccessToken, token)
			res, _, err := appsAPIService.UpdateAppByUser(auth, appName, application)
			if res.Success {
				res, _, err := appsAPIService.FetchAppByUser(auth, application.Name)
				if res.Success {
					for i := 0; i < len(res.Data); i++ {
						cmd.Println("App updated successfully"+"\n"+"Container Id: "+res.Data[i].ContainerId, "Container Port: ", res.Data[i].ContainerPort,
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
	return updateAppCmd
}

//FetchLogsCmd returns a command to fetch logs of an app
func FetchLogsCmd(appsAPIService AppsAPIService) *cobra.Command {
	var fetchLogsCmd = &cobra.Command{
		Use:   "logs [APP_NAME] [TOKEN] [NUMBER_OF_LOGS] ",
		Short: "fetch logs of an app",
		Args:  cobra.MinimumNArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			var localVarOptional openapi.FetchLogsByUserOpts
			token := args[1]
			appName := args[0]
			if len(args) == 3 {
				n, _ := strconv.ParseInt(args[2], 10, 32)
				localVarOptional = openapi.FetchLogsByUserOpts{
					Tail: optional.NewInt32(int32(n)),
				}
			}
			auth := context.WithValue(context.Background(), openapi.ContextAccessToken, token)
			res, _, err := appsAPIService.FetchLogsByUser(auth, appName, &localVarOptional)
			if res.Success {
				if len(res.Data) != 0 {
					for i := 0; i < len(res.Data); i++ {
						cmd.Println(res.Data[i])
					}
				} else {
					cmd.Print("No logs found")
				}
			} else {
				cmd.Println("Error in fetching the logs.", err)
			}
		},
	}
	return fetchLogsCmd
}
