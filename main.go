package main

import (
	"github.com/Alfeenn/article/app"
	"github.com/Alfeenn/article/cmd"
	"github.com/Alfeenn/article/controller"
	"github.com/Alfeenn/article/middleware"
	"github.com/Alfeenn/article/repository"
	"github.com/Alfeenn/article/service"
	"github.com/gin-gonic/gin"
)

func main() {
	migrate := cmd.MigrateCmd()
	if migrate {
		return
	}

	engine := gin.New()
	db := app.NewDB()
	repo := repository.NewRepository()
	service := service.NewService(repo, db)
	controller := controller.NewController(service)
	middleware := middleware.NewMiddleware()
	engine.Use(middleware)
	engine.GET("/api/user", controller.FindAll)
	engine.GET("/api/user/:id", controller.Find)
	engine.PUT("/api/user/:id", controller.Update)
	engine.POST("/api/user", controller.Create)
	engine.POST("/api/user/:id", controller.Delete)
	engine.Run("localhost:8000")
}
