package relativescommands

import (
	"context"
	"fmt"

	relativesdomain "github.com/PhuPhuoc/curanest-patient-service/module/relatives/domain"
)

type CreateRelativeAccountCmdDTO struct {
	AccountInfoDTO
	RelativesInfoDTO
}

type AccountInfoDTO struct {
	RoleName    string `json:"role-name"`
	FullName    string `json:"full-name"`
	PhoneNumber string `json:"phone-number"`
	Email       string `json:"email"`
	Password    string `json:"password"`
}

type RelativesInfoDTO struct {
	Dob      string `json:"dob"`
	Address  string `json:"address"`
	Ward     string `json:"ward"`
	District string `json:"district"`
	City     string `json:"city"`
}

type createRelativesAccountHandler struct {
	cmdRepo    RelativeCommandRepo
	accService ExternalAccountService
}

func NewCreateRelativesAccountHandler(cmdRepo RelativeCommandRepo, accService ExternalAccountService) *createRelativesAccountHandler {
	return &createRelativesAccountHandler{
		cmdRepo:    cmdRepo,
		accService: accService,
	}
}

func (h *createRelativesAccountHandler) Handle(ctx context.Context, dto *CreateRelativeAccountCmdDTO) error {
	// 1. call external service
	dto.RoleName = "relatives"
	accid, err := h.accService.Create(ctx, &dto.AccountInfoDTO)
	if err != nil {
		return err
	}

	// 2. create record in table relatives
	if accid == nil {
		return fmt.Errorf("failed to create account: received nil account ID")
	}
	entity, _ := relativesdomain.NewRelatives(
		*accid,
		dto.Dob, dto.Address, dto.Ward,
		dto.District,
		dto.City,
	)
	if err = h.cmdRepo.Create(ctx, entity); err != nil {
		return err
	}

	return nil
}
