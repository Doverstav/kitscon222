/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/AlecAivazis/survey/v2"
	"github.com/doverstav/kitscon222/cobra_demo/database"
	"github.com/spf13/cobra"
)

// addPresentationCmd represents the presentation command
var addPresentationCmd = &cobra.Command{
	Aliases: []string{"p", "pres"},
	Use:     "presentation",
	Short:   "Add presentation under a KitsCon",
	Long: `Given a conference name, will add a presentation 
under that conference.
	
Accepts title, presenter, description, rating & review as arguments. 
If some or all of these are missing, the missing arguments will be 
prompted for.`,
	Run: func(cmd *cobra.Command, args []string) {
		confName := cmd.Flag("conference").Value.String()

		presTitle := ""
		if len(args) >= 1 {
			presTitle = args[0]
		} else {
			survey.AskOne(&survey.Input{
				Message: "Title:",
			}, &presTitle)
			fmt.Println(presTitle)
		}

		presPresenter := ""
		if len(args) >= 2 {
			presPresenter = args[1]
		} else {
			survey.AskOne(&survey.Input{
				Message: "Presenter:",
			}, &presPresenter)
			fmt.Println(presPresenter)
		}

		presDesc := ""
		if len(args) >= 3 {
			presDesc = args[2]
		} else {
			survey.AskOne(&survey.Input{
				Message: "Description:",
			}, &presDesc)
			fmt.Println(presDesc)
		}

		presRating := 0
		if len(args) >= 4 {
			var err error
			presRating, err = strconv.Atoi(args[3])
			if err != nil {
				fmt.Printf("Could not convert rating: %v", err)
				os.Exit(1)
			}
		} else {
			survey.AskOne(&survey.Input{
				Message: "Rating (any positive number):",
			}, &presRating)
			fmt.Println(presRating)
		}

		presReview := ""
		if len(args) >= 5 {
			presReview = args[4]
		} else {
			survey.AskOne(&survey.Input{
				Message: "Review:",
			}, &presReview)
			fmt.Println(presReview)
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

	addPresentationCmd.Flags().StringP("conference", "c", "", "Conference during which the presentation was held")
	addPresentationCmd.MarkFlagRequired("conference")
}
