package relativesexternalrpc

import (
	"context"

	"github.com/PhuPhuoc/curanest-patient-service/common"
	relativesqueries "github.com/PhuPhuoc/curanest-patient-service/module/relatives/usecase/queries"
)

func (ex *externalAccountService) GetAccountWithFilterRPC(ctx context.Context, filter *relativesqueries.FilterAccountQuery) ([]relativesqueries.AccountDTO, error) {
	token, ok := ctx.Value(common.KeyToken).(string)
	if !ok {
		return nil, common.NewInternalServerError().
			WithReason("cannot get list accounts").WithInner("missing token to fetching data from other service")
	}

	response, err := common.CallExternalAPI(ctx, common.RequestOptions{
		Method:  "POST",
		URL:     ex.apiURL + "/external/rpc/accounts/filter",
		Payload: filter,
		Token:   token,
	})
	if err != nil {
		return nil, common.NewInternalServerError().
			WithReason("cannot get list accounts").WithInner("cannot call external api - " + err.Error())
	}

	success, ok := response["success"].(bool)
	if !ok || !success {
		return nil, common.ExtractErrorFromResponse(response)
	}

	newPaging, _ := common.ExtractDataFromResponse[common.Paging](response, "paging")
	filter.Paging = *newPaging

	return common.ExtractListDataFromResponse[relativesqueries.AccountDTO](response, "data")
}
