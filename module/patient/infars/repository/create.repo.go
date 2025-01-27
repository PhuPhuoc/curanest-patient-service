package patientrepository

import (
	"context"

	"github.com/PhuPhuoc/curanest-patient-service/common"
	patientdomain "github.com/PhuPhuoc/curanest-patient-service/module/patient/domain"
)

func (r *patientRepo) Create(ctx context.Context, entity *patientdomain.Patient) error {
	dto := ToDTO(entity)
	query := common.GenerateSQLQueries(common.INSERT, TABLE, FIELD, nil)
	if _, err := r.db.NamedExec(query, dto); err != nil {
		return err
	}
	return nil
}
