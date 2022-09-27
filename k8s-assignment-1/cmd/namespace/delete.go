/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package namespace

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

// deleteCmd represents the delete namespace command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a namespace",
	Long:  "kube-client delete namespace",
	Run: func(cmd *cobra.Command, args []string) {
		var config *rest.Config
		fmt.Println("Creating in-cluster config")
		config, err := rest.InClusterConfig()
		if err != nil {
			fmt.Println("Failed to create in-cluster config, trying to fetch from global kube config")
			kubeConfigFilepath := filepath.Join(
				os.Getenv("HOME"), ".kube", "config",
			)
			config, err = clientcmd.BuildConfigFromFlags("", kubeConfigFilepath)
			if err != nil {
				panic(err.Error())
			}
		}

		clientset, _ := kubernetes.NewForConfig(config)
		err = clientset.CoreV1().Namespaces().Delete(context.TODO(), "demo-ns", v1.DeleteOptions{})
		if err != nil {
			fmt.Println("Error: ", err)
			return
		}

		fmt.Println("Namespace delete successful")
	},
}

func init() {
	namespaceCmd.AddCommand(deleteCmd)
}
