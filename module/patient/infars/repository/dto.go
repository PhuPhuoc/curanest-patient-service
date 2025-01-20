package patientrepository

import (
	"time"

	patientdomain "github.com/PhuPhuoc/curanest-patient-service/module/patient/domain"
	"github.com/google/uuid"
)

const (
	table   = `patients`
	field   = `id, relatives_id, full_name, dob, phone_number, address, ward, district, city, desc_pathology, note_for_nurse`
	mapping = `:id, :relatives_id, :full_name, :dob, :phone_number, :address, :ward, :district, :city, :desc_pathology, :note_for_nurse`

	getField   = `id, relatives_id, full_name, dob, phone_number, address, ward, district, city, desc_pathology, note_for_nurse, created_at, updated_at`
	getMapping = `:id, :relatives_id, :full_name, :dob, :phone_number, :address, :ward, :district, :city, :desc_pathology, :note_for_nurse, :created_at, :updated_at`
)

type PatientDTO struct {
	Id            uuid.UUID  `db:"id"`
	RelativeId    uuid.UUID  `db:"relatives_id"`
	FullName      string     `db:"full_name"`
	Dob           string     `db:"dob"`
	PhoneNumber   string     `db:"phone_number"`
	Address       string     `db:"address"`
	Ward          string     `db:"ward"`
	District      string     `db:"district"`
	City          string     `db:"city"`
	DescPathology string     `db:"desc_pathology"`
	NoteForNurse  string     `db:"note_for_nurse"`
	CreatedAt     *time.Time `db:"created_at"`
	UpdateAt      *time.Time `db:"updated_at"`
}

func (dto *PatientDTO) ToEntity() (*patientdomain.Patient, error) {
	return patientdomain.NewPatient(
		dto.Id,
		dto.RelativeId,
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
		dto.UpdateAt,
	)
}

func ToDTO(data *patientdomain.Patient) *PatientDTO {
	dto := &PatientDTO{
		Id:            data.GetID(),
		RelativeId:    data.GetRelativesID(),
		FullName:      data.GetFullName(),
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
