package service

import (
	"context"
	"time"

	"github.com/grum261/theater/internal/models"
)

type PerformanceRepository interface {
	Create(
		ctx context.Context, name, location, string,
		startingAt time.Time, duration time.Duration, costumes []int,
	) (models.PerformanceReturn, error)

	Update(
		ctx context.Context, id int, name, location string,
		startingAt time.Time, duration time.Duration, costumes []int,
	) (models.PerformanceReturn, error)

	Delete(ctx context.Context, id int) error
	// GetPerformancesByDate(ctx context.Context) []models.Performance
}
