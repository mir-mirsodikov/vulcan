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

	jsx := false
	cmd.PersistentFlags().BoolVar(&jsx, "jsx", true, "Create the component/page using JSX")
	viper.BindPFlag("next.jsx", addCmd.Flags().Lookup("jsx"))

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Println("No config file found. Using defaults.")
		}
	}

	cmd.Execute()
}
