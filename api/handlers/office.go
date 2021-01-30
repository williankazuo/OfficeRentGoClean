package handlers

import (
	"net/http"
	"officerent/usecase/office"

	"github.com/gin-gonic/gin"
)

func getOffice(service office.UseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		teste := struct {
			Teste string
		}{
			"Teste",
		}

		c.JSON(http.StatusOK, teste)
	}
}

// InitOfficeHandlers Init routes for office
func InitOfficeHandlers(r *gin.Engine, service office.UseCase) {
	r.GET("/office", getOffice(service))
}
