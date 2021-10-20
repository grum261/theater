package pgdb

const (
	clothInsert = `
	INSERT INTO clothes (name, type_id) VALUES ($1, $2) RETURNING id, (SELECT name FROM clothes_types WHERE id = $2)`
	clothColorsInsert = `
	WITH ins AS (INSERT INTO clothes_colors (cloth_id, color_id) (SELECT $1, unnest($2::int[])))
	SELECT name FROM colors WHERE id = any($2::int[])`
	clothMaterialsInsert = `
	WITH ins AS (INSERT INTO clothes_materials (cloth_id, material_id) (SELECT $1, unnest($2::int[])))
	SELECT name FROM materials WHERE id = any($2::int[])`

	clothUpdate       = `UPDATE clothes SET name = $2, type_id = $3 WHERE id = $1 RETURNING (SELECT name FROM clothes_types WHERE id = $3)`
	clothColorsUpdate = `
	WITH d AS (
		DELETE FROM clothes_colors WHERE cloth_id = $1
	), ins AS (
		INSERT INTO clothes_colors (cloth_id, color_id) (SELECT $1, unnest($2::int[]))
	)
	SELECT name FROM colors WHERE id = any($2::int[])`
	clothMaterialsUpdate = `
	WITH d AS (
		DELETE FROM clothes_materials WHERE cloth_id = $1
	), ins AS (
		INSERT INTO clothes_materials (cloth_id, material_id) (SELECT $1, unnest($2::int[]))
	)
	SELECT name FROM materials WHERE id = any($2::int[])`

	clothDelete = `DELETE FROM clothes WHERE id = $1`

	selectClothesLimitOffset = `
	SELECT cl.id, cl.name, ct.name, array(SELECT name FROM colors WHERE cc.color_id = id), 
	array(SELECT name FROM materials WHERE cm.material_id = id)
	FROM clothes cl
	INNER JOIN clothes_type ct ON cl.type_id = ct.id
	INNER JOIN clothes_color cc ON cl.id = cc.costume_id
	INNER JOIN clothes_materials cm ON cl.id = cm.costume_id
	LIMIT $1 OFFSET $2`

	selectClothesByIdArray = `SELECT cl.id, cl.name, ct.name, array(SELECT name FROM colors WHERE cc.color_id = id), 
	array(SELECT name FROM materials WHERE cm.material_id = id)
	FROM clothes cl
	INNER JOIN clothes_type ct ON cl.type_id = ct.id
	INNER JOIN clothes_color cc ON cl.id = cc.costume_id
	INNER JOIN clothes_materials cm ON cl.id = cm.costume_id
	WHERE cl.id = any($1::int[])`
)
