package user

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/apapapap/go-k8s-training/assignment2/journal/utils"
	"golang.org/x/crypto/ssh/terminal"
)

func Login() {
	fmt.Println("Welcome to your Journal application")
	fmt.Println("Please enter your username: ")
	username := fetchStringInputFromTerminal()

	homeDir, err := os.UserHomeDir()
	utils.CheckErr(err)

	appDirName := homeDir + "/journal/" + username
	userExists, err := utils.Exists(appDirName + "/password.txt")
	utils.CheckErr(err)
	if !userExists {
		fmt.Printf("User '%s' doesn't exists, creating new...\n", username)

		fmt.Println("Please set password for new user:")
		password := fetchPasswordFromTerminal()
		pwdFilename := appDirName + "/password.txt"
		createNewUser(pwdFilename, username, password, homeDir, appDirName)
	} else {
		fmt.Println("Please enter password for user:")
		password := fetchPasswordFromTerminal()
		pwdFilename := appDirName + "/password.txt"
		loginExistingUser(pwdFilename, username, password, homeDir)
	}

	userLogin := homeDir + "/journal/.session"
	sessionFile, sessionErr := os.Create(userLogin)
	utils.CheckErr(sessionErr)
	defer sessionFile.Close()

	encUser, err := utils.Encrypt(username)
	utils.CheckErr(err)

	_, err = sessionFile.Write([]byte(encUser))
	utils.CheckErr(err)
}

func fetchStringInputFromTerminal() string {
	reader := bufio.NewReader(os.Stdin)
	str, err := reader.ReadString('\n')
	utils.CheckErr(err)
	str = strings.TrimSuffix(str, "\n")
	return str
}

func fetchPasswordFromTerminal() string {
	bytePassword, err := terminal.ReadPassword(0)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	password := strings.TrimSuffix(string(bytePassword), "\n")
	return password
}

func loginExistingUser(pwdFilename, username, password, homeDir string) {
	dat, err := os.ReadFile(pwdFilename)
	utils.CheckErr(err)

	decodedPwd, err := utils.Decrypt(string(dat))
	utils.CheckErr(err)

	if decodedPwd == password {
		fmt.Printf("User %s login successful\n", username)
	} else {
		fmt.Printf("Invalid credentials, login failed for user %s\n", username)
		os.Exit(1)
	}
}

func createNewUser(pwdFilename, username, password, homeDir, appDirName string) {
	err := os.MkdirAll(appDirName, 0755)
	utils.CheckErr(err)

	// we assume that if user is present the password file will always be present
	// hence we don't check if password file exists
	myfile, err := os.Create(pwdFilename)
	utils.CheckErr(err)
	defer myfile.Close()

	encPwd, err := utils.Encrypt(password)
	utils.CheckErr(err)
	_, err = myfile.Write([]byte(encPwd))
	utils.CheckErr(err)

	fmt.Printf("User %s created and login successful\n", username)
}

func Logout() {
	homeDir, err := os.UserHomeDir()
	utils.CheckErr(err)

	userSession := homeDir + "/journal/.session"

	currentUser := utils.ReadFromFile(userSession)

	err = os.Remove(userSession)
	utils.CheckErr(err)

	fmt.Printf("User %s logout successful\n", currentUser)
}
