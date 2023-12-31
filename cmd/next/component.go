package next

import (
	"fmt"
	"github.com/mir-mirsodikov/vulcan/models"
	"github.com/mir-mirsodikov/vulcan/templates"
	"github.com/spf13/cobra"
	"os"
	"text/template"
)

var componentCmd = &cobra.Command{
	Use:     "component [name]",
	Aliases: []string{"c"},
	Short:   "Add a new component",
	Long:    `Add a new component`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("Please provide a name for the page.")
			os.Exit(1)
		}

		isClientComponent, _ := cmd.Flags().GetBool("client-component")
		desiredPath, _ := cmd.Flags().GetString("path")

		componentOpts := models.Component{
			Name:            args[0],
			ClientComponent: isClientComponent,
		}

		componentOpts.SetPath(desiredPath)

		fmt.Println("Creating a new component named:", componentOpts.Name, " ...")

		rawTmpl, _ := templates.TemplateFiles.ReadFile("next_component.tmpl")
		tmpl, err := template.New("next_component").Parse(string(rawTmpl))

		if err != nil {
			fmt.Println("Error parsing template:", err)
			os.Exit(1)
		}

		err = os.MkdirAll(componentOpts.Path, 0755)

		if err != nil {
			fmt.Println("Error creating directory:", err)
			os.Exit(1)
		}

		newFile, err := os.Create(componentOpts.GetComponentPath())
		if err != nil {
			fmt.Println("Error creating file:", err)
			os.Exit(1)
		}
		defer newFile.Close()

		err = tmpl.Execute(newFile, componentOpts)

		if err != nil {
			fmt.Println("Error executing template:", err)
			os.Exit(1)
		}

		fmt.Println("Successfully created component:", componentOpts.Name)
	},
}

func init() {
	addCmd.AddCommand(componentCmd)
}
