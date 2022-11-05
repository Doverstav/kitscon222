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

// deleteConferenceCmd represents the deleteConference command
var deleteConferenceCmd = &cobra.Command{
	Aliases: []string{"c", "conf"},
	Use:     "conference",
	Short:   "Deletes a conference and all presentations under it",
	Long: `Given a conference name, will remove that conference and 
all presentations added under it`,
	Run: func(cmd *cobra.Command, args []string) {
		confName := cmd.Flag("conference").Value.String()

		kitsconToDelete, err := database.GetKitsconByName(db, confName)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		err = database.DeleteKitscon(db, kitsconToDelete.Id, kitsconToDelete.PresentationIds)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		fmt.Printf("Deleted conference %s and all presentations", kitsconToDelete.Name)
	},
}

func init() {
	deleteCmd.AddCommand(deleteConferenceCmd)

	deleteConferenceCmd.Flags().StringP("conference", "c", "", "Name of conference to delete")
	deleteConferenceCmd.MarkFlagRequired("conference")
}
