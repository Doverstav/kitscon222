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
		confName := cmd.Flag("conf").Value.String()

		presentations := database.GetPresentations(db, confName)

		toPrint := ""
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
	listPresentationsCmd.Flags().StringP("conf", "c", "", "Conference during which the presentation was held")
	listPresentationsCmd.MarkFlagRequired("conf")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// presentationsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// presentationsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
