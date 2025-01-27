package patientrepository

import (
	"context"

	"github.com/PhuPhuoc/curanest-patient-service/common"
	patientdomain "github.com/PhuPhuoc/curanest-patient-service/module/patient/domain"
)

func (r *patientRepo) Update(ctx context.Context, entity *patientdomain.Patient) error {
	dto := ToDTO(entity)
	where := "id=:id"
	query := common.GenerateSQLQueries(common.UPDATE, TABLE, UPDATE_FIELD, &where)
	if _, err := r.db.NamedExec(query, dto); err != nil {
		return err
	}
	return nil
}
