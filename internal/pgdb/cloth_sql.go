package pgdb

import (
	"context"
)

type clothInsertParams struct {
	Name              string
	TypeId            int
	Designer          string
	Location          string
	Condition         string
	Size              int
	Colors, Materials []int
}

type clothUpdateParams struct {
	Id                int
	Name              string
	TypeId            int
	Location          string
	Designer          string
	Condition         string
	Size              int
	Colors, Materials []int
}

type clothReturn struct {
	Id                int
	Name, Type        string
	Location          string
	Designer          string
	Condition         string
	Size              int
	Colors, Materials []string
}

func (q *Queries) insertCloth(ctx context.Context, p clothInsertParams) (clothReturn, error) {
	var (
		clothId  int
		typeName string
	)

	tx, err := q.db.Begin(ctx)
	if err != nil {
		return clothReturn{}, err
	}
	defer tx.Rollback(ctx)

	if err := tx.QueryRow(ctx, clothInsert, p.Name, p.TypeId, p.Location, p.Designer, p.Condition, p.Size).Scan(&clothId, &typeName); err != nil {
		return clothReturn{}, err
	}

	var colors, materials []string

	if len(p.Colors) != 0 {
		rows, err := tx.Query(ctx, clothColorsInsert, clothId, p.Colors)
		if err != nil {
			return clothReturn{}, err
		}
		defer rows.Close()

		for rows.Next() {
			var color string

			if err := rows.Scan(&color); err != nil {
				return clothReturn{}, err
			}

			colors = append(colors, color)
		}
	}

	if len(p.Materials) != 0 {
		rows, err := tx.Query(ctx, clothMaterialsInsert, clothId, p.Materials)
		if err != nil {
			return clothReturn{}, err
		}
		defer rows.Close()

		for rows.Next() {
			var material string

			if err := rows.Scan(&material); err != nil {
				return clothReturn{}, err
			}

			materials = append(materials, material)
		}
	}

	if err := tx.Commit(ctx); err != nil {
		return clothReturn{}, err
	}

	return clothReturn{
		Id:        clothId,
		Name:      p.Name,
		Type:      typeName,
		Colors:    colors,
		Materials: materials,
	}, nil
}

func (q *Queries) updateCloth(ctx context.Context, p clothUpdateParams) (clothReturn, error) {
	var typeName string

	tx, err := q.db.Begin(ctx)
	if err != nil {
		return clothReturn{}, err
	}
	defer tx.Rollback(ctx)

	if err := tx.QueryRow(ctx, clothUpdate, p.Id, p.Name, p.TypeId, p.Location, p.Designer, p.Condition, p.Size).Scan(&typeName); err != nil {
		return clothReturn{}, err
	}

	var colors, materials []string

	if len(p.Colors) != 0 {
		rows, err := tx.Query(ctx, clothColorsUpdate, p.Id, p.Colors)
		if err != nil {
			return clothReturn{}, err
		}
		defer rows.Close()

		for rows.Next() {
			var color string

			if err := rows.Scan(&color); err != nil {
				return clothReturn{}, err
			}

			colors = append(colors, color)
		}
	}

	if len(p.Materials) != 0 {
		rows, err := tx.Query(ctx, clothMaterialsUpdate, p.Id, p.Materials)
		if err != nil {
			return clothReturn{}, err
		}
		defer rows.Close()

		for rows.Next() {
			var material string

			if err := rows.Scan(&material); err != nil {
				return clothReturn{}, err
			}

			materials = append(materials, material)
		}
	}

	if err := tx.Commit(ctx); err != nil {
		return clothReturn{}, err
	}

	return clothReturn{
		Id:        p.Id,
		Name:      p.Name,
		Type:      typeName,
		Colors:    colors,
		Materials: materials,
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

		if err := rows.Scan(&c.Id, &c.Name, &c.Type, &c.Location, &c.Designer, &c.Condition, &c.Size, &c.Colors, &c.Materials); err != nil {
			return nil, err
		}

		r = append(r, c)
	}

	return r, nil
}

func (q *Queries) selectClothesByIdArray(ctx context.Context, ids []int) ([]clothReturn, error) {
	rows, err := q.db.Query(ctx, selectClothesByIdArray, ids)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var clothes []clothReturn

	for rows.Next() {
		var c clothReturn

		if err := rows.Scan(&c.Id, &c.Name, &c.Type, &c.Location, &c.Designer, &c.Condition, &c.Size, &c.Colors, &c.Materials); err != nil {
			return nil, err
		}

		clothes = append(clothes, c)
	}

	return clothes, nil
}
