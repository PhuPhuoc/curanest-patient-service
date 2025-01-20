package patienthttpservice

import (
	patientcommands "github.com/PhuPhuoc/curanest-patient-service/module/patient/usecase/commands"
	"github.com/gin-gonic/gin"
)

type patientHttpService struct {
	cmd patientcommands.Commands
}

func NewPatientHTTPService(cmd patientcommands.Commands) *patientHttpService {
	return &patientHttpService{
		cmd: cmd,
	}
}

func (s *patientHttpService) Routes(g *gin.RouterGroup) {
	patient_route := g.Group("/patients")
	{
		patient_route.POST("", s.handleCreatePatientProfile())
	}
}
