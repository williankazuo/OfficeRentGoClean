package handlers

import (
	"encoding/json"
	"net/http"
	"officerent/api/models"
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

func createOffice(service office.UseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		var oReq models.OfficeCreate
		err := json.NewDecoder(c.Request.Body).Decode(&oReq)
		if err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		id, err := service.CreateOffice(oReq.Title, oReq.Description, oReq.People, oReq.Price)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}

		c.JSON(http.StatusOK, id)
	}
}

// InitOfficeHandlers Init routes for office
func InitOfficeHandlers(r *gin.Engine, service office.UseCase) {
	r.GET("/office", getOffice(service))
	r.POST("/office", createOffice(service))
}
