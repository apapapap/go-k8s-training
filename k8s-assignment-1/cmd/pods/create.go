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

// createCmd represents the create pod command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create pod",
	Long:  "kube-client create pod",
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

		pod := getPodObj(namespace)
		_, err = clientset.CoreV1().Pods(namespace).Create(context.TODO(), pod, v1.CreateOptions{})
		if err != nil {
			fmt.Println("Error: ", err)
			return
		}
		fmt.Println("Pod created successfully")
	},
}

func getPodObj(namespace string) *corev1.Pod {
	return &corev1.Pod{
		TypeMeta: v1.TypeMeta{},
		ObjectMeta: v1.ObjectMeta{
			Name:      "my-pod",
			Namespace: namespace,
		},
		Spec: corev1.PodSpec{
			Containers: []corev1.Container{
				{
					Name:  "nginx-ap",
					Image: "nginx",
				},
			},
		},
	}
}

func init() {
	podsCmd.AddCommand(createCmd)
}
