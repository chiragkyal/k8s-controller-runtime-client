package main

import (
	"context"
	"log"
	"os"

	myclient "github.com/chiragkyal/k8s-controller-runtime-client/client"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/selection"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func main() {
	config := os.Getenv("KUBECONFIG")
	c, err := myclient.NewK8sRestClient(config)
	if err != nil {
		log.Fatalf("Unable to initialize client with error: %s", err.Error())
	}

	pods := &corev1.PodList{}
	namespace := "testlabel"

	listOps := &client.ListOptions{
		Namespace:     namespace,
		LabelSelector: getLabelSelector(),
	}

	err = c.List(context.Background(), pods, listOps)
	if err != nil {
		log.Fatalf("failed to list pods in namespace %s with err %v\n", namespace, err)
	}

	for _, pod := range pods.Items {
		log.Println(pod.Name, ":", pod.Labels, pod.Namespace)
	}

}

func getLabelSelector() labels.Selector {
	// Ref : https://medium.com/coding-kubernetes/using-k8s-label-selectors-in-go-the-right-way-733cde7e8630

	// Like `kubectl label ... -l bss=mybss
	bssReq, err := labels.NewRequirement("bss", selection.Equals, []string{"mybss"})
	if err != nil {
		log.Fatalf("%v\n", err)
	}
	// Like `kubectl label ... -l !service
	serviceReq, err := labels.NewRequirement("service", selection.DoesNotExist, nil)
	if err != nil {
		log.Fatalf("%v\n", err)
	}

	//-------1
	selector := labels.NewSelector()
	selector = selector.Add(*bssReq)
	selector = selector.Add(*serviceReq)
	//-------2
	// selector = selector.Add(
	// 	*bssReq,
	// 	*serviceReq,
	// )
	//-------3
	// selector := labels.NewSelector().Add(
	// 	*bssReq, *serviceReq,
	// )
	log.Printf("Using `%s` labelSelector", selector.String())
	return selector
}
