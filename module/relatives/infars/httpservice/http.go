package relativeshttpservice

import (
	"github.com/PhuPhuoc/curanest-patient-service/middleware"
	relativescommands "github.com/PhuPhuoc/curanest-patient-service/module/relatives/usecase/commands"
	relativesqueries "github.com/PhuPhuoc/curanest-patient-service/module/relatives/usecase/queries"
	"github.com/gin-gonic/gin"
)

type relativesHttpService struct {
	cmd   relativescommands.Commands
	query relativesqueries.Queries
	auth  middleware.AuthClient
}

func NewRelativesHTTPService(cmd relativescommands.Commands, query relativesqueries.Queries) *relativesHttpService {
	return &relativesHttpService{
		cmd:   cmd,
		query: query,
	}
}

func (s *relativesHttpService) AddAuth(auth middleware.AuthClient) *relativesHttpService {
	s.auth = auth
	return s
}

func (s *relativesHttpService) Routes(g *gin.RouterGroup) {
	relatives_route := g.Group("/relatives")
	{
		relatives_route.POST("", s.handleCreateRelativesAccount())
		relatives_route.GET(
			"/me",
			middleware.RequireAuth(s.auth),
			s.handleGetProfile(),
		)
	}
}
