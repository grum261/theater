package pgdb

import (
	"context"
	"time"
)

type performanceInsertParams struct {
	Name       string
	Location   string
	StartingAt time.Time
	Duration   int
	Costumes   []int
}

type performanceUpdateParams struct {
	performanceInsertParams
	Id int
}

func (q *Queries) insertPerformance(ctx context.Context, p performanceInsertParams) (int, error) {
	var id int

	if err := q.db.QueryRow(ctx, performanceInsert, p.Name, p.Location, p.StartingAt, p.Duration).Scan(&id); err != nil {
		return 0, err
	}

	if _, err := q.db.Exec(ctx, performanceCostumesInsert, id, p.Costumes); err != nil {
		return 0, err
	}

	return id, nil
}

func (q *Queries) updatePerformance(ctx context.Context, p performanceUpdateParams) error {
	if _, err := q.db.Exec(ctx, performanceUpdate, p.Id, p.Name, p.Location, p.StartingAt, p.Duration); err != nil {
		return err
	}

	if _, err := q.db.Exec(ctx, performanceCostumesUpdate, p.Id, p.Costumes); err != nil {
		return err
	}

	return nil
}

func (q *Queries) deletePerformace(ctx context.Context, id int) error {
	if _, err := q.db.Exec(ctx, performanceDelete, id); err != nil {
		return err
	}

	return nil
}
