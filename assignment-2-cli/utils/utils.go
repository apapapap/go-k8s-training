package utils

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func ReadFromFile(path string) string {
	data, err := ioutil.ReadFile(path)
	if os.IsNotExist(err) {
		fmt.Println("Please login first")
		os.Exit(1)
	}
	CheckErr(err)

	dataStr, err := Decrypt(string(data))
	CheckErr(err)
	return dataStr
}

func ReadFromFileAsSlice(path string) []string {
	data, err := ioutil.ReadFile(path)
	if os.IsNotExist(err) {
		fmt.Println("No journal entries found")
		return []string{}
	}
	CheckErr(err)

	lines := strings.Split(string(data), "\n")
	for i := 0; i < len(lines)-1; i++ {
		lines[i], err = Decrypt(lines[i])
		CheckErr(err)
	}
	return lines[:len(lines)-1]
}

func CheckErr(e error) {
	if e != nil {
		panic(e)
	}
}

func PrintSlice(s []string) {
	for _, value := range s {
		fmt.Printf("%v\n", value)
	}
}

func Exists(name string) (bool, error) {
	_, err := os.Stat(name)
	if err == nil {
		return true, nil
	}
	if errors.Is(err, os.ErrNotExist) {
		return false, nil
	}
	return false, err
}
