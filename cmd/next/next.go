package next

import (
	"fmt"
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
		if len(args) < 1 {
			fmt.Println("Please provide a name for the page.")
			os.Exit(1)
		}

		isServerComponent, _ := cmd.Flags().GetBool("server-component")

		pageOpts := page{
			Name:            args[0],
			ServerComponent: isServerComponent,
		}

		fmt.Println("Creating a new page named:", pageOpts.Name, " ...")

		rawTmpl, _ := templates.TemplateFiles.ReadFile("next_page.tmpl")

		tmpl, err := template.New("next_page").Parse(string(rawTmpl))

		if err != nil {
			fmt.Println("Error parsing template:", err)
			os.Exit(1)
		}

		err = os.MkdirAll("app/"+pageOpts.Name, 0755)

		if err != nil {
			fmt.Println("Error creating directory:", err)
			os.Exit(1)
		}

		newFile, err := os.Create("app/" + pageOpts.Name + "/page.jsx")
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

		fmt.Println("Successfully created a new page named:", pageOpts.Name, " ...")
	},
}

type page struct {
	Name            string
	ServerComponent bool
}

func init() {
	pageVars := page{}
	pageCmd.Flags().BoolVarP(&pageVars.ServerComponent, "server-component", "s", false, "Create the page as a server component")

	addCmd.AddCommand(pageCmd)
	rootCmd.AddCommand(addCmd)
}

func Command() *cobra.Command {
	return rootCmd
}
