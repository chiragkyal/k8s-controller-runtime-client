package main

import (
	"context"
	"log"
	"os"

	myclient "github.com/chiragkyal/k8s-controller-runtime-client/client"
	corev1 "k8s.io/api/core/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func main() {
	config := os.Getenv("KUBECONFIG")
	c, err := myclient.NewK8sRestClient(config)
	if err != nil {
		log.Fatalf("Unable to initialize client with error: %s", err.Error())
	}

	pods := &corev1.PodList{}
	err = c.List(context.Background(), pods, client.InNamespace("default"))
	if err != nil {
		log.Fatalf("failed to list pods in namespace default with err %v\n", err)
	}

	for _, pod := range pods.Items {
		log.Println(pod.Name)
	}

}
