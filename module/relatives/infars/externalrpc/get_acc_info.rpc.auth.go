package relativesexternalrpc

import (
	"context"
	"fmt"

	"github.com/PhuPhuoc/curanest-patient-service/common"
	relativesqueries "github.com/PhuPhuoc/curanest-patient-service/module/relatives/usecase/queries"
)

func (ex *externalAccountService) GetAccountProfile(ctx context.Context) (*relativesqueries.ResponseAccountDTO, error) {
	token, ok := ctx.Value(common.KeyToken).(string)
	if !ok {
		return nil, fmt.Errorf("missing token to fetching data from other service")
	}
	response, err := common.CallExternalAPI(ctx, common.RequestOptions{
		Method: "GET",
		URL:    ex.apiURL + "/external/rpc/accounts/me",
		Token:  token,
	})
	if err != nil {
		resp := common.NewInternalServerError().WithReason("cannot call external api: " + err.Error())
		return nil, resp
	}
	success, ok := response["success"].(bool)
	if !ok || !success {
		return nil, common.ExtractErrorFromResponse(response)
	}

	return common.ExtractDataFromResponse[relativesqueries.ResponseAccountDTO](response)
}
