package pgdb

import (
	"context"

	"github.com/grum261/theater/internal/models"
	"github.com/jackc/pgx/v4/pgxpool"
	"golang.org/x/sync/errgroup"
)

type Costume struct {
	q *Queries
}

func newCostume(db *pgxpool.Pool) *Costume {
	return &Costume{
		q: newQueries(db),
	}
}

func (c *Costume) Create(ctx context.Context, p models.CostumeInsert) (models.CostumeReturn, error) {
	tx, err := c.q.db.Begin(ctx)
	if err != nil {
		return models.CostumeReturn{}, err
	}
	defer tx.Rollback(ctx)

	id, err := c.q.WithTx(tx).insertCostume(ctx, costumeInsertParams{
		Name:         p.Name,
		Description:  p.Description,
		Location:     p.Location,
		Condition:    p.Condition,
		Designer:     p.Designer,
		Size:         p.Size,
		Clothes:      p.ClothesId,
		IsDecor:      p.IsDecor,
		IsArchived:   p.IsArchived,
		ImageFront:   p.Image.Front,
		ImageBack:    p.Image.Back,
		ImageSideway: p.Image.Sideway,
		ImageDetails: p.Image.Details,
	})
	if err != nil {
		return models.CostumeReturn{}, err
	}

	_out := models.CostumeReturn{Id: id}

	clothes, err := c.q.WithTx(tx).selectClothesByIdArray(ctx, p.ClothesId)
	if err != nil {
		return models.CostumeReturn{}, err
	}

	for _, c := range clothes {
		_out.Clothes = append(_out.Clothes, models.Cloth{
			Id:        c.Id,
			Name:      c.Name,
			Type:      c.Type,
			Colors:    c.Colors,
			Materials: c.Materials,
		})
	}

	if err := tx.Commit(ctx); err != nil {
		return models.CostumeReturn{}, err
	}

	return _out, nil
}

func (c *Costume) Update(ctx context.Context, p models.CostumeUpdate) (models.CostumeReturn, error) {
	err := c.q.updateCostume(ctx, costumeUpdateParams{
		Id:           p.Id,
		Name:         p.Name,
		Description:  p.Description,
		Location:     p.Location,
		Condition:    p.Condition,
		Designer:     p.Designer,
		Size:         p.Size,
		Clothes:      p.ClothesId,
		IsDecor:      p.IsDecor,
		IsArchived:   p.IsArchived,
		ImageFront:   p.Image.Front,
		ImageBack:    p.Image.Back,
		ImageSideway: p.Image.Sideway,
		ImageDetails: p.Image.Details,
	})
	if err != nil {
		return models.CostumeReturn{}, err
	}

	_out := models.CostumeReturn{Id: p.Id}

	clothes, err := c.q.selectClothesByIdArray(ctx, p.ClothesId)
	if err != nil {
		return models.CostumeReturn{}, err
	}

	for _, c := range clothes {
		_out.Clothes = append(_out.Clothes, models.Cloth{
			Id:        c.Id,
			Name:      c.Name,
			Type:      c.Type,
			Colors:    c.Colors,
			Materials: c.Materials,
		})
	}

	return _out, nil
}

func (c *Costume) Delete(ctx context.Context, id int) error {
	return c.q.deleteCostume(ctx, id)
}

func (c *Costume) GetWithLimitOffset(ctx context.Context, limit, offset int) ([]models.CostumeSelect, error) {
	costumes, err := c.q.selectCostumesWithLimitOffset(ctx, limit, offset)
	if err != nil {
		return nil, err
	}

	g, ctx := errgroup.WithContext(ctx)

	clothCh := make(chan []int)
	g.Go(func() error {
		defer close(clothCh)
		for _, co := range costumes {
			select {
			case clothCh <- co.Clothes:
			case <-ctx.Done():
				return err
			}
		}

		return nil
	})

	costumesCh := make(chan models.CostumeSelect)
	for i := 0; i < 20; i++ {
		g.Go(func() error {
			j := 0
			for clothesId := range clothCh {
				clothes, err := c.q.selectClothesByIdArray(ctx, clothesId)
				if err != nil {
					return err
				}

				co := models.CostumeSelect{
					Id:          costumes[j].Id,
					Name:        costumes[j].Name,
					Description: costumes[j].Description,
					Location:    costumes[j].Location,
					Condition:   costumes[j].Condition,
					Designer:    costumes[j].Designer,
					IsDecor:     costumes[j].IsDecor,
					IsArchived:  costumes[j].IsArchived,
					Size:        costumes[j].Size,
					Image: models.Image{
						Front:   costumes[j].ImageFront,
						Back:    costumes[j].ImageBack,
						Sideway: costumes[j].ImageSideway,
						Details: costumes[j].ImageDetails,
					},
				}

				for _, cl := range clothes {
					co.Clothes = append(co.Clothes, models.Cloth{
						Id:        cl.Id,
						Name:      cl.Name,
						Type:      cl.Type,
						Colors:    cl.Colors,
						Materials: cl.Materials,
					})
				}

				select {
				case costumesCh <- co:
					j++
				case <-ctx.Done():
					return ctx.Err()
				}
			}

			return nil
		})
	}

	go func() {
		g.Wait()
		close(costumesCh)
	}()

	var _out []models.CostumeSelect

	for r := range costumesCh {
		_out = append(_out, r)
	}

	return _out, nil
}

func (c *Costume) MakeWriteOff(ctx context.Context, ids []int) error {
	return c.q.updateWriteOff(ctx, ids)
}
