package main

import (
	"log"

	"github.com/elprogramadorgt/StreamAlchemy/router"
	"github.com/gin-gonic/gin"
)

func main() {

	gin.ForceConsoleColor()
	engine := gin.Default()
	engine.ContextWithFallback = true
	router.GinRoutes(engine)
	log.Fatal(engine.Run("0.0.0.0:8080"))
}
