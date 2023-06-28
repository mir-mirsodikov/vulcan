package next

import (
	"fmt"
	"github.com/mir-mirsodikov/vulcan/pkg/chalk"
	"github.com/mir-mirsodikov/vulcan/pkg/next"
	"github.com/mir-mirsodikov/vulcan/templates"
	"github.com/spf13/cobra"
	"os"
	"text/template"
)

var pageCmd = &cobra.Command{
	Use:     "page [name]",
	Aliases: []string{"p"},
	Short:   "Add a new page",
	Long:    `Add a new page`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("Please provide a name for the page.")
			os.Exit(1)
		}

		isServerComponent, _ := cmd.Flags().GetBool("server-component")
		desiredPath, _ := cmd.Flags().GetString("path")

		pageOpts := next.Page{
			Name:            args[0],
			ServerComponent: isServerComponent,
		}

		pageOpts.SetPath(desiredPath)

		fmt.Println("Creating a new page named:", pageOpts.Name, " ...")

		rawTmpl, _ := templates.TemplateFiles.ReadFile("next_page.tmpl")

		tmpl, err := template.New("next_page").Parse(string(rawTmpl))

		if err != nil {
			fmt.Println("Error parsing template:", err)
			os.Exit(1)
		}

		err = os.MkdirAll(pageOpts.Path, 0755)

		if err != nil {
			fmt.Println("Error creating directory:", err)
			os.Exit(1)
		}

		newFile, err := os.Create(pageOpts.GetPagePath())
		if err != nil {
			fmt.Println("Error creating file:", err)
			os.Exit(1)
		}
		defer newFile.Close()

		err = tmpl.Execute(newFile, pageOpts)

		if err != nil {
			fmt.Println("Error executing template:", err)
			os.Exit(1)
		}

		fmt.Println(chalk.Green("`" + pageOpts.Name + "` page created!"))
	},
}

func init() {
	addCmd.AddCommand(pageCmd)
}
