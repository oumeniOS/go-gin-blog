package main

import (
	"github.com/EDDYCJY/gin-blog/pkg/setting"
	"fmt"
	"github.com/EDDYCJY/gin-blog/routers"
	"github.com/fvbock/endless"
	"log"
	"syscall"
)

func main() {

	endless.DefaultReadTimeOut = setting.ReadTimeOut
	endless.DefaultWriteTimeOut = setting.WriteTimeOut
	endless.DefaultMaxHeaderBytes = 1 << 20
	endPoint := fmt.Sprintf(":%d", setting.HTTPPort)

	server := endless.NewServer(endPoint, routers.InitRouters())
	server.BeforeBegin = func(add string) {
		log.Printf("Actual pid is %d", syscall.Getpid())
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Printf("Server err :%v", err)
	}

	//router := routers.InitRouters()
	//
	//server := &http.Server{
	//	Addr:           fmt.Sprintf(":%d", setting.HTTPPort),
	//	Handler:        router,
	//	ReadTimeout:    setting.ReadTimeOut,
	//	WriteTimeout:   setting.WriteTimeOut,
	//	MaxHeaderBytes: 1 << 20,
	//}
	//server.ListenAndServe()
}
