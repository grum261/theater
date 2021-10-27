package service

import (
	"context"

	"github.com/grum261/theater/internal/models"
)

type PerformanceRepository interface {
	Create(ctx context.Context, args models.PerformanceInsertUpdate) (models.PerformanceReturn, error)
	Update(ctx context.Context, args models.PerformanceInsertUpdate) (models.PerformanceReturn, error)
	GetNearest(ctx context.Context) ([]models.Performance, error)
	Delete(ctx context.Context, id int) error
}

type Performance struct {
	repo PerformanceRepository
}

func newPerformance(repo PerformanceRepository) *Performance {
	return &Performance{
		repo: repo,
	}
}

func (p *Performance) Create(ctx context.Context, args models.PerformanceInsertUpdate) (models.PerformanceReturn, error) {
	return p.repo.Create(ctx, args)
}

func (p *Performance) Update(ctx context.Context, args models.PerformanceInsertUpdate) (models.PerformanceReturn, error) {
	return p.repo.Update(ctx, args)
}

func (p *Performance) GetNearest(ctx context.Context) ([]models.Performance, error) {
	return p.repo.GetNearest(ctx)
}

func (p *Performance) Delete(ctx context.Context, id int) error {
	return p.repo.Delete(ctx, id)
}
