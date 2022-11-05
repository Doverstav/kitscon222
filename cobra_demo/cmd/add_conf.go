/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/AlecAivazis/survey/v2"
	"github.com/doverstav/kitscon222/cobra_demo/database"
	"github.com/spf13/cobra"
)

// addConfCmd represents the conf command
var addConfCmd = &cobra.Command{
	Aliases: []string{"c", "conf"},
	Use:     "conference",
	Short:   "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		confName := ""
		if len(args) >= 1 {
			confName = args[0]
		} else {
			survey.AskOne(&survey.Input{
				Message: "Name:",
			}, &confName)
			fmt.Println(confName)
		}
		confDesc := ""
		if len(args) >= 2 {
			confDesc = args[1]
		} else {
			survey.AskOne(&survey.Input{
				Message: "Description",
			}, &confDesc)
			fmt.Println(confDesc)
		}

		err := database.SaveKitscon(db, confName, confDesc)
		if err != nil {
			fmt.Printf("Error when saving new conference: %v", err)
			os.Exit(1)
		}

		fmt.Printf("Saved KitsCon with name %s", confName)
	},
}

func init() {
	addCmd.AddCommand(addConfCmd)
}
