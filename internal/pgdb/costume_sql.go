package pgdb

import (
	"context"
)

type InsertCostumeParams struct {
	Name, Description, Condition               string
	IsDecor                                    bool
	Size                                       int
	PerformaceId, Colors, Materials, Designers []int
}

type UpdateCostumeParams struct {
	InsertCostumeParams
	Id int
}

func (q *Queries) InsertCostume(ctx context.Context, params InsertCostumeParams) (Costumes, error) {
	var costumeId int

	tx, err := q.db.Begin(ctx)
	if err != nil {
		return Costumes{}, err
	}
	defer tx.Rollback(ctx)

	if err := tx.QueryRow(
		ctx, insertCostume, params.Name, params.Description,
		params.IsDecor, params.Condition, params.Size,
	).Scan(&costumeId); err != nil {
		return Costumes{}, err
	}

	var colors, materials, designers, performances []string

	if len(params.PerformaceId) != 0 {
		rows, err := q.db.Query(ctx, insertCostumePerformances, costumeId, params.PerformaceId)
		if err != nil {
			return Costumes{}, err
		}
		defer rows.Close()

		for rows.Next() {
			var performance string

			if err := rows.Scan(&performance); err != nil {
				return Costumes{}, err
			}

			performances = append(performances, performance)
		}
	}

	if len(params.Colors) != 0 {
		rows, err := tx.Query(ctx, insertCostumeColors, costumeId, params.Colors)
		if err != nil {
			return Costumes{}, err
		}
		defer rows.Close()

		for rows.Next() {
			var color string

			if err := rows.Scan(&color); err != nil {
				return Costumes{}, err
			}

			colors = append(colors, color)
		}
	}

	if len(params.Materials) != 0 {
		rows, err := tx.Query(ctx, insertCostumeMaterials, costumeId, params.Materials)
		if err != nil {
			return Costumes{}, err
		}
		defer rows.Close()

		for rows.Next() {
			var material string

			if err := rows.Scan(&material); err != nil {
				return Costumes{}, err
			}

			materials = append(materials, material)
		}
	}

	if len(params.Designers) != 0 {
		rows, err := tx.Query(ctx, insertCostumeDesigners, costumeId, params.Materials)
		if err != nil {
			return Costumes{}, err
		}
		defer rows.Close()

		for rows.Next() {
			var designer string

			if err := rows.Scan(&designer); err != nil {
				return Costumes{}, err
			}

			designers = append(designers, designer)
		}
	}

	if err := tx.Commit(ctx); err != nil {
		return Costumes{}, err
	}

	return Costumes{
		Id:           costumeId,
		Name:         params.Name,
		Description:  params.Description,
		IsDecor:      params.IsDecor,
		Condition:    params.Condition,
		Size:         params.Size,
		Performances: performances,
		Tags: Tags{
			Colors:    colors,
			Materials: materials,
			Designers: designers,
		},
	}, nil
}

func (q *Queries) UpdateCostume(ctx context.Context, params UpdateCostumeParams) (Costumes, error) {
	tx, err := q.db.Begin(ctx)
	if err != nil {
		return Costumes{}, nil
	}
	defer tx.Rollback(ctx)

	if _, err := tx.Exec(ctx, updateCostume, params.Id, params.Name, params.Description, params.IsDecor, params.Condition, params.Size); err != nil {
		return Costumes{}, nil
	}

	var colors, materials, designers []string

	if len(params.Colors) != 0 {
		rows, err := tx.Query(ctx, updateCostumeColors, params.Id, params.Colors)
		if err != nil {
			return Costumes{}, err
		}
		defer rows.Close()

		for rows.Next() {
			var color string

			if err := rows.Scan(&color); err != nil {
				return Costumes{}, err
			}

			colors = append(colors, color)
		}
	}

	if len(params.Materials) != 0 {
		rows, err := tx.Query(ctx, updateCostumeMaterials, params.Id, params.Materials)
		if err != nil {
			return Costumes{}, err
		}
		defer rows.Close()

		for rows.Next() {
			var material string

			if err := rows.Scan(&material); err != nil {
				return Costumes{}, err
			}

			materials = append(materials, material)
		}
	}

	if len(params.Designers) != 0 {
		rows, err := tx.Query(ctx, updateCostumeDesigners, params.Id, params.Materials)
		if err != nil {
			return Costumes{}, err
		}
		defer rows.Close()

		for rows.Next() {
			var designer string

			if err := rows.Scan(&designer); err != nil {
				return Costumes{}, err
			}

			designers = append(designers, designer)
		}
	}

	if err := tx.Commit(ctx); err != nil {
		return Costumes{}, err
	}

	return Costumes{
		Id:          params.Id,
		Name:        params.Name,
		Description: params.Description,
		IsDecor:     params.IsDecor,
		Condition:   params.Condition,
		Size:        params.Size,
		Tags: Tags{
			Colors:    colors,
			Materials: materials,
			Designers: designers,
		},
	}, nil
}

func (q *Queries) SelectById(ctx context.Context, id int) (Costumes, error) {
	var _out Costumes

	if err := q.db.QueryRow(ctx, selectCostumeById, id).Scan(&_out); err != nil {
		return Costumes{}, err
	}

	return Costumes{
		Id:           id,
		Name:         _out.Name,
		Description:  _out.Description,
		Condition:    _out.Condition,
		IsDecor:      _out.IsDecor,
		Size:         _out.Size,
		CreatedAt:    _out.CreatedAt,
		UpdatedAt:    _out.UpdatedAt,
		Performances: _out.Performances,
		Tags: Tags{
			Colors:    _out.Colors,
			Materials: _out.Materials,
			Designers: _out.Designers,
		},
	}, nil
}
