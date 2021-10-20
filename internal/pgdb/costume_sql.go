package pgdb

import "context"

type costumeInsertParams struct {
	Name, Description, Location, Condition, Designer  string
	Size                                              int
	Clothes                                           []int
	IsDecor, IsArchived                               bool
	ImageFront, ImageBack, ImageSideway, ImageDetails string
}

type costumeUpdateParams struct {
	Id                                                int
	Name, Description, Location, Condition, Designer  string
	Size                                              int
	Clothes                                           []int
	IsDecor, IsArchived                               bool
	ImageFront, ImageBack, ImageSideway, ImageDetails string
}

func (q *Queries) insertCostume(ctx context.Context, p costumeInsertParams) (int, error) {
	var costumeId int

	tx, err := q.db.Begin(ctx)
	if err != nil {
		return 0, err
	}
	defer tx.Rollback(ctx)

	// name, description, designer, condition, is_decor, location, is_archived, size,
	// image_front, image_back, image_sideway, image_details
	if err := tx.QueryRow(
		ctx, costumeInsert, p.Name, p.Description, p.Designer, p.Condition, p.IsDecor, p.Location,
		p.IsArchived, p.Size, p.ImageFront, p.ImageBack, p.ImageSideway, p.ImageDetails,
	).Scan(&costumeId); err != nil {
		return 0, err
	}

	if _, err := tx.Exec(ctx, costumeClothesInsert, costumeId, p.Clothes); err != nil {
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
		ctx, costumeClothesUpdate, p.Name, p.Description, p.Condition, p.IsDecor, p.Location,
		p.IsArchived, p.Size, p.ImageFront, p.ImageBack, p.ImageSideway, p.ImageDetails, p.Designer,
	); err != nil {
		return err
	}

	if _, err := tx.Exec(ctx, costumeClothesUpdate, p.Id, p.Clothes); err != nil {
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
