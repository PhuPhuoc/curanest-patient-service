package relativescommands

import (
	"context"

	relativesdomain "github.com/PhuPhuoc/curanest-patient-service/module/relatives/domain"
)

type Commands struct {
	CreateRelativesAccount *createRelativesAccountHandler
}

type Builder interface {
	BuildRelativesCmdRepo() RelativeCommandRepo
	BuildExternalAccountServiceInCmd() ExternalAccountService
}

func NewRelativesCmdWithBuilder(b Builder) Commands {
	return Commands{
		CreateRelativesAccount: NewCreateRelativesAccountHandler(
			b.BuildRelativesCmdRepo(),
			b.BuildExternalAccountServiceInCmd(),
		),
	}
}

type RelativeCommandRepo interface {
	Create(ctx context.Context, entity *relativesdomain.Relatives) error
}

type ExternalAccountService interface {
	CreateAccountRPC(ctx context.Context, entity *AccountInfoDTO) (*ResponseCreateAccountDTO, error)
}
