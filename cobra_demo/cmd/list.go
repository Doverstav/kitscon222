/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strings"

	"github.com/doverstav/kitscon222/cobra_demo/database"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Args: cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		kitscons := database.GetKitscons(db)
		toPrint := ""

		for _, kitscon := range kitscons {
			toPrint += fmt.Sprintf("### %s ###\n%s\n", kitscon.Name, kitscon.Desc)
			if len(kitscon.PresentationIds) == 0 {
				toPrint += "\n\tNo presentations added\n"
			}

			presentations, _ := database.GetPresentations(db, kitscon.Id)
			for _, presentation := range presentations {
				tempString := fmt.Sprintf("\n### %s by %s ###\n%s\n%s\n%s\n",
					presentation.PresentationTitle,
					presentation.Presenter,
					presentation.Desc,
					strings.Repeat("⭐", presentation.Rating),
					presentation.Review)

				toPrint += strings.ReplaceAll(tempString, "\n", "\n\t")
			}

			toPrint += "\n"
		}

		fmt.Println(toPrint)
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
