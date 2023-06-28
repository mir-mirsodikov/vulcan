package cmd

import (
	"fmt"
	"github.com/mir-mirsodikov/vulcan/configs"
	"github.com/spf13/cobra"
	"os"
)

var initCmd = &cobra.Command{
	Use:   "init [framework]",
	Short: "Initialize a new project",
	Long:  `Initialize a new project`,
	Run: func(cmd *cobra.Command, args []string) {
		framework := args[0]

		switch framework {
		case "next":
			fmt.Println("Initializing a new Next.js project...")
			file, err := os.Create("vulcan.yaml")

			config, _ := configs.ConfigFiles.ReadFile("next.yaml")

			if err != nil {
				fmt.Println("Error creating vulcan.yaml file:", err)
				os.Exit(1)
			}

			defer file.Close()

			file.WriteString(string(config))

			fmt.Println("Done!")
			break
		default:
			fmt.Println("Please provide a valid framework.")
			break
		}

	},
}
