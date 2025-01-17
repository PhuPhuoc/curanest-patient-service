package api

import (
	"log"
	"net/http"

	"github.com/PhuPhuoc/curanest-patient-service/builder"
	"github.com/PhuPhuoc/curanest-patient-service/config"
	"github.com/PhuPhuoc/curanest-patient-service/docs"
	"github.com/PhuPhuoc/curanest-patient-service/middleware"
	relativeshttpservice "github.com/PhuPhuoc/curanest-patient-service/module/relatives/infars/httpservice"
	relativescommands "github.com/PhuPhuoc/curanest-patient-service/module/relatives/usecase/commands"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type server struct {
	port string
	db   *sqlx.DB
}

func InitServer(port string, db *sqlx.DB) *server {
	return &server{
		port: port,
		db:   db,
	}
}

const (
	env_local = "local"
	env_vps   = "vps"

	urlacc_local = "http://localhost:8001/"
	urlacc_prod  = "http://auth_service:8080/"
)

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
	var urlAccServices string
	envDevlopment := config.AppConfig.EnvDev
	if envDevlopment == env_local {
		docs.SwaggerInfo.BasePath = "/"
		urlAccServices = urlacc_local
	}

	if envDevlopment == env_vps {
		gin.SetMode(gin.ReleaseMode)
		docs.SwaggerInfo.BasePath = "/patient"
		urlAccServices = urlacc_prod
	}

	router := gin.New()
	router.Use(middleware.SkipSwaggerLog(), gin.Recovery())
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	router.GET("/ping", func(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"message": "curanest-patient-service - pong"}) })

	// *** usecase: command vs query
	relatives_cmd_builder := relativescommands.NewRelativesCmdWithBuilder(
		builder.NewRelativesBuilder(sv.db).AddUrlPathAccountService(urlAccServices),
	)

	api := router.Group("/api/v1")
	{
		relativeshttpservice.NewAccountHTTPService(relatives_cmd_builder).Routes(api)
	}

	// rpc := router.Group("/internal/rpc")
	// {
	// }
	log.Println("server start listening at port: ", sv.port)
	return router.Run(sv.port)
}
