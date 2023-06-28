package main

import (
	"fmt"
	"github.com/mir-mirsodikov/vulcan/cmd"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigName("vulcan")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	nextConfig := map[string]any{
		"jsx": false,
	}

	viper.SetDefault("next", nextConfig)

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Println("No config file found. Using defaults.")
		}
	}

	cmd.Execute()
}
