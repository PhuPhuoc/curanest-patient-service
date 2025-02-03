package patientrepository

import (
	"context"

	"github.com/PhuPhuoc/curanest-patient-service/common"
	patientdomain "github.com/PhuPhuoc/curanest-patient-service/module/patient/domain"
	"github.com/google/uuid"
)

func (r *patientRepo) GetPatientsByRelId(ctx context.Context, relativesId uuid.UUID) ([]patientdomain.Patient, error) {
	var dtos []PatientDTO
	where := "relatives_id=?"
	query := common.GenerateSQLQueries(common.FIND_WITH_CREATED_AT, TABLE, FIELD, &where)
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
