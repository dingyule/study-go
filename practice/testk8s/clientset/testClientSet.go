package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	logrus.Info("---clientset开始调用kubernetes---")
	config, err := clientcmd.BuildConfigFromFlags("", "/root/.kube/config")
	if err != nil {
		panic(err)
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}

	podClient := clientset.CoreV1().Pods(apiv1.NamespaceDefault)

	list, err := podClient.List(metav1.ListOptions{Limit: 500})
	if err != nil {
		panic(err)
	}

	for _, d := range list.Items {
		fmt.Printf("namespace:%v \t name:%v \t status:%v", d.Namespace, d.Name, d.Status.Phase)
	}

	logrus.Info("---clientset调用结束---")

}
