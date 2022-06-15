package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"os/user"
	"strings"
)

func main() {
	username := getCurrentUser()
	hostname := getHostname()
	pwd := getPwd()
	fmt.Printf("%s@%s:~%s$ ", username, hostname, pwd)

	reader := bufio.NewReader(os.Stdin)
	cmdStr, err := reader.ReadString('\n')
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	runCommand(cmdStr)
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

	switch cmdStr {
	case "exit":
		os.Exit(0)
	default:
	}

	fmt.Print("\n")
	cmd := exec.Command(cmdStr)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	if err != nil {
		errMsg := fmt.Sprintf("Error running command: %s\n%v", cmdStr, err)
		fmt.Fprintln(os.Stderr, errMsg)
	}
}
