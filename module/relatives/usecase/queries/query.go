package relativesqueries

import (
	"context"

	relativesdomain "github.com/PhuPhuoc/curanest-patient-service/module/relatives/domain"
	"github.com/google/uuid"
)

type Queries struct {
	GetMyProfile *getMyProfileHandler
}

type Builder interface {
	BuildRelativesQueryRepo() RelativesQueryRepo
	BuildExternalAccountServiceInQuery() ExternalAccountService
}

func NewRelativesQueryWithBuilder(b Builder) Queries {
	return Queries{
		GetMyProfile: NewGetMyRelativesAccountHandler(
			b.BuildRelativesQueryRepo(),
			b.BuildExternalAccountServiceInQuery(),
		),
	}
}

type RelativesQueryRepo interface {
	FindById(ctx context.Context, id uuid.UUID) (*relativesdomain.Relatives, error)
}

type ExternalAccountService interface {
	GetAccountProfile(ctx context.Context) (*MyAccountDTO, error)
}
