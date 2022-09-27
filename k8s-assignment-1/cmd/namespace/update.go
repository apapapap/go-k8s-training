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
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

// updateCmd represents the update namespace command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update a namespace",
	Long:  "kube-client update namespace",
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
		namespace, err := GetNamespace(clientset)
		if err != nil {
			fmt.Println("Error: ", err)
			return
		}

		if namespace.ObjectMeta.Labels == nil {
			namespace.ObjectMeta.Labels = make(map[string]string)
		}
		namespace.ObjectMeta.Labels["type"] = "frontend"
		_, err = clientset.CoreV1().Namespaces().Update(context.TODO(), namespace, metav1.UpdateOptions{})
		if err != nil {
			fmt.Println("Error: ", err)
			return
		}

		fmt.Println("Namespace update successful")
	},
}

func init() {
	namespaceCmd.AddCommand(updateCmd)
}
