/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package pods

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

// deleteCmd represents the delete pod command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a pod",
	Long:  "kube-client delete pod",
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
		namespace := "default"

		err = clientset.CoreV1().Pods(namespace).Delete(context.TODO(), "my-pod", v1.DeleteOptions{})
		if err != nil {
			fmt.Println("Error: ", err)
			return
		}

		fmt.Println("Pod delete successful")
	},
}

func init() {
	podsCmd.AddCommand(deleteCmd)
}
