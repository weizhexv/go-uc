package web

import (
	"dghire.com/libs/go-web/dgroute"
	"dghire.com/libs/go-web/middleware"
	"dghire.com/libs/go-web/wrapper"
	"github.com/gin-gonic/gin"
	"go-uc/client"
	"go-uc/internal/store"
	"go-uc/vconfig"
	"go-uc/web/controller"
	"log"
	"strconv"
)

var engine = initEngine()

func initEngine() *gin.Engine {
	e := wrapper.NewEngine(store.DB(), vconfig.ServiceName(), vconfig.ServicePort(),
		middleware.Recover(), middleware.Logger(), middleware.Monitor())

	g := e.Group("/uc")
	controller.Auth.Bind(g)
	controller.Employee.Bind(g)
	controller.Login.Bind(g)
	controller.User.Bind(g)
	client.UcClient.Bind(g)

	err := dgroute.Loader.RefreshCheckpoint()
	if err != nil {
		panic(err)
	}
	return e
}

func Run() {
	log.Fatal(engine.Run("0.0.0.0:" + strconv.Itoa(vconfig.ServicePort())))
}
