package relativesrepository

import (
	"context"

	"github.com/PhuPhuoc/curanest-patient-service/common"
	relativesdomain "github.com/PhuPhuoc/curanest-patient-service/module/relatives/domain"
	"github.com/google/uuid"
)

func (r *relativesRepo) FindById(ctx context.Context, id uuid.UUID) (*relativesdomain.Relatives, error) {
	var dto RelativesDTO
	where := "id=?"
	query := common.GenerateSQLQueries(common.FIND_WITH_OUT_CREATED_AT, TABLE, FIELD, &where)
	if err := r.db.Get(&dto, query, id); err != nil {
		return nil, err
	}
	return dto.ToEntity()
}
