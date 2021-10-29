package pgdb

const (
	clothesTypesInsert = `INSERT INTO clothes_types (name) (SELECT unnest($1::varchar[])) RETURNING id`
	clothesTypesUpdate = `UPDATE clothes_type SET name = $2 WHERE id = $1`
	clothesTypesDelete = `DELETE FROM clothes_types WHERE id = $1`
	clothesTypesSelect = `SELECT id, name FROM clothes_types`

	tagsInsertWithoutType = `INSERT INTO tags (name, type_id) (SELECT unnest($1::varchar[]), (SELECT id FROM tags_types WHERE name = 'Другое'))`
	tagsInsertWithTypes   = `INSERT INTO tags (name, type_id) VALUES ($1, $2)`
	tagUpdate             = `UPDATE tags (name, type_id) VALUES ($1, $2)`

	tagsTypesInsert   = `INSERT INTO tags_types (name) (SELECT unnest($1::varchar[])) RETURNING id`
	tagTypeNameUpdate = `UPDATE tags_types SET name = $2 WHERE id = $1`
)
