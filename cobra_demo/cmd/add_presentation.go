/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/doverstav/kitscon222/cobra_demo/database"
	"github.com/spf13/cobra"
)

// addPresentationCmd represents the presentation command
var addPresentationCmd = &cobra.Command{
	Aliases: []string{"p", "pres"},
	Use:     "presentation",
	Short:   "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		confName := cmd.Flag("conf").Value.String()

		presTitle := ""
		if len(args) >= 1 {
			presTitle = args[0]
		}
		presPresenter := ""
		if len(args) >= 2 {
			presPresenter = args[1]
		}
		presDesc := ""
		if len(args) >= 3 {
			presDesc = args[2]
		}
		presRating := 0
		if len(args) >= 4 {
			var err error
			presRating, err = strconv.Atoi(args[3])
			if err != nil {
				fmt.Printf("Could not convert rating: %v", err)
				os.Exit(1)
			}
		}
		presReview := ""
		if len(args) >= 5 {
			presReview = args[4]
		}

		parentKitscon, err := database.GetKitsconByName(db, confName)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		err = database.SavePresentation(
			db, parentKitscon.Id,
			presTitle,
			presPresenter,
			presDesc,
			presRating,
			presReview,
		)
		if err != nil {
			fmt.Printf("Could not save presentation: %v", err)
			os.Exit(1)
		}

		fmt.Printf("Saved presentation %s by %s", presTitle, presPresenter)
	},
}

func init() {
	addCmd.AddCommand(addPresentationCmd)

	addPresentationCmd.Flags().StringP("conf", "c", "", "Conference during which the presentation was held")
	addPresentationCmd.MarkFlagRequired("conf")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// presentationCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// presentationCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}