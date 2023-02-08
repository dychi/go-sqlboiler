package main

import (
	"github.com/dychi/go-sqlboiler/config"
	"github.com/dychi/go-sqlboiler/handler"
	"github.com/dychi/go-sqlboiler/middleware"
	"github.com/gin-gonic/gin"
)

func main() {
	// config
	_ = config.NewDB()

	// router
	e := gin.New()
	// e.Use(middleware.CustomLogFormatter()) // Custom Formatter
	e.Use(middleware.JSONLogFormatter()) // Custom JSON Formatter
	e.Use(gin.Recovery())                // Recovery
	router := handler.SetApiRoutes(e)
	if err := router.Run(); err != nil {
		panic(err)
	}
}
