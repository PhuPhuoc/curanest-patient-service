package builder

import (
	patientrepository "github.com/PhuPhuoc/curanest-patient-service/module/patient/infars/repository"
	patientcommands "github.com/PhuPhuoc/curanest-patient-service/module/patient/usecase/commands"
	relativesrepository "github.com/PhuPhuoc/curanest-patient-service/module/relatives/infars/repository"
	"github.com/jmoiron/sqlx"
)

type builderOfPatient struct {
	db *sqlx.DB
}

func NewPatientBuilder(db *sqlx.DB) builderOfPatient {
	return builderOfPatient{db: db}
}

func (s builderOfPatient) BuildPatientCmdRepo() patientcommands.PatientCommandRepo {
	return patientrepository.NewPatientRepo(s.db)
}

func (s builderOfPatient) BuildReletivesFetcherCmdRepo() patientcommands.ReletiveFetcher {
	return relativesrepository.NewRelativesRepo(s.db)
}
