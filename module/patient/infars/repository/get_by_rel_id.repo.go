package patientrepository

import (
	"context"

	patientdomain "github.com/PhuPhuoc/curanest-patient-service/module/patient/domain"
	"github.com/google/uuid"
)

func (r *patientRepo) GetPatientsByRelId(ctx context.Context, relativesId uuid.UUID) ([]patientdomain.Patient, error) {
	query := "select " + field + " from " + table + ` where relatives_id=?`

	var dtos []PatientDTO
	if err := r.db.SelectContext(ctx, &dtos, query, relativesId); err != nil {
		return nil, err
	}

	entities := make([]patientdomain.Patient, len(dtos))
	for i := range dtos {
		entity, _ := dtos[i].ToEntity()
		entities[i] = *entity
	}
	return entities, nil
}
