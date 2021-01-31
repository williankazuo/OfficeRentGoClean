package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"officerent/config"
	"officerent/infra/repository"
	"officerent/usecase/office"

	"officerent/api/handlers"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// Config database
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?parseTime=true", config.DB_USER, config.DB_PASSWORD, config.DB_HOST, config.DB_DATABASE)
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer db.Close()

	// Gin route, and healthcheck route
	r := gin.Default()
	r.GET("/ping", pingHandler)

	officeRepo := repository.NewOfficeMySQL(db)
	officeService := office.NewService(officeRepo)
	handlers.InitOfficeHandlers(r, officeService)

	r.Run()
}

func pingHandler(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}
