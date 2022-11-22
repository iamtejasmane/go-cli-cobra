/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package info

import (
	"github.com/spf13/cobra"
)

// info/infoCmd represents the info/info command
var InfoCmd = &cobra.Command{
	Use:   "info",
	Short: "All things information",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// info/infoCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// info/infoCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
