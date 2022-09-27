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

// listCmd represents the view command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "kube-client pods list",
	Long:  "List pods",
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
		pods, err := clientset.CoreV1().Pods(namespace).List(context.TODO(), v1.ListOptions{})
		if err != nil {
			fmt.Println("Error: ", err)
			return
		}

		fmt.Printf("Namespace: %s\n", namespace)
		fmt.Printf("Number of pods in the cluster: %d\n", len(pods.Items))

		for i, pod := range pods.Items {
			fmt.Printf("%d. %s\n", i+1, pod.Name)
		}
	},
}

func init() {
	podsCmd.AddCommand(listCmd)
}
