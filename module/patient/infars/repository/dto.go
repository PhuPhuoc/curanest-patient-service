package patientrepository

import (
	"time"

	patientdomain "github.com/PhuPhuoc/curanest-patient-service/module/patient/domain"
	"github.com/google/uuid"
)

var (
	TABLE        = `patients`
	FIELD        = []string{"id", "relatives_id", "full_name", "gender", "dob", "phone_number", "address", "ward", "district", "city", "desc_pathology", "note_for_nurse"}
	UPDATE_FIELD = []string{"full_name", "gender", "dob", "phone_number", "address", "ward", "district", "city", "desc_pathology", "note_for_nurse"}
)

type PatientDTO struct {
	Id            uuid.UUID  `db:"id"`
	RelativeId    uuid.UUID  `db:"relatives_id"`
	FullName      string     `db:"full_name"`
	Gender        bool       `db:"gender"`
	Dob           string     `db:"dob"`
	PhoneNumber   string     `db:"phone_number"`
	Address       string     `db:"address"`
	Ward          string     `db:"ward"`
	District      string     `db:"district"`
	City          string     `db:"city"`
	DescPathology string     `db:"desc_pathology"`
	NoteForNurse  string     `db:"note_for_nurse"`
	CreatedAt     *time.Time `db:"created_at"`
}

func (dto *PatientDTO) ToEntity() (*patientdomain.Patient, error) {
	return patientdomain.NewPatient(
		dto.Id,
		dto.RelativeId,
		dto.Gender,
		dto.FullName,
		dto.Dob,
		dto.PhoneNumber,
		dto.Address,
		dto.Ward,
		dto.District,
		dto.City,
		dto.DescPathology,
		dto.NoteForNurse,
		dto.CreatedAt,
	)
}

func ToDTO(data *patientdomain.Patient) *PatientDTO {
	dto := &PatientDTO{
		Id:            data.GetID(),
		RelativeId:    data.GetRelativesID(),
		FullName:      data.GetFullName(),
		Gender:        data.GetGender(),
		Dob:           data.GetDOB(),
		PhoneNumber:   data.GetPhoneNumber(),
		Address:       data.GetAddress(),
		Ward:          data.GetWard(),
		District:      data.GetDistrict(),
		City:          data.GetCity(),
		DescPathology: data.GetDescPathology(),
		NoteForNurse:  data.GetNoteForNurse(),
	}
	return dto
}
