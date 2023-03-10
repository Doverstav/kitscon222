/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/doverstav/kitscon222/cobra_demo/database"
	"github.com/spf13/cobra"
)

// deletePresentationCmd represents the deletePresentation command
var deletePresentationCmd = &cobra.Command{
	Aliases: []string{"p", "pres"},
	Use:     "presentation",
	Short:   "Deletes a presentation under a KitsCon",
	Long: `Given a presentation name and a conference name, 
will remove that presentation from that conference.`,
	Run: func(cmd *cobra.Command, args []string) {
		confName := cmd.Flag("conference").Value.String()
		presName := cmd.Flag("presentation").Value.String()

		fmt.Println("deletePresentation called")
		parentKitscon, err := database.GetKitsconByName(db, confName)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		presToDelete, err := database.GetPresentationByName(db, parentKitscon.Id, presName)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		database.RemovePresentation(db, parentKitscon.Id, presToDelete.Id)
	},
}

func init() {
	deleteCmd.AddCommand(deletePresentationCmd)

	deletePresentationCmd.Flags().StringP("conference", "c", "", "Conference during which the presentation was held")
	deletePresentationCmd.MarkFlagRequired("conference")
	deletePresentationCmd.Flags().StringP("presentation", "p", "", "Presentation you wish to delete")
	deletePresentationCmd.MarkFlagRequired("presentation")
}
