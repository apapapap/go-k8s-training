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
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

// updateCmd represents the update deployment command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update a deployment",
	Long:  "kube-client update deployment",
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
		deployment, err := GetDeployment(clientset)
		if err != nil {
			fmt.Println("Error: ", err)
			return
		}

		if deployment.ObjectMeta.Labels == nil {
			deployment.ObjectMeta.Labels = make(map[string]string)
		}
		deployment.ObjectMeta.Labels["type1"] = "frontend1"
		_, err = clientset.AppsV1().Deployments(namespace).Update(context.TODO(), deployment, metav1.UpdateOptions{})
		if err != nil {
			fmt.Println("Error: ", err)
			return
		}

		fmt.Println("Deployment update successful")
	},
}

func init() {
	deploymentsCmd.AddCommand(updateCmd)
}
