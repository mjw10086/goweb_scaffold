package main

import (
	"net/http"

	"example.com/goweb/config"
	"example.com/goweb/handler"
	"example.com/goweb/mlogger"
	"example.com/goweb/models"
	"example.com/goweb/tasks"
)

func main() {
	config.Init()
	mlogger.Init(nil)
	models.Init()

	go tasks.Init()

	server := &http.Server{
		Addr:    ":" + config.Port,
		Handler: handler.New(),
	}

	mlogger.Logger.Info("[server start] listen on ", config.Port)
	if err := server.ListenAndServe(); err != nil {
		mlogger.Logger.Fatal("server start error: ", err)
	}
}
