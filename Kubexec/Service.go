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

type SvcList struct {
	Name string
	Namespace string
	Label map[string]string
}

func GetSvcList(clientset *kubernetes.Clientset, namespace string) []SvcList {
	svclist := make([]SvcList, 0)
	svc, err := clientset.CoreV1().Services(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}
	for _, v := range svc.Items {
		tmp := &SvcList{Name: v.Name, Namespace: v.Namespace, Label: v.Labels}
		svclist = append(svclist, *tmp)
	}
	return svclist
}

func NewSvc(clientset *kubernetes.Clientset, abspath string, namespace string) bool {
	svcset := &v1.Service{}
	config, err := ioutil.ReadFile(abspath)
	if err != nil {
		panic(err.Error())
	}
	data, err := yaml.ToJSON(config)
	if err != nil {
		panic(err.Error())
	}
	if err := yaml.Unmarshal(data, svcset); err != nil {
		panic(err.Error())
	}
	if _, err := clientset.CoreV1().Services(namespace).Get(context.TODO(), svcset.Name, metav1.GetOptions{}); err != nil {
		if !errors.IsNotFound(err) {
			return false
		}
		if _, err := clientset.CoreV1().Services(namespace).Create(context.TODO(), svcset, metav1.CreateOptions{}); err != nil {
			log.Fatal()
			return false
		}
		return true
	}else {
		if _, err := clientset.CoreV1().Services(namespace).Update(context.TODO(), svcset, metav1.UpdateOptions{}); err != nil {
			panic(err.Error())
			return false
		}
		return true
	}
}

func DeleteSvc(clientset *kubernetes.Clientset, namespace string, svcname string) bool {
	if err := clientset.CoreV1().Services(namespace).Delete(context.TODO(), svcname, metav1.DeleteOptions{}); err != nil {
		panic(err.Error())
		return false
	}
	return true
}