/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"fmt"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

var (
	kubeConfigPath string
	namespace      string
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List deployments in default namespace",
	Run: func(cmd *cobra.Command, args []string) {
		log.Debug().Msg("Running list command")
		kubeConfigPath = viper.GetString("config")
		client, err := getKubeClient(kubeConfigPath)
		if err != nil {
			log.Error().Msg("Failed to create Kubernetes client")
			return
		}

		deployments, err := client.AppsV1().Deployments(namespace).List(context.Background(), metav1.ListOptions{})
		if err != nil {
			fmt.Printf("Error listing deployments: %v\n", err)
			return
		}
		fmt.Printf("Deployments in %s namespace:\n", namespace)
		for _, item := range deployments.Items {
			fmt.Println(item.Name)
		}
		sf, err := client.AppsV1().StatefulSets(namespace).List(context.Background(), metav1.ListOptions{})
		if err != nil {
			fmt.Printf("Error listing StatefulSets: %v\n", err)
			return
		}
		fmt.Printf("StatefulSets in %s namespace:\n", namespace)
		for _, item := range sf.Items {
			fmt.Println(item.Name)
		}
	},
}

func getKubeClient(kubeConfigPath string) (*kubernetes.Clientset, error) {
	config, err := clientcmd.BuildConfigFromFlags("", kubeConfigPath)
	if err != nil {
		log.Error().Err(err).Msg("Failed to build Kubernetes config")
		return nil, err
	}
	return kubernetes.NewForConfig(config)
}

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.Flags().StringVarP(&kubeConfigPath, "config", "c", "", "Path to the k8s config file")
	listCmd.Flags().StringVarP(&namespace, "namespace", "n", "default", "namespace to list resources from")
	viper.BindPFlag("config", listCmd.Flags().Lookup("config"))
	viper.BindEnv("config", "KUBECONFIG")
	viper.SetDefault("config", "~/.kube/config")

}
