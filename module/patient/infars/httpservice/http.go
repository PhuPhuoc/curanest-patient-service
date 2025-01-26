package patienthttpservice

import (
	"github.com/PhuPhuoc/curanest-patient-service/middleware"
	patientcommands "github.com/PhuPhuoc/curanest-patient-service/module/patient/usecase/commands"
	patientqueries "github.com/PhuPhuoc/curanest-patient-service/module/patient/usecase/queries"
	"github.com/gin-gonic/gin"
)

type patientHttpService struct {
	cmd   patientcommands.Commands
	query patientqueries.Queries
	auth  middleware.AuthClient
}

func NewPatientHTTPService(cmd patientcommands.Commands, query patientqueries.Queries) *patientHttpService {
	return &patientHttpService{
		cmd:   cmd,
		query: query,
	}
}

func (s *patientHttpService) AddAuth(auth middleware.AuthClient) *patientHttpService {
	s.auth = auth
	return s
}

func (s *patientHttpService) Routes(g *gin.RouterGroup) {
	patient_route := g.Group("/patients")
	{
		patient_route.POST(
			"",
			middleware.RequireAuth(s.auth),
			middleware.RequireRole("relatives"),
			s.handleCreatePatientProfile(),
		)

		patient_route.PUT(
			":patient-id",
			middleware.RequireAuth(s.auth),
			middleware.RequireRole("relatives"),
			s.handleUpdatePatientProfile(),
		)

		patient_route.GET("/relatives",
			middleware.RequireAuth(s.auth),
			middleware.RequireRole("relatives"),
			s.handleGetMyPatient(),
		)
	}
}
