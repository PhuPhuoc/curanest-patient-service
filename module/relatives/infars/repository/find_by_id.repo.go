package relativesrepository

import (
	"context"

	relativesdomain "github.com/PhuPhuoc/curanest-patient-service/module/relatives/domain"
	"github.com/google/uuid"
)

func (r *relativesRepo) FindById(ctx context.Context, id uuid.UUID) (*relativesdomain.Relatives, error) {
	var dto RelativesDTO
	query := `select ` + field + ` from ` + table + ` where id=?`
	if err := r.db.Get(&dto, query, id); err != nil {
		return nil, err
	}
	return dto.ToEntity()
}
