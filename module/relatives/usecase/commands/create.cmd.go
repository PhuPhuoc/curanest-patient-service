package relativescommands

import (
	"context"
	"fmt"

	"github.com/PhuPhuoc/curanest-patient-service/common"
	relativesdomain "github.com/PhuPhuoc/curanest-patient-service/module/relatives/domain"
	"github.com/google/uuid"
)

type CreateRelativeAccountCmdDTO struct {
	FullName    string `json:"full-name"`
	Gender      bool   `json:"gender"`
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

type ResponseCreateAccountDTO struct {
	Id string `json:"id"`
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

	resp, err := h.accService.CreateAccountRPC(ctx, accdto)
	if err != nil {
		fmt.Println("error: ", err)
		return err
	}

	if resp.Id == "" {
		return common.NewInternalServerError().
			WithReason("cannot create account for relatives - cannot get account id")
	}

	// 2. create record in table relatives
	accid := uuid.MustParse(resp.Id)
	entity, _ := relativesdomain.NewRelatives(
		accid,
		dto.Gender,
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
