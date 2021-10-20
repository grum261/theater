package pgdb

const (
	colorsInsert = `INSERT INTO colors (name) (SELECT unnest($1::varchar[])) RETURNING id`
	colorsUpdate = `UPDATE colors SET name = $2 WHERE id = $1`
	colorsSelect = `SELECT id, name FROM colors`
	colorsDelete = `DELETE FROM colors WHERE id = $1`

	materialsInsert = `INSERT INTO materials (name) (SELECT unnest($1::varchar[])) RETURNING id`
	materialsUpdate = `UPDATE materials SET name = $2 WHERE id = $1`
	materialsSelect = `SELECT id, name FROM materials`
	materialsDelete = `DELETE FROM materials WHERE id = $1`

	clothesTypesInsert = `INSERT INTO clothes_types (name) (SELECT unnest($1::varchar[])) RETURNING id`
	clothesTypesUpdate = `UPDATE clothes_type SET name = $2 WHERE id = $1`
	clothesTypesDelete = `DELETE FROM clothes_types WHERE id = $1`
	clothesTypesSelect = `SELECT id, name FROM clothes_types`
)
