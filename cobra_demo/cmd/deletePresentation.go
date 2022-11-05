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
	Short:   "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		confName := cmd.Flag("conf").Value.String()
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
	deletePresentationCmd.Flags().StringP("conf", "c", "", "Conference during which the presentation was held")
	deletePresentationCmd.MarkFlagRequired("conf")
	deletePresentationCmd.Flags().StringP("presentation", "p", "", "Presentation you wish to delete")
	deletePresentationCmd.MarkFlagRequired("presentation")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deletePresentationCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deletePresentationCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
