package relativesqueries

import (
	"context"

	"github.com/PhuPhuoc/curanest-patient-service/common"
)

type getMyProfileHandler struct {
	queryRepo RelativesQueryRepo
	accRPC    ExternalAccountService
}

func NewGetMyRelativesAccountHandler(queryRepo RelativesQueryRepo, accRPC ExternalAccountService) *getMyProfileHandler {
	return &getMyProfileHandler{
		queryRepo: queryRepo,
		accRPC:    accRPC,
	}
}

func (h *getMyProfileHandler) Handle(ctx context.Context) (*ResponseProfileDTO, error) {
	accdto, err := h.accRPC.GetAccountProfileRPC(ctx)
	if err != nil {
		return nil, err
	}

	reldto, err := h.queryRepo.FindById(ctx, accdto.Id)
	if err != nil {
		return nil, common.NewInternalServerError().
			WithReason("cannot relatives info").WithInner(err.Error())
	}

	resp := &ResponseProfileDTO{
		accdto, toDTO(reldto),
	}

	return resp, nil
}
