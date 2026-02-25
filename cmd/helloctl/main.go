package main

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

func newRootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "helloctl",
		Short: "helloctl is a CLI for interacting with Kubernetes",
	}

	cmd.AddCommand(newVersionCmd())
	cmd.AddCommand(newListPodsCmd())
	return cmd
}

func newVersionCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Print the CLI version",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("helloctl dev")
		},
	}
}

func newListPodsCmd() *cobra.Command {
	var namespace string

	cmd := &cobra.Command{
		Use:   "listpods",
		Short: "List pods in a namespace",
		RunE: func(cmd *cobra.Command, args []string) error {
			// Conecta ao cluster Kubernetes
			clientset, err := getClientset()
			if err != nil {
				return fmt.Errorf("failed to create k8s client: %w", err)
			}

			// Lista os pods no namespace especificado
			pods, err := clientset.CoreV1().Pods(namespace).List(context.Background(), metav1.ListOptions{})
			if err != nil {
				return fmt.Errorf("failed to list pods: %w", err)
			}

			// Exibe os pods encontrados
			fmt.Printf("Pods in namespace '%s':\n", namespace)
			if len(pods.Items) == 0 {
				fmt.Println("No pods found.")
				return nil
			}

			for _, pod := range pods.Items {
				fmt.Printf("- %s (%s)\n", pod.Name, pod.Status.Phase)
			}
			
			return nil
		},
	}

	// Adiciona a flag para especificar o namespace
	cmd.Flags().StringVarP(&namespace, "namespace", "n", "default", "Kubernetes namespace to list pods from")
	
	return cmd
}

// Função auxiliar para conectar ao cluster Kubernetes
func getClientset() (*kubernetes.Clientset, error) {
	var kubeconfig string
	
	// Tenta encontrar o kubeconfig no diretório padrão
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = filepath.Join(home, ".kube", "config")
	}

	// Cria a configuração a partir do kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		return nil, fmt.Errorf("failed to build config from flags: %w", err)
	}

	// Cria o clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, fmt.Errorf("failed to create clientset: %w", err)
	}
	
	return clientset, nil
}

func main() {
	root := newRootCmd()
	if err := root.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}