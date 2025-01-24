package api

import (
	"log"
	"net/http"
	"time"

	"github.com/PhuPhuoc/curanest-patient-service/builder"
	"github.com/PhuPhuoc/curanest-patient-service/common"
	"github.com/PhuPhuoc/curanest-patient-service/config"
	"github.com/PhuPhuoc/curanest-patient-service/docs"
	"github.com/PhuPhuoc/curanest-patient-service/middleware"
	patienthttpservice "github.com/PhuPhuoc/curanest-patient-service/module/patient/infars/httpservice"
	patientcommands "github.com/PhuPhuoc/curanest-patient-service/module/patient/usecase/commands"
	patientqueries "github.com/PhuPhuoc/curanest-patient-service/module/patient/usecase/queries"
	relativeshttpservice "github.com/PhuPhuoc/curanest-patient-service/module/relatives/infars/httpservice"
	relativescommands "github.com/PhuPhuoc/curanest-patient-service/module/relatives/usecase/commands"
	relativesqueries "github.com/PhuPhuoc/curanest-patient-service/module/relatives/usecase/queries"
	"github.com/gin-contrib/cors"
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

	urlacc_local = "http://localhost:8001"
	urlacc_prod  = "http://auth_service:8080"
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
		// gin.SetMode(gin.ReleaseMode)
		docs.SwaggerInfo.BasePath = "/patient"
		urlAccServices = urlacc_prod
	}

	router := gin.New()

	configcors := cors.DefaultConfig()
	configcors.AllowAllOrigins = true
	configcors.AllowMethods = []string{"POST", "GET", "PUT", "OPTIONS"}
	configcors.AllowHeaders = []string{"Origin", "Content-Type", "Authorization", "Accept", "User-Agent", "Cache-Control", "Pragma"}
	configcors.ExposeHeaders = []string{"Content-Length"}
	configcors.AllowCredentials = true
	configcors.MaxAge = 12 * time.Hour

	router.Use(cors.New(configcors))
	router.Use(middleware.SkipSwaggerLog(), gin.Recovery())
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	router.GET("/ping", func(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"message": "curanest-patient-service - pong"}) })

	authClient := common.NewJWTx(config.AppConfig.Key)
	// *** usecase: command vs query
	relatives_cmd_builder := relativescommands.NewRelativesCmdWithBuilder(
		builder.NewRelativesBuilder(sv.db).AddUrlPathAccountService(urlAccServices),
	)
	relatives_query_builder := relativesqueries.NewRelativesQueryWithBuilder(
		builder.NewRelativesBuilder(sv.db).AddUrlPathAccountService(urlAccServices),
	)

	patient_cmd_builder := patientcommands.NewPatientCmdWithBuilder(
		builder.NewPatientBuilder(sv.db),
	)
	patient_query_builder := patientqueries.NewPatientQueryWithBuilder(
		builder.NewPatientBuilder(sv.db),
	)

	api := router.Group("/api/v1")
	{
		relativeshttpservice.NewRelativesHTTPService(relatives_cmd_builder, relatives_query_builder).AddAuth(authClient).Routes(api)
		patienthttpservice.NewPatientHTTPService(patient_cmd_builder, patient_query_builder).AddAuth(authClient).Routes(api)
	}

	// rpc := router.Group("/internal/rpc")
	// {
	// }
	log.Println("server start listening at port: ", sv.port)
	return router.Run(sv.port)
}
