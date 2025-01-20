package relativesexternalrpc

import (
	"context"
	"fmt"

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

func (ex *externalAccountService) Create(ctx context.Context, entity *relativescommands.AccountInfoDTO) (*uuid.UUID, error) {
	response, err := common.CallExternalAPI(ctx, common.RequestOptions{
		Method:  "POST",
		URL:     ex.apiURL + "/external/rpc/accounts",
		Payload: entity,
	})
	if err != nil {
		return nil, fmt.Errorf("cannot call external api - %v", err)
	}
	success, ok := response["success"]
	if ok {
		flagSuccess, _ := success.(bool)
		fmt.Println("flagSuccess: ", flagSuccess)
		if !flagSuccess {
			responseErr, _ := response["error"].(string)
			fmt.Println("responseErr: ", responseErr)
			return nil, fmt.Errorf("%v", responseErr)
		}
	}
	accid, ok := response["data"]
	if ok {
		accuuid := uuid.MustParse(accid.(string))
		return &accuuid, nil
	} else {
		return nil, fmt.Errorf("field data not found")
	}
}
