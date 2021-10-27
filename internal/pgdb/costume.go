package pgdb

import (
	"context"

	"github.com/grum261/theater/internal/models"
	"github.com/jackc/pgx/v4/pgxpool"
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
			Location:  c.Location,
			Designer:  c.Designer,
			Condition: c.Condition,
			Size:      c.Size,
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
			Location:  c.Location,
			Designer:  c.Designer,
			Condition: c.Condition,
			Size:      c.Size,
			Colors:    c.Colors,
			Materials: c.Materials,
		})
	}

	return _out, nil
}

func (c *Costume) Delete(ctx context.Context, id int) error {
	return c.q.deleteCostume(ctx, id)
}

func (c *Costume) GetWithLimitOffset(ctx context.Context, limit, offset int) ([]models.Costume, error) {
	costumes, err := c.q.selectCostumesWithLimitOffset(ctx, limit, offset)
	if err != nil {
		return nil, err
	}

	var _out []models.Costume

	for _, co := range costumes {
		cos := models.Costume{
			Id:          co.Id,
			Name:        co.Name,
			Description: co.Description,
			IsDecor:     co.IsDecor,
			IsArchived:  co.IsArchived,
			Image: models.Image{
				Front:   co.ImageFront,
				Back:    co.ImageBack,
				Sideway: co.ImageSideway,
				Details: co.ImageDetails,
			},
		}

		clothes, err := c.q.selectClothesByIdArray(ctx, co.Clothes)
		if err != nil {
			return nil, err
		}

		for _, cl := range clothes {
			cos.Clothes = append(cos.Clothes, models.Cloth{
				Id:        cl.Id,
				Name:      cl.Name,
				Type:      cl.Type,
				Location:  cl.Location,
				Designer:  cl.Designer,
				Condition: cl.Condition,
				Size:      cl.Size,
				Colors:    cl.Colors,
				Materials: cl.Materials,
			})
		}

		_out = append(_out, cos)
	}

	return _out, nil
}

func (c *Costume) MakeWriteOff(ctx context.Context, ids []int) error {
	return c.q.updateWriteOff(ctx, ids)
}
