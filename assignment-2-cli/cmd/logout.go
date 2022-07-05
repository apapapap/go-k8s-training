/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/apapapap/go-k8s-training/assignment2/journal/user"
	"github.com/spf13/cobra"
)

// logoutCmd represents the logout command
var logoutCmd = &cobra.Command{
	Use:   "logout",
	Short: "Logout current logged-in user",
	Long:  "Initiate logout operation for the current logged-in user from the journal application",
	Run: func(cmd *cobra.Command, args []string) {
		user.Logout()
	},
}

func init() {
	userCmd.AddCommand(logoutCmd)
}
