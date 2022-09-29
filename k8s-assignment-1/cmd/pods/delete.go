/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package pods

import (
	"context"
	"fmt"

	cmdv1 "github.com/apapapap/k8s-dev-training/assignment-1/kube-client/cmd"
	"github.com/spf13/cobra"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// deleteCmd represents the delete pod command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a pod",
	Long:  "kube-client delete pod",
	Run: func(cmd *cobra.Command, args []string) {
		namespace := "default"
		var err error
		if cmdv1.UseCtrlRuntime {
			pod := &corev1.Pod{
				ObjectMeta: metav1.ObjectMeta{
					Namespace: namespace,
					Name:      "my-pod",
				},
			}
			err = cmdv1.CtrlClient.Delete(context.Background(), pod)
		} else {
			err = cmdv1.ClientSet.CoreV1().Pods(namespace).Delete(context.TODO(), "my-pod", metav1.DeleteOptions{})
		}
		if err != nil {
			fmt.Println("Failed to delete pod. Error: ", err)
			return
		}

		fmt.Println("Pod delete successful")
	},
}

func init() {
	podsCmd.AddCommand(deleteCmd)
}
