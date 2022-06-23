/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/apapapap/go-k8s-training/assignment2/journal/utils"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "journal entry add {TEXT-TO-ADD}",
	Long:  "journal entry add {TEXT-TO-ADD}",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		journalText := args[0]

		if journalText != "" {
			homeDir, err := os.UserHomeDir()
			utils.CheckErr(err)

			userSession := homeDir + "/journal/.session"
			currentUser := utils.ReadFromFile(userSession)

			journalFile := homeDir + "/journal/" + currentUser + "/entry.txt"
			f, err := os.OpenFile(journalFile, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
			utils.CheckErr(err)
			defer f.Close()

			journalText = time.Now().Round(0).String() + ": " + journalText
			encText, err := utils.Encrypt(journalText)
			utils.CheckErr(err)
			_, err = f.WriteString(encText + "\n")
			utils.CheckErr(err)
		} else {
			fmt.Println("Please enter some text to add to your journal")
		}
	},
}

func init() {
	entryCmd.AddCommand(addCmd)
}
