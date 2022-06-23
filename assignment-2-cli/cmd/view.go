/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"os"

	"github.com/apapapap/go-k8s-training/assignment2/journal/utils"
	"github.com/spf13/cobra"
)

// viewCmd represents the view command
var viewCmd = &cobra.Command{
	Use:   "view",
	Short: "View journal records",
	Long:  "View your journal records limited to the latest 50 entries",
	Run: func(cmd *cobra.Command, args []string) {
		homeDir, err := os.UserHomeDir()
		utils.CheckErr(err)

		userSession := homeDir + "/journal/.session"
		currentUser := utils.ReadFromFile(userSession)

		path := homeDir + "/journal/" + currentUser + "/entry.txt"
		contentArr := utils.ReadFromFileAsSlice(path)
		contentLen := len(contentArr)
		startIndex := contentLen - 50
		if startIndex > 0 {
			fmt.Println(startIndex)
			contentArr = contentArr[startIndex:]
		}
		utils.PrintSlice(contentArr)
	},
}

func init() {
	entryCmd.AddCommand(viewCmd)
}
