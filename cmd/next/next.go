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
		fmt.Println("Hello, World from the Next handler!")
	},
}

func NextCommand() *cobra.Command {
	return rootCmd
}
