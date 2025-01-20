package patientrepository

import (
	"context"

	patientdomain "github.com/PhuPhuoc/curanest-patient-service/module/patient/domain"
)

func (r *patientRepo) Create(ctx context.Context, entity *patientdomain.Patient) error {
	dto := ToDTO(entity)
	query := `insert into ` + table + ` (` + field + `) values (` + mapping + `)`
	if _, err := r.db.NamedExec(query, dto); err != nil {
		return err
	}
	return nil
}
