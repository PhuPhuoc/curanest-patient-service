package patientqueries

import (
	"context"

	patientdomain "github.com/PhuPhuoc/curanest-patient-service/module/patient/domain"
	"github.com/google/uuid"
)

type Queries struct {
	GetPatientByRelativesId *getPatientByRelIdHandler
}

type Builder interface {
	BuildPatientQueryRepo() PatientQueryRepo
}

func NewPatientQueryWithBuilder(b Builder) Queries {
	return Queries{
		GetPatientByRelativesId: NewGetPatientByRelIdHandler(
			b.BuildPatientQueryRepo(),
		),
	}
}

type PatientQueryRepo interface {
	GetPatientsByRelId(ctx context.Context, relativesId uuid.UUID) ([]patientdomain.Patient, error)
}
