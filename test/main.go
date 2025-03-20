package main

import (
	"context"
	"fmt"
	"log"
	"path/filepath"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	metricsv "k8s.io/metrics/pkg/apis/metrics/v1beta1"
)

func main() {
	// Load kubeconfig from the default location or set it manually
	kubeconfig := filepath.Join(homedir.HomeDir(), ".kube", "config")
	clientConfig, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		log.Fatalf("failed to build kubeconfig: %v", err)
	}

	// Create a clientset to interact with the Kubernetes cluster
	clientset, err := kubernetes.NewForConfig(clientConfig)
	if err != nil {
		log.Fatalf("failed to create clientset: %v", err)
	}

	// Access the metrics client using the MetricsV1beta1 API
	metricsv.Build

	// Get the metrics for the pods in the "default" namespace
	podMetricsList, err := metricsClient.PodMetricses("default").List(context.Background(), metav1.ListOptions{})
	if err != nil {
		log.Fatalf("failed to get pod metrics: %v", err)
	}

	// Iterate over pods and containers to print only CPU usage
	for _, podMetrics := range podMetricsList.Items {
		fmt.Printf("Pod: %s\n", podMetrics.Name)
		for _, container := range podMetrics.Containers {
			// Only print CPU usage (no memory or other resources)
			cpuUsage, found := container.Usage["cpu"]
			if found {
				fmt.Printf("\tContainer: %s\n", container.Name)
				fmt.Printf("\t\tCPU Usage: %s\n", cpuUsage.String())
			}
		}
	}
}
