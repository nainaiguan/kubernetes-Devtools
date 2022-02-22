package Client

import (
	"flag"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func NewClientSet() *kubernetes.Clientset {
	kubeconfig := flag.String("kubeconfig", "/root/.kube/config", "")
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err.Error())
	}
	clientset, err:= kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	return clientset
}