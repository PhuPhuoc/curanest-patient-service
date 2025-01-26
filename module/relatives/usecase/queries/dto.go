package relativesqueries

import (
	"time"

	"github.com/PhuPhuoc/curanest-patient-service/common"
	relativesdomain "github.com/PhuPhuoc/curanest-patient-service/module/relatives/domain"
	"github.com/google/uuid"
)

type ResponseProfileDTO struct {
	*ResponseAccountDTO
	*RelativesInfoDTO
}

// get_account_info
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
	Gender   bool   `json:"gender"`
	Dob      string `json:"dob"`
	Address  string `json:"address"`
	Ward     string `json:"ward"`
	District string `json:"district"`
	City     string `json:"city"`
}

func toDTO(data *relativesdomain.Relatives) *RelativesInfoDTO {
	dto := &RelativesInfoDTO{
		Gender:   data.GetGender(),
		Dob:      data.GetDOB(),
		Address:  data.GetAddress(),
		Ward:     data.GetWard(),
		District: data.GetDistrict(),
		City:     data.GetCity(),
	}
	return dto
}

// get_accs_filter
type FilterAccountQuery struct {
	Paging common.Paging      `json:"paging"`
	Filter FieldFilterAccount `json:"filter"`
}

type FieldFilterAccount struct {
	Role        string `form:"role" json:"role"`
	FullName    string `form:"full-name" json:"full-name"`
	Email       string `form:"email" json:"email"`
	PhoneNumber string `form:"phone-number" json:"phone-number"`
}

type AccountDTO struct {
	Id          uuid.UUID `json:"id"`
	FullName    string    `json:"full-name"`
	Email       string    `json:"email"`
	PhoneNumber string    `json:"phone-number"`
	Avatar      string    `json:"avatar"`
	CreatedAt   time.Time `json:"created-at"`
}
