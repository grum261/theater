package pgdb

import (
	"context"

	"github.com/grum261/theater/internal/models"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Tag struct {
	q *Queries
}

func newTag(db *pgxpool.Pool) *Tag {
	return &Tag{
		q: newQueries(db),
	}
}

func (t *Tag) Create(ctx context.Context, names []string, table string) ([]models.Tag, error) {
	tags, err := t.q.insertTags(ctx, names, table)
	if err != nil {
		return nil, err
	}

	var _out []models.Tag

	for _, tag := range tags {
		_out = append(_out, models.Tag{
			Id:   tag.Id,
			Name: tag.Name,
		})
	}

	return _out, nil
}

func (t *Tag) Update(ctx context.Context, id int, name, table string) error {
	return t.q.updateTag(ctx, id, name, table)
}

func (t *Tag) Delete(ctx context.Context, id int, table string) error {
	return t.q.deleteTag(ctx, id, table)
}

func (t *Tag) GetAll(ctx context.Context, table string) ([]models.Tag, error) {
	tags, err := t.q.selectTags(ctx, table)
	if err != nil {
		return nil, err
	}

	var _out []models.Tag

	for _, tag := range tags {
		_out = append(_out, models.Tag{
			Id:   tag.Id,
			Name: tag.Name,
		})
	}

	return _out, nil
}
