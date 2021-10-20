package pgdb

import (
	"context"

	"github.com/grum261/theater/internal/models"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Costume struct {
	q *Queries
}

func newCostume(db *pgxpool.Pool) *Costume {
	return &Costume{
		q: newQueries(db),
	}
}

func (c *Costume) Create(ctx context.Context, p models.CostumeInsert) (models.CostumeReturn, error) {
	id, err := c.q.insertCostume(ctx, costumeInsertParams{
		Name:         p.Name,
		Description:  p.Description,
		Location:     p.Location,
		Condition:    p.Condition,
		Designer:     p.Designer,
		Size:         p.Size,
		Clothes:      p.ClothesId,
		IsDecor:      p.IsDecor,
		IsArchived:   p.IsArchived,
		ImageFront:   p.Image.Front,
		ImageBack:    p.Image.Back,
		ImageSideway: p.Image.Sideway,
		ImageDetails: p.Image.Details,
	})
	if err != nil {
		return models.CostumeReturn{}, err
	}

	_out := models.CostumeReturn{Id: id}

	cl := Cloth{}

	_out.Clothes, err = cl.GetByIdArray(ctx, p.ClothesId)
	if err != nil {
		return _out, err
	}

	return _out, nil
}

func (c *Costume) Update(ctx context.Context, p models.CostumeUpdate) (models.CostumeReturn, error) {
	err := c.q.updateCostume(ctx, costumeUpdateParams{
		Id:           p.Id,
		Name:         p.Name,
		Description:  p.Description,
		Location:     p.Location,
		Condition:    p.Condition,
		Designer:     p.Designer,
		Size:         p.Size,
		Clothes:      p.ClothesId,
		IsDecor:      p.IsDecor,
		IsArchived:   p.IsArchived,
		ImageFront:   p.Image.Front,
		ImageBack:    p.Image.Back,
		ImageSideway: p.Image.Sideway,
		ImageDetails: p.Image.Details,
	})
	if err != nil {
		return models.CostumeReturn{}, err
	}

	_out := models.CostumeReturn{Id: p.Id}

	cl := Cloth{}

	_out.Clothes, err = cl.GetByIdArray(ctx, p.ClothesId)
	if err != nil {
		return _out, err
	}

	return _out, nil
}

func (c *Costume) Delete(ctx context.Context, id int) error {
	return c.q.deleteCostume(ctx, id)
}
