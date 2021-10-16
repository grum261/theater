package service

import (
	"context"
	"fmt"

	"github.com/grum261/theater/internal/models"
)

type CostumeRepository interface {
	Create(ctx context.Context, costume models.Costume, colorsId, materialsId, designersId, performacesId []int) (models.Costume, error)
	Update(ctx context.Context, costume models.Costume, colorsId, materialsId, designersId, performacesId []int) (models.Costume, error)
	FindById(ctx context.Context, id int) (models.Costume, error)
	Delete(ctx context.Context, id int) error
}

type Costume struct {
	repo CostumeRepository
}

func NewCostume(repo CostumeRepository) *Costume {
	return &Costume{
		repo: repo,
	}
}

func (c *Costume) Create(ctx context.Context, costume models.Costume, colorsId, materialsId, designersId, performacesId []int) (models.Costume, error) {
	_out, err := c.repo.Create(ctx, costume, colorsId, materialsId, designersId, performacesId)
	if err != nil {
		return models.Costume{}, fmt.Errorf("service.Costume.Create ошибка репозитория: %w", err)
	}

	return _out, nil
}

func (c *Costume) Update(ctx context.Context, costume models.Costume, colorsId, materialsId, designersId, performacesId []int) (models.Costume, error) {
	_out, err := c.repo.Update(ctx, costume, colorsId, materialsId, designersId, performacesId)
	if err != nil {
		return models.Costume{}, fmt.Errorf("service.Costume.Update ошибка репозитория: %w", err)
	}

	return _out, nil
}

func (c *Costume) FindById(ctx context.Context, id int) (models.Costume, error) {
	_out, err := c.repo.FindById(ctx, id)
	if err != nil {
		return models.Costume{}, fmt.Errorf("service.Costume.Find ошибка репозитория: %w", err)
	}

	return _out, nil
}

func (c *Costume) Delete(ctx context.Context, id int) error {
	if err := c.repo.Delete(ctx, id); err != nil {
		return err
	}

	return nil
}
