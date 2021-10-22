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

func (c *Cloth) Create(ctx context.Context, name string, typeId int, colors, materials []int) (models.Cloth, error) {
	cl, err := c.q.insertCloth(ctx, clothInsertParams{
		Name:      name,
		TypeId:    typeId,
		Colors:    colors,
		Materials: materials,
	})
	if err != nil {
		return models.Cloth{}, err
	}

	return models.Cloth{
		Id:        cl.Id,
		Name:      name,
		Type:      cl.Type,
		Colors:    cl.Colors,
		Materials: cl.Materials,
	}, nil
}

func (c *Cloth) Update(ctx context.Context, id, typeId int, name string, colors, materials []int) (models.Cloth, error) {
	cl, err := c.q.updateCloth(ctx, clothUpdateParams{
		Id:        id,
		Name:      name,
		TypeId:    typeId,
		Colors:    colors,
		Materials: materials,
	})
	if err != nil {
		return models.Cloth{}, err
	}

	return models.Cloth{
		Id:        id,
		Name:      name,
		Type:      cl.Type,
		Colors:    cl.Colors,
		Materials: cl.Materials,
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
			Id:        c.Id,
			Name:      c.Name,
			Type:      c.Type,
			Colors:    c.Colors,
			Materials: c.Materials,
		})
	}

	return _out, nil
}

// func (c *Cloth) GetByIdArray(ctx context.Context, ids []int) ([]models.Cloth, error) {
// 	clothes, err := c.q.selectByIdArray(ctx, ids)
// 	if err != nil {
// 		return nil, err
// 	}

// 	var _out []models.Cloth

// 	for _, c := range clothes {
// 		_out = append(_out, models.Cloth{
// 			Id:        c.Id,
// 			Name:      c.Name,
// 			Type:      c.Type,
// 			Colors:    c.Colors,
// 			Materials: c.Materials,
// 		})
// 	}

// 	return _out, err
// }
