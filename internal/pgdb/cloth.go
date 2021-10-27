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

func (c *Cloth) Create(ctx context.Context, name, location, designer, condition string, typeId int, colors, materials []int) (models.Cloth, error) {
	cl, err := c.q.insertCloth(ctx, clothInsertParams{
		Name:      name,
		TypeId:    typeId,
		Location:  location,
		Designer:  designer,
		Condition: condition,
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
		Location:  location,
		Designer:  designer,
		Condition: condition,
		Colors:    cl.Colors,
		Materials: cl.Materials,
	}, nil
}

func (c *Cloth) Update(ctx context.Context, id, typeId int, name, location, designer, condition string, colors, materials []int) (models.Cloth, error) {
	cl, err := c.q.updateCloth(ctx, clothUpdateParams{
		Id:        id,
		Name:      name,
		TypeId:    typeId,
		Location:  location,
		Designer:  designer,
		Condition: condition,
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
		Location:  location,
		Designer:  designer,
		Condition: condition,
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
			Location:  c.Location,
			Designer:  c.Designer,
			Condition: c.Condition,
			Colors:    c.Colors,
			Materials: c.Materials,
		})
	}

	return _out, nil
}
