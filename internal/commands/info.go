/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	// buildVersion - версия сборки.
	BuildVersion string
	// buildDate - дата сборки.
	BuildDate string
	// buildCommit - комментарии к сборке.
	BuildCommit string
)

// infoCmd represents the info command
var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "Версия утилиты",
	Long:  `При вызове отображается версия утилиты, номер сборки, и время сборки.`,
	Run: func(cmd *cobra.Command, args []string) {
		if BuildVersion == "" {
			BuildVersion = "N/A"
		}
		if BuildDate == "" {
			BuildDate = "N/A"
		}
		if BuildCommit == "" {
			BuildCommit = "N/A"
		}
		fmt.Printf("\nGophKeeper \nSimple cli utility for storing passwords, bank card data, text and binary data.\nBuild: %s \nCommit: %s \nBuild Time: %s\n", BuildVersion, BuildCommit, BuildDate)
	},
}

func init() {
	rootCmd.AddCommand(infoCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// infoCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// infoCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
