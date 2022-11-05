/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/doverstav/kitscon222/cobra_demo/database"
	"github.com/spf13/cobra"
)

// listPresentationsCmd represents the presentations command
var listPresentationsCmd = &cobra.Command{
	Aliases: []string{"p", "pres"},
	Use:     "presentations",
	Short:   "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		confName := cmd.Flag("conference").Value.String()

		parentKitscon, err := database.GetKitsconByName(db, confName)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		presentations, err := database.GetPresentations(db, parentKitscon.Id)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		toPrint := fmt.Sprintf("=== %s presentations ===\n\n", parentKitscon.Name)
		for _, presentation := range presentations {
			toPrint += fmt.Sprintf("### %s by %s ###\n%s\n%s\n%s\n\n",
				presentation.PresentationTitle,
				presentation.Presenter,
				presentation.Desc,
				strings.Repeat("⭐", presentation.Rating),
				presentation.Review)
		}

		fmt.Println(toPrint)
	},
}

func init() {
	listCmd.AddCommand(listPresentationsCmd)

	listPresentationsCmd.Flags().StringP("conference", "c", "", "Conference during which the presentation was held")
	listPresentationsCmd.MarkFlagRequired("conference")
}
