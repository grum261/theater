package pgdb

import "github.com/jackc/pgx/v4/pgxpool"

type Store struct {
	*Tag
	*Cloth
	*Costume
	*Performance
}

func NewStore(db *pgxpool.Pool) *Store {
	return &Store{
		newTag(db),
		newCloth(db),
		newCostume(db),
		newPerformance(db),
	}
}
