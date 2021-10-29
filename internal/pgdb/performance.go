package pgdb

import (
	"context"
	"fmt"

	"github.com/grum261/theater/internal/models"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Performance struct {
	q *Queries
}

func newPerformance(db *pgxpool.Pool) *Performance {
	return &Performance{
		q: newQueries(db),
	}
}

func (p *Performance) Create(ctx context.Context, args models.PerformanceInsertUpdate) (models.PerformanceReturn, error) {
	tx, err := p.q.db.Begin(ctx)
	if err != nil {
		return models.PerformanceReturn{}, err
	}
	defer tx.Rollback(ctx)

	id, err := p.q.WithTx(tx).insertPerformance(ctx, performanceInsertParams{
		Name:       args.Name,
		Location:   args.Location,
		StartingAt: args.StartingAt,
		Duration:   args.Duration,
		Costumes:   args.Costumes,
	})
	if err != nil {
		return models.PerformanceReturn{}, err
	}

	if err := tx.Commit(ctx); err != nil {
		return models.PerformanceReturn{}, err
	}

	costumes, err := p.q.selectCostumesAndClothesByPerformance(ctx, id)
	if err != nil {
		return models.PerformanceReturn{}, err
	}

	performances := models.PerformanceReturn{Id: id}

	for _, co := range costumes {
		cos := models.Costume{
			Id:          co.Id,
			Name:        co.Name,
			Description: co.Description,
			Image: models.Image{
				Front:   co.ImageFront,
				Back:    co.ImageBack,
				Sideway: co.ImageSideway,
				Details: co.ImageDetails,
			},
		}

		clothes, err := p.q.selectClothesByCostumeId(ctx, co.Id)
		if err != nil {
			return models.PerformanceReturn{}, err
		}

		for _, cl := range clothes {
			cos.Clothes = append(cos.Clothes, models.Cloth{
				Id:   cl.Id,
				Name: cl.Name,
				Type: cl.Type,
			})
		}

		performances.Costumes = append(performances.Costumes, cos)
	}

	return performances, nil
}

func (p *Performance) Update(ctx context.Context, args models.PerformanceInsertUpdate) (models.PerformanceReturn, error) {
	tx, err := p.q.db.Begin(ctx)
	if err != nil {
		return models.PerformanceReturn{}, err
	}
	defer tx.Rollback(ctx)

	if err := p.q.WithTx(tx).updatePerformance(ctx, performanceUpdateParams{
		Id: args.Id,
		performanceInsertParams: performanceInsertParams{
			Name:       args.Name,
			Location:   args.Location,
			StartingAt: args.StartingAt,
			Duration:   args.Duration,
			Costumes:   args.Costumes,
		},
	}); err != nil {
		return models.PerformanceReturn{}, err
	}

	if err := tx.Commit(ctx); err != nil {
		return models.PerformanceReturn{}, err
	}

	costumes, err := p.q.selectCostumesAndClothesByPerformance(ctx, args.Id)
	if err != nil {
		return models.PerformanceReturn{}, err
	}

	performances := models.PerformanceReturn{Id: args.Id}

	for _, co := range costumes {
		cos := models.Costume{
			Id:          co.Id,
			Name:        co.Name,
			Description: co.Description,
			Image: models.Image{
				Front:   co.ImageFront,
				Back:    co.ImageBack,
				Sideway: co.ImageSideway,
				Details: co.ImageDetails,
			},
		}

		clothes, err := p.q.selectClothesByCostumeId(ctx, co.Id)
		if err != nil {
			return models.PerformanceReturn{}, err
		}

		for _, cl := range clothes {
			cos.Clothes = append(cos.Clothes, models.Cloth{
				Id:   cl.Id,
				Name: cl.Name,
				Type: cl.Type,
			})
		}

		performances.Costumes = append(performances.Costumes, cos)
	}

	return performances, nil
}

func (p *Performance) Delete(ctx context.Context, id int) error {
	return p.q.deletePerformace(ctx, id)
}

func (p *Performance) GetNearest(ctx context.Context) ([]models.Performance, error) {
	return nil, fmt.Errorf("метод не имплементирован")
}
