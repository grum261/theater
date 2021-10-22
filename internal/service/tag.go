package service

import (
	"context"

	"github.com/grum261/theater/internal/models"
)

type TagRepository interface {
	Create(ctx context.Context, names []string, table string) ([]models.Tag, error)
	Update(ctx context.Context, id int, name string, table string) error
	GetAll(ctx context.Context, table string) ([]models.Tag, error)
	Delete(ctx context.Context, id int, table string) error
}

type Tag struct {
	repo TagRepository
}

func newTag(repo TagRepository) *Tag {
	return &Tag{
		repo: repo,
	}
}

func (t *Tag) Create(ctx context.Context, names []string, table string) ([]models.Tag, error) {
	_out, err := t.repo.Create(ctx, names, table)
	if err != nil {
		return nil, err
	}

	return _out, nil
}

func (t *Tag) Update(ctx context.Context, id int, name string, table string) error {
	return t.repo.Update(ctx, id, name, table)
}

func (t *Tag) GetAll(ctx context.Context, table string) ([]models.Tag, error) {
	_out, err := t.repo.GetAll(ctx, table)
	if err != nil {
		return nil, err
	}

	return _out, nil
}

func (t *Tag) Delete(ctx context.Context, id int, table string) error {
	return t.repo.Delete(ctx, id, table)
}
