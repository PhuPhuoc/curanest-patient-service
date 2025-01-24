package relativesexternalrpc

import (
	"context"
	"fmt"

	"github.com/PhuPhuoc/curanest-patient-service/common"
	relativesqueries "github.com/PhuPhuoc/curanest-patient-service/module/relatives/usecase/queries"
)

func (ex *externalAccountService) GetAccountProfile(ctx context.Context) (*relativesqueries.MyAccountDTO, error) {
	token, ok := ctx.Value(common.KeyToken).(string)
	if !ok {
		return nil, fmt.Errorf("missing token to fetching data from other service")
	}
	response, err := common.CallExternalAPI(ctx, common.RequestOptions{
		Method: "GET",
		URL:    ex.apiURL + "/external/rpc/accounts/me",
		Token:  token,
	})
	fmt.Println("URL: " + ex.apiURL + "/external/rpc/accounts/me")
	if err != nil {
		return nil, fmt.Errorf("cannot call external api - %v", err)
	}
	success, ok := response["success"]
	if ok {
		flagSuccess, _ := success.(bool)
		if !flagSuccess {
			responseErr, _ := response["error"].(string)
			return nil, fmt.Errorf("%v", responseErr)
		}
	}

	fmt.Println("response: ", response)
	return nil, nil
}
