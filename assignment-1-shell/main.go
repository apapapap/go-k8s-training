package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"os/user"
	"strings"
)

var cmdHistory []string

func main() {
	username := getCurrentUser()
	hostname := getHostname()

	for {
		pwd := getPwd()
		fmt.Printf("%s@%s:~%s$ ", username, hostname, pwd)

		reader := bufio.NewReader(os.Stdin)
		cmdStr, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}

		if cmdStr != "" && cmdStr != "\n" {
			cmdHistory = append(cmdHistory, cmdStr)
		}
		runCommand(cmdStr)
	}
}

func getCurrentUser() string {
	currentUser, err := user.Current()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	return currentUser.Username
}

func getHostname() string {
	hostname, err := os.Hostname()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	return hostname
}

func getPwd() string {
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	return pwd
}

func runCommand(cmdStr string) {
	cmdStr = strings.TrimSuffix(cmdStr, "\n")
	cmdStrArr := strings.Fields(cmdStr)

	if len(cmdStrArr) != 0 {
		switch cmdStrArr[0] {
		case "exit":
			os.Exit(0)
		case "history":
			printHistory()
			return
		case "cd":
			changeDir(cmdStrArr)
			return
		default:
		}

		cmd := exec.Command(cmdStrArr[0], cmdStrArr[1:]...)
		cmd.Stderr = os.Stderr
		cmd.Stdout = os.Stdout
		err := cmd.Run()
		if err != nil {
			errMsg := fmt.Sprintf("Error running command: %s\n%v", cmdStr, err)
			fmt.Fprintln(os.Stderr, errMsg)
		}
	}
}

func printHistory() {
	for i, cmd := range cmdHistory {
		fmt.Printf("%d  %v", i+1, cmd)
	}
}

func changeDir(cmdStrArr []string) {
	newDir := ""
	var err error
	if len(cmdStrArr) > 1 {
		newDir = cmdStrArr[1]
	} else {
		newDir, err = os.UserHomeDir()
		if err != nil {
			errMsg := fmt.Sprintf("Error running command: %v\n%v", cmdStrArr, err)
			fmt.Fprintln(os.Stderr, errMsg)
			return
		}
	}

	err = os.Chdir(newDir)
	if err != nil {
		errMsg := fmt.Sprintf("Error running command: %v\n%v", cmdStrArr, err)
		fmt.Fprintln(os.Stderr, errMsg)
	}
}
