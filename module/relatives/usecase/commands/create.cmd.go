package relativescommands

import (
	"context"

	"github.com/PhuPhuoc/curanest-patient-service/common"
	relativesdomain "github.com/PhuPhuoc/curanest-patient-service/module/relatives/domain"
)

type CreateRelativeAccountCmdDTO struct {
	FullName    string `json:"full-name"`
	PhoneNumber string `json:"phone-number"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	Dob         string `json:"dob"`
	Address     string `json:"address"`
	Ward        string `json:"ward"`
	District    string `json:"district"`
	City        string `json:"city"`
}

type AccountInfoDTO struct {
	RoleName    string `json:"role-name"`
	FullName    string `json:"full-name"`
	PhoneNumber string `json:"phone-number"`
	Email       string `json:"email"`
	Password    string `json:"password"`
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
	accdto := &AccountInfoDTO{
		RoleName:    "relatives",
		FullName:    dto.FullName,
		PhoneNumber: dto.PhoneNumber,
		Email:       dto.Email,
		Password:    dto.Password,
	}
	accid, err := h.accService.Create(ctx, accdto)
	if err != nil {
		return common.NewInternalServerError().
			WithReason("cannot create account for relatives").
			WithInner(err.Error())
	}

	// 2. create record in table relatives
	if accid == nil {
		return common.NewInternalServerError().
			WithReason("cannot create account for relatives - cannot get account id")
	}
	entity, _ := relativesdomain.NewRelatives(
		*accid,
		dto.Dob, dto.Address, dto.Ward,
		dto.District,
		dto.City,
	)
	if err = h.cmdRepo.Create(ctx, entity); err != nil {
		return common.NewInternalServerError().
			WithReason("cannot create relatives info").
			WithInner(err.Error())
	}

	return nil
}
