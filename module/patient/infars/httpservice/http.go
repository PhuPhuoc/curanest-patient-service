package patienthttpservice

import (
	"github.com/PhuPhuoc/curanest-patient-service/middleware"
	patientcommands "github.com/PhuPhuoc/curanest-patient-service/module/patient/usecase/commands"
	"github.com/gin-gonic/gin"
)

type patientHttpService struct {
	cmd  patientcommands.Commands
	auth middleware.AuthClient
}

func NewPatientHTTPService(cmd patientcommands.Commands) *patientHttpService {
	return &patientHttpService{
		cmd: cmd,
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
	}
}
