package relativeshttpservice

import (
	relativescommands "github.com/PhuPhuoc/curanest-patient-service/module/relatives/usecase/commands"
	"github.com/gin-gonic/gin"
)

type relativesHttpService struct {
	cmd relativescommands.Commands
}

func NewRelativesHTTPService(cmd relativescommands.Commands) *relativesHttpService {
	return &relativesHttpService{
		cmd: cmd,
	}
}

func (s *relativesHttpService) Routes(g *gin.RouterGroup) {
	acc_route := g.Group("/relatives")
	{
		acc_route.POST("", s.handleCreateRelativesAccount())
	}
}
