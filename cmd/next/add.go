package next

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var addCmd = &cobra.Command{
	Use:   "add [page|component|route]",
	Short: "Add a new page, component, or route.",
	Long:  `Add a new page, component, or route.`,
}

func init() {
	addCmd.PersistentFlags().BoolP("client-component", "c", false, "Create the page as a client component")
	addCmd.PersistentFlags().StringP("path", "p", "", "Specify the path for the page")
	addCmd.PersistentFlags().Bool("jsx", false, "Create the component/page using JSX")
	viper.BindPFlag("next.jsx", addCmd.PersistentFlags().Lookup("jsx"))
	rootCmd.AddCommand(addCmd)
}
