package Kubexec

import (
	"context"
	"io/ioutil"
	v1 "k8s.io/api/apps/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/yaml"
	"k8s.io/client-go/kubernetes"
)

type deploymentList struct {
	Name string
	Namespace string
	Replicas *int32
}

func GetDeploymentList(clientset *kubernetes.Clientset, namespace string) []deploymentList {
	dplist := make([]deploymentList, 0)
	deployment, err := clientset.AppsV1().Deployments(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}
	for _, v := range deployment.Items {
		tmp := &deploymentList{Name: v.Name, Namespace: v.Namespace, Replicas: v.Spec.Replicas}
		dplist = append(dplist, *tmp)
	}
	return dplist
}

func NewDeployment(clientset *kubernetes.Clientset, abspath string, namespace string) bool {
	deploymentset := &v1.Deployment{}
	config, err := ioutil.ReadFile(abspath)
	if err != nil {
		panic(err.Error())
	}
	data, err := yaml.ToJSON(config)
	if err != nil {
		panic(err.Error())
	}
	if err := yaml.Unmarshal(data, deploymentset); err != nil {
		panic(err.Error())
	}
	if _, err := clientset.AppsV1().Deployments(namespace).Get(context.TODO(), deploymentset.Name, metav1.GetOptions{}); err != nil {
		if !errors.IsNotFound(err) {
			panic(err.Error())
		}
		_, err := clientset.AppsV1().Deployments(namespace).Create(context.TODO(), deploymentset, metav1.CreateOptions{})
		if err != nil {
			panic(err.Error())
		}
		return true
	}else {
		_, err := clientset.AppsV1().Deployments(namespace).Update(context.TODO(), deploymentset, metav1.UpdateOptions{})
		if err != nil {
			panic(err.Error())
		}
		return true
	}
}

func DeleteDeployment(clientset *kubernetes.Clientset, namespace string, dpname string) bool {
	if _, err := clientset.AppsV1().Deployments(namespace).Get(context.TODO(), dpname, metav1.GetOptions{}); err != nil {
		if !errors.IsNotFound(err) {
			panic(err.Error())
			return false
		}
		return false
	}else {
		if err := clientset.AppsV1().Deployments(namespace).Delete(context.TODO(), dpname, metav1.DeleteOptions{}); err != nil {
			panic(err.Error())
		}
		return true
	}
}