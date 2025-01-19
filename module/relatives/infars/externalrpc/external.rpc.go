package relativesexternalrpc

import (
	"context"

	"github.com/PhuPhuoc/curanest-patient-service/common"
	relativescommands "github.com/PhuPhuoc/curanest-patient-service/module/relatives/usecase/commands"
	"github.com/google/uuid"
)

type externalAccountService struct {
	apiURL string
}

func NewAccountService(apiURL string) relativescommands.ExternalAccountService {
	return &externalAccountService{apiURL: apiURL}
}

func (ex *externalAccountService) Create(ctx context.Context, entity *relativescommands.AccountInfoDTO) (uuid.UUID, error) {
	accid := common.GenUUID()
	return accid, nil
}
