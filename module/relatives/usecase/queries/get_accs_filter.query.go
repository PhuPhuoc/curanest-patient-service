package relativesqueries

import (
	"context"
)

type getAccountWithFilterHandler struct {
	accRPC ExternalAccountService
}

func NewGetAccountWithFilterHandler(accRPC ExternalAccountService) *getAccountWithFilterHandler {
	return &getAccountWithFilterHandler{
		accRPC: accRPC,
	}
}

func (h *getAccountWithFilterHandler) Handle(ctx context.Context, filter *FilterAccountQuery) ([]AccountDTO, error) {
	filter.Paging.Process()
	accs, err := h.accRPC.GetAccountWithFilterRPC(ctx, filter)
	if err != nil {
		return nil, err
	}

	return accs, nil
}
