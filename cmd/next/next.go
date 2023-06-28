package next

import (
	"fmt"
	"github.com/mir-mirsodikov/vulcan/models"
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

func init() {
	pageVars := models.Component{}
	var path string
	addCmd.PersistentFlags().BoolVarP(&pageVars.ServerComponent, "server-component", "s", false, "Create the page as a server component")
	addCmd.PersistentFlags().StringVarP(&path, "path", "p", "", "Specify the path for the page")
}

func Command() *cobra.Command {
	return rootCmd
}
