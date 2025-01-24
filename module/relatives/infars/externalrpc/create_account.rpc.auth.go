package relativesexternalrpc

import (
	"context"
	"fmt"

	"github.com/PhuPhuoc/curanest-patient-service/common"
	relativescommands "github.com/PhuPhuoc/curanest-patient-service/module/relatives/usecase/commands"
)

func (ex *externalAccountService) CreateAccount(ctx context.Context, entity *relativescommands.AccountInfoDTO) (*relativescommands.ResponseCreateAccountDTO, error) {
	response, err := common.CallExternalAPI(ctx, common.RequestOptions{
		Method:  "POST",
		URL:     ex.apiURL + "/external/rpc/accounts",
		Payload: entity,
	})
	if err != nil {
		return nil, fmt.Errorf("cannot call external api - %v", err)
	}
	success, ok := response["success"].(bool)
	if !ok || !success {
		return nil, common.ExtractErrorFromResponse(response)
	}

	return common.ExtractDataFromResponse[relativescommands.ResponseCreateAccountDTO](response)
}
