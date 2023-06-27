package next

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "next",
	Short: "Generate code for NextJS",
	Long:  `Generate code for NextJS`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hello, World from the Next handler!")
	},
}

var addCmd = &cobra.Command{
	Use:   "add [page|component|route]",
	Short: "Add a new page, component, route, etc.",
	Long:  `Add a new page, component, route, etc.`,
	Run: func(cmd *cobra.Command, args []string) {
		cwd, _ := os.Getwd()
		fmt.Println("Hello, World from the Next add handler!")
		fmt.Println("Current working directory:", cwd)
	},
}

var pageCmd = &cobra.Command{
	Use:   "page [name]",
	Short: "Add a new page",
	Long:  `Add a new page`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hello, World from the Next page handler!")

		name := args[0]
		if name == "" {
			fmt.Println("Please provide a name for the page.")
			os.Exit(1)
		}

		newFile, err := os.Create(name)
		if err != nil {
			fmt.Println("Error creating file:", err)
			os.Exit(1)
		}
		defer newFile.Close()
	},
}

func init() {
	pageCmd.Flags().StringP("name", "n", "", "Name of the page to add.")

	addCmd.AddCommand(pageCmd)
	rootCmd.AddCommand(addCmd)
}

func NextCommand() *cobra.Command {
	return rootCmd
}
