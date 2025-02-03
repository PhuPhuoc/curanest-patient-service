package patientrepository

import (
	"context"

	"github.com/PhuPhuoc/curanest-patient-service/common"
	patientdomain "github.com/PhuPhuoc/curanest-patient-service/module/patient/domain"
	"github.com/google/uuid"
)

func (r *patientRepo) FindByID(ctx context.Context, id uuid.UUID) (*patientdomain.Patient, error) {
	var dto PatientDTO
	where := "id=?"
	query := common.GenerateSQLQueries(common.FIND_WITH_CREATED_AT, TABLE, FIELD, &where)
	if err := r.db.Get(&dto, query, id); err != nil {
		return nil, err
	}

	return dto.ToEntity()
}
