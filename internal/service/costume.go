package service

import (
	"context"

	"github.com/grum261/theater/internal/models"
)

type CostumeRepository interface {
	Create(ctx context.Context, p models.CostumeInsert) (models.CostumeReturn, error)
	Update(ctx context.Context, p models.CostumeUpdate) (models.CostumeReturn, error)
	Delete(ctx context.Context, id int) error
	GetWithLimitOffset(ctx context.Context, limit, offset int) ([]models.CostumeSelect, error)
	MakeWriteOff(ctx context.Context, ids []int) error
}

type Costume struct {
	repo CostumeRepository
}

func newCostume(repo CostumeRepository) *Costume {
	return &Costume{
		repo: repo,
	}
}

func (c *Costume) Create(ctx context.Context, p models.CostumeInsert) (models.CostumeReturn, error) {
	return c.repo.Create(ctx, p)
}

func (c *Costume) Update(ctx context.Context, p models.CostumeUpdate) (models.CostumeReturn, error) {
	return c.repo.Update(ctx, p)
}

func (c *Costume) Delete(ctx context.Context, id int) error {
	return c.repo.Delete(ctx, id)
}

func (c *Costume) MakeWriteOff(ctx context.Context, ids []int) error {
	return c.repo.MakeWriteOff(ctx, ids)
}

func (c *Costume) GetWithLimitOffset(ctx context.Context, limit, offset int) ([]models.CostumeSelect, error) {
	return c.repo.GetWithLimitOffset(ctx, limit, offset)
}
