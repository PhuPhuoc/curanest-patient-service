package relativesrepository

import (
	"context"

	relativesdomain "github.com/PhuPhuoc/curanest-patient-service/module/relatives/domain"
)

func (r *relativesRepo) Create(ctx context.Context, entity *relativesdomain.Relatives) error {
	dto := ToDTO(entity)
	query := `insert into ` + table + ` (` + field + `) values (` + mapping + `)`
	if _, err := r.db.NamedExec(query, dto); err != nil {
		return err
	}
	return nil
}
