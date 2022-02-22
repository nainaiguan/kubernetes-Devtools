package Path

import (
	"SRE/Kubexec"
	"github.com/gin-gonic/gin"
	"k8s.io/client-go/kubernetes"
)

func GetPodList(engine *gin.Engine, clientset *kubernetes.Clientset) {
	engine.GET("/getpodlist/:namespace", func(ctx *gin.Context) {
		namespace := ctx.Param("namespace")
		go Kubexec.GetPodList(clientset, namespace)
	})
}

func NewPod(engine *gin.Engine, clientset *kubernetes.Clientset) {
	engine.POST("/newpod/:namespace/:abspath", func(ctx *gin.Context) {
		namespace := ctx.Param("namespace")
		abspath := ctx.Param("abspath")
		go Kubexec.NewPod(clientset, abspath, namespace)
	})
}

func DeletePod(engine *gin.Engine, clientset *kubernetes.Clientset) {
	engine.DELETE("/deletepod/:namespace/:pdname", func(ctx *gin.Context) {
		namespace := ctx.Param("namespace")
		pdname := ctx.Param("pdname")
		go Kubexec.DeletePod(clientset, namespace, pdname)
	})
}

func GetDeploymentList(engine *gin.Engine, clientset *kubernetes.Clientset) {
	engine.GET("/getdeploymentlist/:namespace", func(ctx *gin.Context) {
		namespace := ctx.Param("namespace")
		go Kubexec.GetDeploymentList(clientset, namespace)
	})
}

func NewDeployment(engine *gin.Engine, clientset *kubernetes.Clientset) {
	engine.POST("/newdeployment/:namespace/:abspath", func(ctx *gin.Context) {
		namespace := ctx.Param("namespace")
		abspath := ctx.Param("abspath")
		go Kubexec.NewDeployment(clientset, abspath, namespace)
	})
}

func DeleteDeployment(engine *gin.Engine, clientset *kubernetes.Clientset) {
	engine.DELETE("/deletedeployment/:namespace/:dpname", func(ctx *gin.Context) {
		namespace := ctx.Param("namespace")
		dpname := ctx.Param("dpname")
		go Kubexec.DeleteDeployment(clientset, namespace, dpname)
	})
}

func GetSvcList(engine *gin.Engine, clientset *kubernetes.Clientset) {
	engine.GET("/getsvclist/:namespace", func(ctx *gin.Context) {
		namespace := ctx.Param("namespace")
		go Kubexec.GetSvcList(clientset, namespace)
	})
}

func NewSvc(engine *gin.Engine, clientset *kubernetes.Clientset) {
	engine.POST("/newsvc/:namespace/:abspath", func(ctx *gin.Context) {
		namespace := ctx.Param("namespace")
		abspath := ctx.Param("abspath")
		go Kubexec.NewSvc(clientset, abspath, namespace)
	})
}

func DeleteSvc(engine *gin.Engine, clientset *kubernetes.Clientset) {
	engine.DELETE("/deletesvc/:namespace/:svcname", func(ctx *gin.Context) {
		namespace := ctx.Param("namespace")
		svcname := ctx.Param("svcname")
		go Kubexec.DeleteSvc(clientset, namespace, svcname)
	})
}