package main

import (
	"github.com/oumeniOS/go-gin-blog/pkg/setting"
	"fmt"
	"github.com/oumeniOS/go-gin-blog/routers"
	"github.com/fvbock/endless"
	"log"
	"syscall"
	"github.com/oumeniOS/go-gin-blog/models"
)

func main() {

	setting.Setup()
	models.Setup()

	endless.DefaultReadTimeOut = setting.ServerSetting.ReadTimeout
	endless.DefaultWriteTimeOut = setting.ServerSetting.WriteTimeout
	endless.DefaultMaxHeaderBytes = 1 << 20
	endPoint := fmt.Sprintf(":%d", setting.ServerSetting.HttpPort)

	server := endless.NewServer(endPoint, routers.InitRouters())
	server.BeforeBegin = func(add string) {
		log.Printf("Actual pid is %d", syscall.Getpid())
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Printf("Server err :%v", err)
	}
}
