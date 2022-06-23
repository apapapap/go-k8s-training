/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/apapapap/go-k8s-training/assignment2/journal/user"
	"github.com/spf13/cobra"
)

// loginCmd represents the login command
var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "User login to journal application",
	Long:  "Initiate login operation for a user to journal application",
	Run: func(cmd *cobra.Command, args []string) {
		user.Login()
	},
}

func init() {
	userCmd.AddCommand(loginCmd)
}
