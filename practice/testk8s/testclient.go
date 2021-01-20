package main

import (
	"context"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"

	"github.com/sirupsen/logrus"
)

func main() {

	logrus.Info("---开始调用kubernetes---")
	config, err := clientcmd.BuildConfigFromFlags("", "D:/kubeconfig/config")

	if err != nil {
		panic(err)
	}

	config.APIPath = "api"
	config.GroupVersion = &corev1.SchemeGroupVersion
	config.NegotiatedSerializer = scheme.Codecs

	restClient, err := rest.RESTClientFor(config)

	if err != nil {
		panic(err)
	}

	result := &corev1.PodList{}

	err = restClient.Get().Namespace("kube-system").Resource("pods").
		VersionedParams(&metav1.ListOptions{Limit: 500}, scheme.ParameterCodec).
		Do(context.TODO()).Into(result)
	if err != nil {
		panic(err)
	}

	for _, d := range result.Items {
		logrus.Infof("Namespace {%s} Name: {%s}", d.Namespace, d.Name)
	}

	logrus.Info("---调用结束---")

}
