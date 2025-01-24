package relativesqueries

import (
	"context"
	"fmt"
	"time"

	"github.com/PhuPhuoc/curanest-patient-service/common"
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

type MyProfileDTO struct {
	MyAccountDTO
	MyRelativesInfoDTO
}

type MyAccountDTO struct {
	Id          uuid.UUID `json:"id"`
	Role        string    `json:"role"`
	FullName    string    `json:"full-name"`
	Email       string    `json:"email"`
	PhoneNumber string    `json:"phone-number"`
	Avatar      string    `json:"avatar"`
	CreatedAt   time.Time `json:"created-at"`
}

type MyRelativesInfoDTO struct {
	Dob      string `json:"dob"`
	Address  string `json:"address"`
	Ward     string `json:"ward"`
	District string `json:"district"`
	City     string `json:"city"`
}

func (h *getMyProfileHandler) Handle(ctx context.Context) (*MyProfileDTO, error) {
	accdto, err := h.accRPC.GetAccountProfile(ctx)
	if err != nil {
		return nil, common.NewInternalServerError().
			WithReason("cannot get profile account info of relatives").
			WithInner(err.Error())
	}

	fmt.Println("accdto: ", accdto)
	return nil, nil
}
