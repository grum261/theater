package pgdb

import (
	"context"

	"github.com/grum261/theater/internal/models"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Costume struct {
	q *Queries
}

func NewCostume(db *pgxpool.Pool) *Costume {
	return &Costume{
		q: NewQueries(db),
	}
}

func (c *Costume) Create(ctx context.Context, costume models.Costume, colorsId, materialsId, designersId, performacesId []int) (models.Costume, error) {
	_out, err := c.q.InsertCostume(ctx, InsertCostumeParams{
		Name:         costume.Name,
		Description:  costume.Description,
		Condition:    costume.Condition,
		IsDecor:      costume.IsDecor,
		Size:         costume.Size,
		PerformaceId: performacesId,
		Colors:       colorsId,
		Materials:    materialsId,
		Designers:    designersId,
	})
	if err != nil {
		return models.Costume{}, err
	}

	return models.Costume{
		Id:          _out.Id,
		Name:        _out.Name,
		Description: _out.Description,
		Condition:   _out.Condition,
		Size:        _out.Size,
		IsDecor:     _out.IsDecor,
		Performaces: _out.Performances,
		Tags: models.Tag{
			Colors:    _out.Colors,
			Materials: _out.Materials,
			Designers: _out.Designers,
		},
	}, nil
}

func (c *Costume) Update(ctx context.Context, costume models.Costume, colorsId, materialsId, designersId, performacesId []int) (models.Costume, error) {
	_out, err := c.q.UpdateCostume(ctx, UpdateCostumeParams{
		Id: costume.Id,
		InsertCostumeParams: InsertCostumeParams{
			Name:         costume.Name,
			Description:  costume.Description,
			Condition:    costume.Condition,
			IsDecor:      costume.IsDecor,
			Size:         costume.Size,
			PerformaceId: performacesId,
			Colors:       colorsId,
			Materials:    materialsId,
			Designers:    designersId,
		},
	})
	if err != nil {
		return models.Costume{}, err
	}

	return models.Costume{
		Id:          _out.Id,
		Name:        _out.Name,
		Description: _out.Description,
		Condition:   _out.Condition,
		Size:        _out.Size,
		IsDecor:     _out.IsDecor,
		Performaces: _out.Performances,
		Tags: models.Tag{
			Colors:    _out.Colors,
			Materials: _out.Materials,
			Designers: _out.Designers,
		},
	}, nil
}

func (c *Costume) FindById(ctx context.Context, id int) (models.Costume, error) {
	_out, err := c.q.SelectById(ctx, id)
	if err != nil {
		return models.Costume{}, err
	}

	return models.Costume{
		Id:          id,
		Name:        _out.Name,
		Description: _out.Description,
		Condition:   _out.Condition,
		Size:        _out.Size,
		IsDecor:     _out.IsDecor,
		Performaces: _out.Performances,
		Tags: models.Tag{
			Colors:    _out.Colors,
			Materials: _out.Materials,
			Designers: _out.Designers,
		},
	}, nil
}

func (c *Costume) Delete(ctx context.Context, id int) error {
	return nil
}
