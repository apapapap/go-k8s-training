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
	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

// getCmd represents the get pod command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "kube-client pods get",
	Long:  "Get a pod",
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
		pod, err := GetPod(clientset)
		if err != nil {
			fmt.Println("Error: ", err)
			return
		}
		fmt.Println("Pod name: ", pod.Name)
	},
}

func init() {
	podsCmd.AddCommand(getCmd)
}

func GetPod(clientset *kubernetes.Clientset) (*corev1.Pod, error) {
	namespace := "default"
	pod, err := clientset.CoreV1().Pods(namespace).Get(context.TODO(), "my-pod", v1.GetOptions{})
	if err != nil {
		return nil, err
	}
	return pod, nil
}
