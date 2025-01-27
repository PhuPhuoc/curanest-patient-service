package relativesrepository

import (
	"context"
	"fmt"

	"github.com/google/uuid"
)

func (r *relativesRepo) VerifyRelatives(ctx context.Context, relativeId *uuid.UUID) error {
	var verify bool
	query := `select exists (select 1 from ` + TABLE + ` where id=?)`
	if err := r.db.GetContext(ctx, &verify, query, relativeId.String()); err != nil {
		return err
	}
	if !verify {
		return fmt.Errorf("relative-id invalid")
	}
	return nil
}
