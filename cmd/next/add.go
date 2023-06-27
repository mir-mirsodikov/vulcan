package next

import (
	"fmt"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add [page|component|route]",
	Short: "Add a new page, component, or route.",
	Long:  `Add a new page, component, or route.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Unknown command. Please use `vulcan next add --help` for more information.")
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
