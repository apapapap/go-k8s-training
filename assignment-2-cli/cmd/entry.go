/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/spf13/cobra"
)

// entryCmd represents the entry command
var entryCmd = &cobra.Command{
	Use:   "entry",
	Short: "Add/View entries in your journal",
	Long:  "Add/View entries in your journal",
}

func init() {
	rootCmd.AddCommand(entryCmd)
}
