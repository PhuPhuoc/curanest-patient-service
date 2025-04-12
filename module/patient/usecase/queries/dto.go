package patientqueries

import (
	patientdomain "github.com/PhuPhuoc/curanest-patient-service/module/patient/domain"
	"github.com/google/uuid"
)

type PatientDTO struct {
	Id            uuid.UUID `json:"id"`
	RelativesId   uuid.UUID `json:"-"`
	FullName      string    `json:"full-name"`
	Gender        bool      `json:"gender"`
	Dob           string    `json:"dob"`
	PhoneNumber   string    `json:"phone-number"`
	Address       string    `json:"address"`
	Ward          string    `json:"ward"`
	District      string    `json:"district"`
	City          string    `json:"city"`
	DescPathology string    `json:"desc-pathology"`
	NoteForNurse  string    `json:"note-for-nurse"`
}

func toDTO(data *patientdomain.Patient) *PatientDTO {
	dto := &PatientDTO{
		Id:            data.GetID(),
		RelativesId:   data.GetRelativesID(),
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
