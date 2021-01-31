package handlers

import (
	"net/http"
	"officerent/usecase/office"

	"github.com/gin-gonic/gin"
)

func getOffice(service office.UseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		offices, err := service.GetAllOffices()
		if err != nil {
			c.JSON()
		}

		c.JSON(http.StatusOK, offices)
	}
}

// InitOfficeHandlers Init routes for office
func InitOfficeHandlers(r *gin.Engine, service office.UseCase) {
	r.GET("/office", getOffice(service))
}
