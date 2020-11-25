package main

import (
	"fmt"
	"gin-blog/models"
	"gin-blog/pkg/logging"
	"gin-blog/pkg/setting"
	"gin-blog/routers"
	"net/http"
)

func main() {
	setting.Setup()
	models.Setup()
	logging.Setup()

	router := routers.InitRouter()
	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.ServerSetting.HttpPort),
		Handler:        router,
		ReadTimeout:    setting.ServerSetting.ReadTimeout,
		WriteTimeout:   setting.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()

}
