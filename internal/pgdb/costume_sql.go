package pgdb

import (
	"context"
)

type costumeInsertParams struct {
	Name, Description, Designer, Condition, Size, Location string
	Clothes                                                []int
	ImageFront, ImageBack, ImageSideway, ImageDetails      string
	IsDecor, IsArchived                                    bool
}

type costumeUpdateParams struct {
	Id int
	costumeInsertParams
}

type costumeReturn struct {
	Id                                                     int
	Name, Description, Designer, Condition, Size, Location string
	ImageFront, ImageBack, ImageSideway, ImageDetails      string
	IsDecor, IsArchived                                    bool
	Tags                                                   []string
}

func (q *Queries) insertCostume(ctx context.Context, p costumeInsertParams) (int, error) {
	var costumeId int

	tx, err := q.db.Begin(ctx)
	if err != nil {
		return 0, err
	}
	defer tx.Rollback(ctx)

	if err := tx.QueryRow(
		ctx, costumeInsert, p.Name, p.Description, p.ImageFront, p.ImageBack, p.ImageSideway,
		p.ImageDetails, p.Designer, p.Location, p.Size, p.Condition, p.IsDecor, p.IsArchived,
	).Scan(&costumeId); err != nil {
		return 0, err
	}

	if _, err := tx.Exec(ctx, costumeClothesInsert, costumeId, p.Clothes); err != nil {
		return 0, err
	}

	if err := tx.Commit(ctx); err != nil {
		return 0, err
	}

	return costumeId, nil
}

func (q *Queries) updateCostume(ctx context.Context, p costumeUpdateParams) error {
	tx, err := q.db.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	if _, err := tx.Exec(
		ctx, costumeClothesUpdate, p.Name, p.Description, p.ImageFront, p.ImageBack, p.ImageSideway,
		p.ImageDetails, p.ImageDetails, p.Designer, p.Location, p.Size, p.Condition, p.IsDecor, p.IsArchived,
	); err != nil {
		return err
	}

	if _, err := tx.Exec(ctx, costumeClothesUpdate, p.Id, p.Clothes); err != nil {
		return err
	}

	if err := tx.Commit(ctx); err != nil {
		return err
	}

	return nil
}

func (q *Queries) deleteCostume(ctx context.Context, id int) error {
	if _, err := q.db.Exec(ctx, costumeDelete, id); err != nil {
		return err
	}

	return nil
}

func (q *Queries) selectCostumesWithLimitOffset(ctx context.Context, limit, offset int) ([]costumeReturn, error) {
	rows, err := q.db.Query(ctx, costumeSelectWithLimitOffset, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var costumes []costumeReturn

	for rows.Next() {
		var c costumeReturn

		if err := rows.Scan(
			&c.Id, &c.Name, &c.Description, &c.ImageFront, &c.ImageBack, &c.ImageSideway,
			&c.ImageDetails, &c.Designer, &c.Location, &c.Size, &c.Condition, &c.IsDecor, &c.IsArchived,
		); err != nil {
			return nil, err
		}

		costumes = append(costumes, c)
	}

	return costumes, nil
}

func (q *Queries) updateWriteOff(ctx context.Context, ids []int) error {
	if _, err := q.db.Exec(ctx, costumeWriteOff, ids); err != nil {
		return err
	}

	return nil
}

func (q *Queries) selectCostumesAndClothesByPerformance(ctx context.Context, performanceId int) ([]costumeReturn, error) {
	rows, err := q.db.Query(ctx, costumeSelectByPerformanceId, performanceId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var costumes []costumeReturn

	for rows.Next() {
		var c costumeReturn

		if err := rows.Scan(
			&c.Id, &c.Name, &c.Description, &c.ImageFront, &c.ImageBack, &c.ImageSideway,
			&c.ImageDetails, &c.Designer, &c.Location, &c.Size, &c.Condition, &c.IsDecor, &c.IsArchived,
		); err != nil {
			return nil, err
		}

		costumes = append(costumes, c)
	}

	return costumes, nil
}
