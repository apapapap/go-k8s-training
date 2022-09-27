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
	appsv1 "k8s.io/api/apps/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

// getCmd represents the get deployment command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "kube-client deployment get",
	Long:  "Get a deployment",
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
		deployment, err := GetDeployment(clientset)
		if err != nil {
			fmt.Println("Error: ", err)
			return
		}
		fmt.Println("Deployment name: ", deployment.Name)
	},
}

func init() {
	deploymentsCmd.AddCommand(getCmd)
}

func GetDeployment(clientset *kubernetes.Clientset) (*appsv1.Deployment, error) {
	namespace := "default"
	deployment, err := clientset.AppsV1().Deployments(namespace).Get(context.TODO(), "demo-deployment", v1.GetOptions{})
	if err != nil {
		return nil, err
	}
	return deployment, nil
}
