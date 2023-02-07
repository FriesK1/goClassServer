package main

import (
	config "github.com/FriesK1/goMartialConfig"

	"github.com/sirupsen/logrus"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func init() {
	config.SetAppName("myServerApp", "msa")
	config.SetAppVersion("1.0.0")

	viper.SetDefault("server.address", "0.0.0.0")
	pflag.StringP("server.address", "a", viper.GetString("server.address"), "Address to bind HTTP Server to")

	viper.SetDefault("server.port", 3001)
	pflag.IntP("server.port", "p", viper.GetInt("server.port"), "Port to bind HTTP Server to")
}

func main() {
	config.GetConfigs()

	if err := StartWebServer(); err != nil {
		logrus.Fatal(err)
	}
}
