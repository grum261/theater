package service

import (
	"context"

	"github.com/grum261/theater/internal/models"
)

type ClothRepository interface {
	Create(ctx context.Context, p models.ClothInsertUpdate) (models.Cloth, error)
	Update(ctx context.Context, p models.ClothInsertUpdate) (models.Cloth, error)
	GetWithLimitOffset(ctx context.Context, limit, offset int) ([]models.Cloth, error)
	Delete(ctx context.Context, id int) error
}

type Cloth struct {
	repo ClothRepository
}

func newCloth(repo ClothRepository) *Cloth {
	return &Cloth{
		repo: repo,
	}
}

func (c *Cloth) Create(ctx context.Context, p models.ClothInsertUpdate) (models.Cloth, error) {
	// p.Name, p.Location, p.Designer, p.Condition, p.TypeId, p.Size, p.Colors, p.Materials
	_out, err := c.repo.Create(ctx, p)
	if err != nil {
		return models.Cloth{}, err
	}

	return _out, nil
}

func (c *Cloth) Update(ctx context.Context, p models.ClothInsertUpdate) (models.Cloth, error) {
	_out, err := c.repo.Update(ctx, p)
	if err != nil {
		return models.Cloth{}, err
	}

	return _out, nil
}

func (c *Cloth) GetWithLimitOffset(ctx context.Context, limit, offset int) ([]models.Cloth, error) {
	_out, err := c.repo.GetWithLimitOffset(ctx, limit, offset)
	if err != nil {
		return nil, err
	}

	return _out, nil
}

func (c *Cloth) Delete(ctx context.Context, id int) error {
	return c.repo.Delete(ctx, id)
}
