package next

import (
	"fmt"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "next",
	Short: "Generate code for NextJS",
	Long:  `Generate code for NextJS`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Unknown command. Please use `vulcan next --help` for more information.")
	},
}

func Command() *cobra.Command {
	return rootCmd
}
