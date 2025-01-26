package patientrepository

import (
	"context"

	patientdomain "github.com/PhuPhuoc/curanest-patient-service/module/patient/domain"
	"github.com/google/uuid"
)

func (r *patientRepo) FindByID(ctx context.Context, id uuid.UUID) (*patientdomain.Patient, error) {
	var dto PatientDTO
	query := `select ` + field + ` from ` + table + ` where id=?`
	if err := r.db.Get(&dto, query, id); err != nil {
		return nil, err
	}
	return dto.ToEntity()
}
