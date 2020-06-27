package middlewares

import (
	"fmt"
	"os"

	openapi "github.com/sdslabs/gctl/client"
	"github.com/spf13/viper"
)

//ReadAppJSON reads a json file for app details
func ReadAppJSON(filename string) openapi.Application {
	var application openapi.Application
	viper.SetConfigName(filename)
	viper.SetConfigType("json")
	path, _ := os.Getwd()
	viper.AddConfigPath(path)
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %s", err))
	}
	err = viper.Unmarshal(&application)
	if err != nil {
		panic("unable to decode json file into struct")
	}
	return application
}
