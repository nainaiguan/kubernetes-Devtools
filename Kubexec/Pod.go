package Kubexec

import (
	"context"
	"io/ioutil"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/yaml"
	"k8s.io/client-go/kubernetes"
	"log"
)

type PodList struct {
	Name string
	Namespace string
}

func GetPodList(clientset *kubernetes.Clientset, namespace string) []PodList {
	podlist := make([]PodList, 0)
	pod, err := clientset.CoreV1().Pods(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}
	for _, v := range pod.Items {
		tmp := &PodList{Name: v.Name, Namespace: v.Namespace}
		podlist = append(podlist, *tmp)
	}
	return podlist
}

func NewPod(clientset *kubernetes.Clientset, abspath string, namespace string) bool {
	podset := &v1.Pod{}
	config, err := ioutil.ReadFile(abspath)
	if err != nil {
		panic(err.Error())
	}
	data, err := yaml.ToJSON(config)
	if err != nil {
		panic(err.Error())
	}
	if err := yaml.Unmarshal(data, podset); err != nil {
		panic(err.Error())
	}
	if _, err := clientset.CoreV1().Pods(namespace).Get(context.TODO(), podset.Name, metav1.GetOptions{}); err != nil {
		if !errors.IsNotFound(err) {
			return false
		}
		if _, err := clientset.CoreV1().Pods(namespace).Create(context.TODO(), podset, metav1.CreateOptions{}); err != nil {
			log.Fatal()
			return false
		}
		return true
	}else {
		if _, err := clientset.CoreV1().Pods(namespace).Update(context.TODO(), podset, metav1.UpdateOptions{}); err != nil {
			panic(err.Error())
			return false
		}
		return true
	}
}

func DeletePod(clientset *kubernetes.Clientset, namespace string, pdname string) bool {
	if err := clientset.CoreV1().Pods(namespace).Delete(context.TODO(), pdname, metav1.DeleteOptions{}); err != nil {
		panic(err.Error())
		return false
	}
	return true
}