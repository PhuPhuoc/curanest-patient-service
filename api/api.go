package api

import (
	"log"
	"net/http"

	"github.com/PhuPhuoc/curanest-patient-service/docs"
	"github.com/PhuPhuoc/curanest-patient-service/middleware"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type server struct {
	address string
	db      *sqlx.DB
}

func InitServerTemp(addr string) *server {
	return &server{
		address: addr,
	}
}

func InitServer(addr string, db *sqlx.DB) *server {
	return &server{
		address: addr,
		db:      db,
	}
}

// @BasePath		/api
// @Summary		ping server
// @Description	ping server
// @Tags			ping
// @Accept			json
// @Produce		json
// @Success		200	{object}	map[string]any	"message success"
// @Failure		400	{object}	error			"Bad request error"
// @Router			/ping [get]
func (sv *server) RunApp() error {
	docs.SwaggerInfo.BasePath = "/api"
	gin.SetMode(gin.ReleaseMode)

	router := gin.New()
	router.Use(
		middleware.SkipSwaggerLog(),
		gin.Recovery(),
	)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	api := router.Group("/api")

	/* ping - test */
	api.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "curanest-patient-service - pong x777"})
	})
	/*
		! start server here
	*/
	log.Println("server start listening at port: ", sv.address)
	return router.Run(sv.address)
}
