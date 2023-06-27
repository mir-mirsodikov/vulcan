package next

import (
	"fmt"
	"github.com/mir-mirsodikov/vulcan/pkg/chalk"
	"github.com/mir-mirsodikov/vulcan/templates"
	"github.com/spf13/cobra"
	"os"
	"text/template"
)

var rootCmd = &cobra.Command{
	Use:   "next",
	Short: "Generate code for NextJS",
	Long:  `Generate code for NextJS`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Unknown command. Please use `vulcan next --help` for more information.")
	},
}

var addCmd = &cobra.Command{
	Use:   "add [page|component|route]",
	Short: "Add a new page, component, or route.",
	Long:  `Add a new page, component, or route.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Unknown command. Please use `vulcan next add --help` for more information.")
	},
}

var pageCmd = &cobra.Command{
	Use:   "page [name]",
	Short: "Add a new page",
	Long:  `Add a new page`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("Please provide a name for the page.")
			os.Exit(1)
		}

		isServerComponent, _ := cmd.Flags().GetBool("server-component")
		desiredPath, _ := cmd.Flags().GetString("path")

		pageOpts := page{
			Name:            args[0],
			ServerComponent: isServerComponent,
		}

		path := "app/" + pageOpts.Name + "/"

		if desiredPath != "" {
			path = "app/" + desiredPath + "/" + pageOpts.Name + "/"
		}

		fmt.Println("Creating a new page named:", pageOpts.Name, " ...")

		rawTmpl, _ := templates.TemplateFiles.ReadFile("next_page.tmpl")

		tmpl, err := template.New("next_page").Parse(string(rawTmpl))

		if err != nil {
			fmt.Println("Error parsing template:", err)
			os.Exit(1)
		}

		err = os.MkdirAll(path, 0755)

		if err != nil {
			fmt.Println("Error creating directory:", err)
			os.Exit(1)
		}

		newFile, err := os.Create(path + "/page.jsx")
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

		fmt.Println(chalk.Green("Successfully created a new page named:", pageOpts.Name, " ..."))
	},
}

type page struct {
	Name            string
	ServerComponent bool
}

func init() {
	pageVars := page{}
	var path string
	pageCmd.Flags().BoolVarP(&pageVars.ServerComponent, "server-component", "s", false, "Create the page as a server component")
	pageCmd.Flags().StringVarP(&path, "path", "p", "", "Specify the path for the page")

	addCmd.AddCommand(pageCmd)
	rootCmd.AddCommand(addCmd)
}

func Command() *cobra.Command {
	return rootCmd
}
