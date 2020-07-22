package middlewares

import (
	"os"

	openapi "github.com/sdslabs/gctl/client"
	"github.com/spf13/viper"
)

//ReadAppJSON reads a json file for app details
func ReadAppJSON(filename string) (openapi.Application, error) {
	var application openapi.Application
	viper.SetConfigName(filename)
	viper.SetConfigType("json")
	path, _ := os.Getwd()
	viper.AddConfigPath(path)
	err := viper.ReadInConfig()
	if err != nil {
		return application, err
	}
	err = viper.Unmarshal(&application)
	if err != nil {
		return application, err
	}
	if viper.IsSet("git.repo_url") {
		application.Git.RepoUrl = viper.GetString("git.repo_url")
	}
	if viper.IsSet("git.access_token") {
		application.Git.AccessToken = viper.GetString("git.access_token")
	}
	if viper.IsSet("context.rc_file") {
		application.Context.RcFile = viper.GetBool("context.rc_file")
	} else {
		application.Context.RcFile = false
	}
	return application, nil
}
