package pgdb

import (
	"context"

	"github.com/grum261/theater/internal/models"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Cloth struct {
	q *Queries
}

func newCloth(db *pgxpool.Pool) *Cloth {
	return &Cloth{
		q: newQueries(db),
	}
}

func (c *Cloth) Create(ctx context.Context, p models.ClothInsertUpdate) (models.Cloth, error) {
	cl, err := c.q.insertCloth(ctx, clothInsertParams{
		Name:   p.Name,
		TypeId: p.TypeId,
	})
	if err != nil {
		return models.Cloth{}, err
	}

	return models.Cloth{
		Id:   cl.Id,
		Name: p.Name,
		Type: cl.Type,
	}, nil
}

func (c *Cloth) Update(ctx context.Context, p models.ClothInsertUpdate) (models.Cloth, error) {
	cl, err := c.q.updateCloth(ctx, clothUpdateParams{
		Id: p.Id,
		clothInsertParams: clothInsertParams{
			Name:   p.Name,
			TypeId: p.TypeId,
		},
	})
	if err != nil {
		return models.Cloth{}, err
	}

	return models.Cloth{
		Id:   cl.Id,
		Name: p.Name,
		Type: cl.Type,
	}, nil
}

func (c *Cloth) Delete(ctx context.Context, id int) error {
	return c.q.deleteCloth(ctx, id)
}

func (c *Cloth) GetWithLimitOffset(ctx context.Context, limit, offset int) ([]models.Cloth, error) {
	clothes, err := c.q.selectWithLimitOffset(ctx, limit, offset)
	if err != nil {
		return nil, err
	}

	var _out []models.Cloth

	for _, c := range clothes {
		_out = append(_out, models.Cloth{
			Id:   c.Id,
			Name: c.Name,
			Type: c.Type,
		})
	}

	return _out, nil
}
