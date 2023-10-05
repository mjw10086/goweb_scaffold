package handler

import (
	"path/filepath"

	"example.com/goweb/config"
	"example.com/goweb/middlewares"
	"example.com/goweb/models"
	"github.com/gin-gonic/gin"
)

func New() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	engine := gin.New()

	engine.LoadHTMLGlob(filepath.Join(config.BaseDir, "views/*"))
	engine.Static("/static", filepath.Join(config.BaseDir, "static/"))

	engine.Use(middlewares.Logger())

	registerRouter(engine)

	return engine
}

func registerRouter(r *gin.Engine) {
	r.GET("/", Index)
}

func Index(c *gin.Context) {
	c.HTML(200, "index.html", gin.H{
		"message": models.Store.GetMessage(),
	})
}
