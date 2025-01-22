package main

import (
	"log"

	"github.com/PhuPhuoc/curanest-patient-service/api"
	"github.com/PhuPhuoc/curanest-patient-service/config"
	"github.com/PhuPhuoc/curanest-patient-service/db/mysql"
)

// @title						Patient Service
// @version					1.0
// @description				Auth-service: https://api.curanest.com.vn/auth/swagger/index.html.
// @description				Patient-service: https://api.curanest.com.vn/auth/swagger/index.html.
// @description				Nurse-service: https://api.curanest.com.vn/auth/swagger/index.html.
// @description				Appointment-service (not ready - expected): https://api.curanest.com.vn/auth/swagger/index.html.
// @description				Notification-service (not ready - expected): https://api.curanest.com.vn/auth/swagger/index.html.
// @securityDefinitions.apikey	ApiKeyAuth
// @in							header
// @name						Authorization
func main() {
	config.LoadConfig()
	db := mysql.ConnectDB()
	if err_ping := db.Ping(); err_ping != nil {
		log.Println("Cannot ping db: ", err_ping)
	}
	defer db.Close()

	server := api.InitServer(config.AppConfig.AppPort, db)
	if err_run_server := server.RunApp(); err_run_server != nil {
		log.Fatal("Cannot run app: ", err_run_server)
	}
}
