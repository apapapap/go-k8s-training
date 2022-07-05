/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/spf13/cobra"
)

// userCmd represents the user command
var userCmd = &cobra.Command{
	Use:   "user",
	Short: "Perform operations for a user",
	Long:  "Perform operations for a user like login, logout",
}

func init() {
	rootCmd.AddCommand(userCmd)
}
