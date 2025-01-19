package relativesdomain

import "github.com/google/uuid"

type Relatives struct {
	id       uuid.UUID
	dob      string
	address  string
	ward     string
	district string
	city     string
}

func (a *Relatives) GetID() uuid.UUID {
	return a.id
}

func (a *Relatives) GetDOB() string {
	return a.dob
}

func (a *Relatives) GetAddress() string {
	return a.address
}

func (a *Relatives) GetWard() string {
	return a.ward
}

func (a *Relatives) GetDistrict() string {
	return a.district
}

func (a *Relatives) GetCity() string {
	return a.city
}

func NewRelatives(id uuid.UUID, dob, address, ward, district, city string) (*Relatives, error) {
	return &Relatives{
		id:       id,
		dob:      dob,
		address:  address,
		ward:     ward,
		district: district,
		city:     city,
	}, nil
}
