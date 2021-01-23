package main

import (
	"net/http"

	"officerent/api/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", pingHandler)

	handlers.InitOfficeHandlers(r)

	r.Run()
}

func pingHandler(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}
