package relativescommands

import (
	"context"

	"github.com/PhuPhuoc/curanest-patient-service/common"
	relativesdomain "github.com/PhuPhuoc/curanest-patient-service/module/relatives/domain"
	"github.com/google/uuid"
)

type updateRelativesAccountHandler struct {
	cmdRepo    RelativeCommandRepo
	accService ExternalAccountService
}

func NewUpdateRelativesAccountHandler(cmdRepo RelativeCommandRepo, accService ExternalAccountService) *updateRelativesAccountHandler {
	return &updateRelativesAccountHandler{
		cmdRepo:    cmdRepo,
		accService: accService,
	}
}

func (h *updateRelativesAccountHandler) Handle(ctx context.Context, id *uuid.UUID, dto *UpdateRelativeAccountCmdDTO) error {
	if err := h.accService.UpdateAccountRPC(ctx, id, &dto.UpdateAccountInfoDTO); err != nil {
		return err
	}

	entity, _ := relativesdomain.NewRelatives(
		*id,
		dto.Gender,
		dto.Dob,
		dto.Address,
		dto.Ward,
		dto.District,
		dto.City,
	)

	if err := h.cmdRepo.Update(ctx, entity); err != nil {
		return common.NewInternalServerError().
			WithReason("cannot update relatives info").
			WithInner(err.Error())
	}

	return nil
}
