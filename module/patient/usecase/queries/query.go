package patientqueries

import (
	"context"

	patientdomain "github.com/PhuPhuoc/curanest-patient-service/module/patient/domain"
	"github.com/google/uuid"
)

type Queries struct {
	FindById                *findByIdHandler
	GetPatientByRelativesId *getPatientByRelIdHandler

	FindRelatviesIdByPatientId *findRelativesIdHandler
}

type Builder interface {
	BuildPatientQueryRepo() PatientQueryRepo
}

func NewPatientQueryWithBuilder(b Builder) Queries {
	return Queries{
		FindById: NewFindByIdHandler(
			b.BuildPatientQueryRepo(),
		),
		GetPatientByRelativesId: NewGetPatientByRelIdHandler(
			b.BuildPatientQueryRepo(),
		),

		FindRelatviesIdByPatientId: NewFindRelatviesIdHandler(
			b.BuildPatientQueryRepo(),
		),
	}
}

type PatientQueryRepo interface {
	FindByID(ctx context.Context, patientId uuid.UUID) (*patientdomain.Patient, error)
	GetPatientsByRelId(ctx context.Context, relativesId uuid.UUID) ([]patientdomain.Patient, error)
}
