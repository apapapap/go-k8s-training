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
	apiv1 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

// getCmd represents the get namespace command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "kube-client namespace get",
	Long:  "Get a namespace",
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
		fmt.Println("Namespace name: ", namespace.Name)
	},
}

func init() {
	namespaceCmd.AddCommand(getCmd)
}

func GetNamespace(clientset *kubernetes.Clientset) (*apiv1.Namespace, error) {
	namespace, err := clientset.CoreV1().Namespaces().Get(context.TODO(), "demo-ns", v1.GetOptions{})
	if err != nil {
		return nil, err
	}
	return namespace, nil
}
