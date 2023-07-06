package cmd

import (
	"fmt"
	"os"

	"github.com/mir-mirsodikov/vulcan/cmd/next"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd *cobra.Command

func NewRootCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "vulcan",
		Short: "Vulcan is a tool to help speed up your JS meta framework work.",
		Long:  `A code generator for your favorite JS meta frameworks like Next, Nuxt, SvelteKit, etc. Quickly add new pages, components, routes, etc.`,
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd = NewRootCommand()
	rootCmd.AddCommand(next.Command())
	rootCmd.AddCommand(initCmd)
}

func initConfig() {
	viper.SetConfigName("vulcan")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	nextConfig := map[string]interface{}{
		"jsx": false,
	}

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
		}
	}

	viper.SetDefault("next", nextConfig)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
