package relativesrepository

import (
	relativesdomain "github.com/PhuPhuoc/curanest-patient-service/module/relatives/domain"
	"github.com/google/uuid"
)

var (
	TABLE        = `relatives`
	FIELD        = []string{"id", "gender", "dob", "address", "ward", "district", "city"}
	UPDATE_FIELD = []string{"gender", "dob", "address", "ward", "district", "city"}
)

type RelativesDTO struct {
	Id       uuid.UUID `db:"id"`
	Gender   bool      `db:"gender"`
	Dob      string    `db:"dob"`
	Address  string    `db:"address"`
	Ward     string    `db:"ward"`
	District string    `db:"district"`
	City     string    `db:"city"`
}

func (dto *RelativesDTO) ToEntity() (*relativesdomain.Relatives, error) {
	return relativesdomain.NewRelatives(
		dto.Id,
		dto.Gender,
		dto.Dob,
		dto.Address,
		dto.Ward,
		dto.District,
		dto.City,
	)
}

func ToDTO(data *relativesdomain.Relatives) *RelativesDTO {
	dto := &RelativesDTO{
		Id:       data.GetID(),
		Gender:   data.GetGender(),
		Dob:      data.GetDOB(),
		Address:  data.GetAddress(),
		Ward:     data.GetWard(),
		District: data.GetDistrict(),
		City:     data.GetCity(),
	}
	return dto
}
