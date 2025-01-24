package relativesqueries

import (
	"context"
	"time"

	"github.com/PhuPhuoc/curanest-patient-service/common"
	relativesdomain "github.com/PhuPhuoc/curanest-patient-service/module/relatives/domain"
	"github.com/google/uuid"
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

type ResponseProfileDTO struct {
	*ResponseAccountDTO
	*RelativesInfoDTO
}

type ResponseAccountDTO struct {
	Id          uuid.UUID `json:"id"`
	Role        string    `json:"role"`
	FullName    string    `json:"full-name"`
	Email       string    `json:"email"`
	PhoneNumber string    `json:"phone-number"`
	Avatar      string    `json:"avatar"`
	CreatedAt   time.Time `json:"created-at"`
}

type RelativesInfoDTO struct {
	Dob      string `json:"dob"`
	Address  string `json:"address"`
	Ward     string `json:"ward"`
	District string `json:"district"`
	City     string `json:"city"`
}

func toDTO(data *relativesdomain.Relatives) *RelativesInfoDTO {
	dto := &RelativesInfoDTO{
		Dob:      data.GetDOB(),
		Address:  data.GetAddress(),
		Ward:     data.GetWard(),
		District: data.GetDistrict(),
		City:     data.GetCity(),
	}
	return dto
}

func (h *getMyProfileHandler) Handle(ctx context.Context) (*ResponseProfileDTO, error) {
	accdto, err := h.accRPC.GetAccountProfile(ctx)
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
