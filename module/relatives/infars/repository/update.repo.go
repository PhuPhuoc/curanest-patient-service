package relativesrepository

import (
	"context"

	"github.com/PhuPhuoc/curanest-patient-service/common"
	relativesdomain "github.com/PhuPhuoc/curanest-patient-service/module/relatives/domain"
)

func (r *relativesRepo) Update(ctx context.Context, entity *relativesdomain.Relatives) error {
	dto := ToDTO(entity)
	where := "id=:id"
	query := common.GenerateSQLQueries(common.UPDATE, TABLE, UPDATE_FIELD, &where)
	if _, err := r.db.NamedExec(query, dto); err != nil {
		return err
	}
	return nil
}
