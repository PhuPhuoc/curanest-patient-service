package patientrepository

import (
	"context"

	patientdomain "github.com/PhuPhuoc/curanest-patient-service/module/patient/domain"
)

func (r *patientRepo) Update(ctx context.Context, entity *patientdomain.Patient) error {
	setField := SetFieldForPatient()
	dto := ToDTO(entity)
	query := `update ` + table + ` set ` + setField + ` where id = :id`

	if _, err := r.db.NamedExec(query, dto); err != nil {
		return err
	}
	return nil
}
