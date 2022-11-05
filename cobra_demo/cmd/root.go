/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/dgraph-io/badger/v3"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
)

var db *badger.DB

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "kitscon-cobra",
	Short: "A CLI for managing your KitsCon experience",
	Long: `Have you always had a hard time keeping track of 
what you listened to at KitsCon, and what you thought about it?
	
Fret no more, for the KitsCon CLI (built with Cobra) will help 
you keep track of every KitsCon, presentations you've attended 
and review you've written!`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	// Get app dir
	dir, _ := homedir.Dir()
	expandedDir, _ := homedir.Expand(dir)
	appDir := expandedDir + "/kitscon-cli"
	fmt.Println(appDir)

	// Setup db
	var err error
	db, err = badger.Open(badger.DefaultOptions(appDir).WithLoggingLevel(badger.ERROR))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
