package relativesrepository

import (
	"context"

	"github.com/PhuPhuoc/curanest-patient-service/common"
	relativesdomain "github.com/PhuPhuoc/curanest-patient-service/module/relatives/domain"
)

func (r *relativesRepo) Create(ctx context.Context, entity *relativesdomain.Relatives) error {
	dto := ToDTO(entity)
	query := common.GenerateSQLQueries(common.INSERT, TABLE, FIELD, nil)
	if _, err := r.db.NamedExec(query, dto); err != nil {
		return err
	}
	return nil
}
