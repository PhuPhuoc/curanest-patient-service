package relativesrepository

import "github.com/jmoiron/sqlx"

type relativesRepo struct {
	db *sqlx.DB
}

func NewRelativesRepo(db *sqlx.DB) *relativesRepo {
	return &relativesRepo{
		db: db,
	}
}
