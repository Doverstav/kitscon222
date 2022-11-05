/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/doverstav/kitscon222/cobra_demo/database"
	"github.com/spf13/cobra"
)

// listConfsCmd represents the cons command
var listConfsCmd = &cobra.Command{
	Aliases: []string{"c", "confs"},
	Use:     "conferences",
	Short:   "List all conferences",
	Run: func(cmd *cobra.Command, args []string) {
		kitscons := database.GetKitscons(db)

		toPrint := ""
		for _, kitscon := range kitscons {
			toPrint += fmt.Sprintf("=== %s ===\n%s\n\n", kitscon.Name, kitscon.Desc)
		}

		fmt.Println(toPrint)
	},
}

func init() {
	listCmd.AddCommand(listConfsCmd)
}
