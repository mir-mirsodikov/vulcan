package cmd

import (
	"fmt"
	"github.com/mir-mirsodikov/vulcan/cmd/next"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd *cobra.Command

func NewRootCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "vulcan",
		Short: "Vulcan is a tool to help speed up your JS meta framework work.",
		Long:  `A code generator for your favorite JS meta frameworks like Next, Nuxt, SvelteKit, etc. Quickly add new pages, components, routes, etc.`,
		Run: func(cmd *cobra.Command, args []string) {
			_, err := fmt.Fprintf(cmd.OutOrStdout(), "Hello, World!\n")
			if err != nil {
				return
			}
		},
	}
}

func init() {
	rootCmd = NewRootCommand()
	rootCmd.AddCommand(next.Command())
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
