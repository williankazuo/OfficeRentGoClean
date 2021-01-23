package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func getOffice(teste string) gin.HandlerFunc {
	return func(c *gin.Context) {
		teste := struct {
			Teste string
		}{
			"Teste",
		}

		c.JSON(http.StatusOK, teste)
	}
}

func InitOfficeHandlers(r *gin.Engine) {
	r.GET("/office", getOffice("OI"))
}
