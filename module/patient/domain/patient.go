package patientdomain

import (
	"time"

	"github.com/google/uuid"
)

type Patient struct {
	id             uuid.UUID
	relatives_id   uuid.UUID
	full_name      string
	phone_number   string
	dob            string
	address        string
	ward           string
	district       string
	city           string
	desc_pathology string
	note_for_nurse string
	created_at     *time.Time
	updated_at     *time.Time
}

func NewPatient(
	id, relativesId uuid.UUID,
	fullName, dob, phoneNumber,
	address, ward, district, city,
	descPathology, noteForNurse string,
	createdAt, updatedAt *time.Time,
) (*Patient, error) {
	return &Patient{
		id:             id,
		relatives_id:   relativesId,
		full_name:      fullName,
		phone_number:   phoneNumber,
		dob:            dob,
		address:        address,
		ward:           ward,
		district:       district,
		city:           city,
		desc_pathology: descPathology,
		note_for_nurse: noteForNurse,
		created_at:     createdAt,
		updated_at:     updatedAt,
	}, nil
}

func (a *Patient) GetID() uuid.UUID {
	return a.id
}

func (a *Patient) GetRelativesID() uuid.UUID {
	return a.relatives_id
}

func (a *Patient) GetFullName() string {
	return a.full_name
}

func (a *Patient) GetPhoneNumber() string {
	return a.phone_number
}

func (a *Patient) GetDOB() string {
	return a.dob
}

func (a *Patient) GetAddress() string {
	return a.address
}

func (a *Patient) GetWard() string {
	return a.ward
}

func (a *Patient) GetDistrict() string {
	return a.district
}

func (a *Patient) GetCity() string {
	return a.city
}

func (a *Patient) GetDescPathology() string {
	return a.desc_pathology
}

func (a *Patient) GetNoteForNurse() string {
	return a.note_for_nurse
}

func (a *Patient) GetCreatedAt() *time.Time {
	return a.created_at
}

func (a *Patient) GetUpdatedAt() *time.Time {
	return a.updated_at
}
