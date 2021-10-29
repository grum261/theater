package pgdb

import (
	"context"
)

type clothInsertParams struct {
	Name   string
	TypeId int
}

type clothUpdateParams struct {
	Id int
	clothInsertParams
}

type clothReturn struct {
	Id         int
	Name, Type string
}

func (q *Queries) insertCloth(ctx context.Context, p clothInsertParams) (clothReturn, error) {
	var (
		clothId  int
		typeName string
	)

	if err := q.db.QueryRow(ctx, clothInsert, p.Name, p.TypeId).Scan(&clothId, &typeName); err != nil {
		return clothReturn{}, err
	}

	return clothReturn{
		Id:   clothId,
		Name: p.Name,
		Type: typeName,
	}, nil
}

func (q *Queries) updateCloth(ctx context.Context, p clothUpdateParams) (clothReturn, error) {
	var typeName string

	if err := q.db.QueryRow(ctx, clothUpdate, p.Id, p.Name, p.TypeId).Scan(&typeName); err != nil {
		return clothReturn{}, err
	}

	return clothReturn{
		Id:   p.Id,
		Name: p.Name,
		Type: typeName,
	}, nil
}

func (q *Queries) deleteCloth(ctx context.Context, id int) error {
	if _, err := q.db.Exec(ctx, clothDelete, id); err != nil {
		return err
	}

	return nil
}

func (q *Queries) selectWithLimitOffset(ctx context.Context, limit, offset int) ([]clothReturn, error) {
	rows, err := q.db.Query(ctx, selectClothesLimitOffset, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var r []clothReturn

	for rows.Next() {
		var c clothReturn

		if err := rows.Scan(&c.Id, &c.Name, &c.Type); err != nil {
			return nil, err
		}

		r = append(r, c)
	}

	return r, nil
}

func (q *Queries) selectClothesByCostumeId(ctx context.Context, costumeId int) ([]clothReturn, error) {
	rows, err := q.db.Query(ctx, selectClothesByCostumeId, costumeId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var clothes []clothReturn

	for rows.Next() {
		var c clothReturn

		if err := rows.Scan(&c.Id, &c.Name, &c.Type); err != nil {
			return nil, err
		}

		clothes = append(clothes, c)
	}

	return clothes, nil
}
