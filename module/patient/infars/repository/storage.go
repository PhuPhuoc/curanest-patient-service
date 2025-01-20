package patientrepository

import "github.com/jmoiron/sqlx"

type patientRepo struct {
	db *sqlx.DB
}

func NewPatientRepo(db *sqlx.DB) *patientRepo {
	return &patientRepo{
		db: db,
	}
}
