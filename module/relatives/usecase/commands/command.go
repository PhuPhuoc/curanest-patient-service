package relativescommands

import (
	"context"

	relativesdomain "github.com/PhuPhuoc/curanest-patient-service/module/relatives/domain"
	"github.com/google/uuid"
)

type Commands struct {
	CreateRelativesAccount *createRelativesAccountHandler
}

type Builder interface {
	BuildRelativesCmdRepo() RelativeCommandRepo
	BuildExternalAccountService() ExternalAccountService
}

func NewRelativesCmdWithBuilder(b Builder) Commands {
	return Commands{
		CreateRelativesAccount: NewCreateRelativesAccountHandler(
			b.BuildRelativesCmdRepo(),
			b.BuildExternalAccountService(),
		),
	}
}

type RelativeCommandRepo interface {
	Create(ctx context.Context, entity *relativesdomain.Relatives) error
}

type ExternalAccountService interface {
	Create(ctx context.Context, entity *AccountInfoDTO) (*uuid.UUID, error)
}
