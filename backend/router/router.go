package router

import (
	"chatapp/middleware"
	"io"
	"net/http"
	"os"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

func test(c *gin.Context) {
	name := c.Param("name")
	c.String(http.StatusOK, "Hello %s", name)
}

func getAllChat(c *gin.Context) {
	middleware.ReadCollection()
	c.String(http.StatusOK, "OKay")
}

func Router() *gin.Engine {
	router := gin.Default()
	f, _ := os.Create("gin.log")
	// gin.DefaultWriter = io.MultiWriter(f)
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.GET("/get/:name", test)
	router.GET("/db", getAllChat)
	router.Use(static.Serve("/", static.LocalFile("../frontend/build", true)))
	return router
}
