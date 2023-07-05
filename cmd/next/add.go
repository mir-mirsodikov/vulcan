package next

import (
	"fmt"
	"github.com/mir-mirsodikov/vulcan/models"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var addCmd *cobra.Command

func NewAddCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "add [page|component|route]",
		Short: "Add a new page, component, or route.",
		Long:  `Add a new page, component, or route.`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Unknown command. Please use `vulcan next add --help` for more information.")
		},
	}
}

func BuildAddCommand(cmd *cobra.Command) {
	pageVars := models.Component{}
	var path string
	var jsx bool
	cmd.PersistentFlags().BoolVarP(&pageVars.ServerComponent, "server-component", "s", false, "Create the page as a server component")
	cmd.PersistentFlags().StringVarP(&path, "path", "p", "", "Specify the path for the page")
	cmd.PersistentFlags().BoolVar(&jsx, "jsx", true, "Create the component/page using JSX")
	viper.BindPFlag("next.jsx", cmd.Flags().Lookup("jsx"))

}

func init() {
	addCmd = NewAddCommand()
	rootCmd.AddCommand(addCmd)
}
