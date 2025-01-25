package relativesqueries

import (
	"context"

	relativesdomain "github.com/PhuPhuoc/curanest-patient-service/module/relatives/domain"
	"github.com/google/uuid"
)

type Queries struct {
	GetMyProfile     *getMyProfileHandler
	GetAccWithFilter *getAccountWithFilterHandler
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
		GetAccWithFilter: NewGetAccountWithFilterHandler(
			b.BuildExternalAccountServiceInQuery(),
		),
	}
}

type RelativesQueryRepo interface {
	FindById(ctx context.Context, id uuid.UUID) (*relativesdomain.Relatives, error)
}

type ExternalAccountService interface {
	GetAccountProfileRPC(ctx context.Context) (*ResponseAccountDTO, error)
	GetAccountWithFilterRPC(ctx context.Context, filter *FilterAccountQuery) ([]AccountDTO, error)
}
