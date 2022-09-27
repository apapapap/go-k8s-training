/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package deployments

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

// listCmd represents the list deployments command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "kube-client deployments list",
	Long:  "List deployments",
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
		deployments, err := clientset.AppsV1().Deployments(namespace).List(context.TODO(), v1.ListOptions{})
		if err != nil {
			fmt.Println("Error: ", err)
			return
		}

		fmt.Printf("Namespace: %s\n", namespace)
		fmt.Printf("Number of deployments in the cluster: %d\n", len(deployments.Items))

		for i, deployment := range deployments.Items {
			fmt.Printf("%d. %s\n", i+1, deployment.Name)
		}
	},
}

func init() {
	deploymentsCmd.AddCommand(listCmd)
}
