package patientcommands

import (
	"context"

	patientdomain "github.com/PhuPhuoc/curanest-patient-service/module/patient/domain"
	"github.com/google/uuid"
)

type Commands struct {
	CreatePatientProfile createPatientHandler
}

type Builder interface {
	BuildPatientCmdRepo() PatientCommandRepo
	BuildReletivesFetcherCmdRepo() ReletiveFetcher
}

func NewPatientCmdWithBuilder(b Builder) Commands {
	return Commands{
		CreatePatientProfile: *NewCreatePatientHandler(
			b.BuildPatientCmdRepo(),
			b.BuildReletivesFetcherCmdRepo(),
		),
	}
}

type PatientCommandRepo interface {
	Create(ctx context.Context, entity *patientdomain.Patient) error
}

type ReletiveFetcher interface {
	VerifyRelatives(ctx context.Context, relativeId *uuid.UUID) error
}
