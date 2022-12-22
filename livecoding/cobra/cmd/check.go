/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/spf13/cobra"
)

// checkCmd represents the check command
var checkCmd = &cobra.Command{
	Aliases: []string{"c", "kitscondemo"},
	Use:     "check [url]",
	Short:   "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("check called")
		url := args[0]
		times, _ := strconv.Atoi(cmd.Flag("times").Value.String())

		c := &http.Client{Timeout: 10 * time.Second}
		for times > 0 {
			res, err := c.Get(url)
			if err != nil {
				fmt.Printf("Broke: %v", err)
				os.Exit(1)
			}

			fmt.Println(res.Status)

			times--
		}
	},
}

func init() {
	rootCmd.AddCommand(checkCmd)

	checkCmd.Flags().IntP("times", "t", 1, "How many times to check the url")
	checkCmd.MarkFlagRequired("times")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// checkCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// checkCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
