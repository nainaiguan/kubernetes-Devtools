package main

import (
	"SRE/Client"
	"SRE/Path"
	"github.com/gin-gonic/gin"
)

func main() {
	client := Client.NewClientSet()
	engine := gin.Default()
	Path.GetPodList(engine, client)
	Path.NewPod(engine, client)
	Path.DeletePod(engine, client)
	Path.GetDeploymentList(engine, client)
	Path.NewDeployment(engine, client)
	Path.DeleteDeployment(engine, client)
	Path.GetSvcList(engine, client)
	Path.NewSvc(engine, client)
	Path.DeleteSvc(engine, client)
	if err := engine.Run(":8080"); err != nil {
		panic(err.Error())
	}
}
