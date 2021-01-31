package handlers

import (
	"net/http"
	"officerent/entity"
	"officerent/usecase/office"

	"github.com/gin-gonic/gin"
)

func getOffice(service office.UseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		offices, err := service.GetAllOffices()
		if err != nil && err != entity.ErrNotFound {
			c.JSON(http.StatusInternalServerError, err)
			return
		}
		if err == entity.ErrNotFound {
			c.JSON(http.StatusNotFound, err)
			return
		}

		c.JSON(http.StatusOK, offices)
		return
	}
}

// InitOfficeHandlers Init routes for office
func InitOfficeHandlers(r *gin.Engine, service office.UseCase) {
	r.GET("/office", getOffice(service))
}
