package relativesexternalrpc

import (
	"context"

	"github.com/google/uuid"

	"github.com/PhuPhuoc/curanest-patient-service/common"
	relativescommands "github.com/PhuPhuoc/curanest-patient-service/module/relatives/usecase/commands"
)

func (ex *externalAccountService) UpdateAccountRPC(ctx context.Context, relativesId *uuid.UUID, entity *relativescommands.UpdateAccountInfoDTO) error {
	token, ok := ctx.Value(common.KeyToken).(string)
	if !ok {
		return common.NewInternalServerError().
			WithReason("cannot update accounts").WithInner("missing token to fetching data from other service")
	}

	response, err := common.CallExternalAPI(ctx, common.RequestOptions{
		Method:  "PUT",
		URL:     ex.apiURL + "/external/rpc/accounts/" + relativesId.String(),
		Token:   token,
		Payload: entity,
	})
	if err != nil {
		resp := common.NewInternalServerError().WithReason("cannot call external api: " + err.Error())
		return resp
	}

	success, ok := response["success"].(bool)
	if !ok || !success {
		return common.ExtractErrorFromResponse(response)
	}
	return nil
}
