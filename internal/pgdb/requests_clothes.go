package pgdb

const (
	clothInsert = `
	INSERT INTO clothes (name, type_id, location, designer, condition, size, is_decor, is_archived) VALUES ($1, $2, $3, $4, $5, $6, $7, $8) 
	RETURNING id, (SELECT name FROM clothes_types WHERE id = $2)`
	clothColorsInsert = `
	WITH ins AS (INSERT INTO clothes_colors (cloth_id, color_id) (SELECT $1, unnest($2::int[])))
	SELECT name FROM colors WHERE id = any($2::int[])`
	clothMaterialsInsert = `
	WITH ins AS (INSERT INTO clothes_materials (cloth_id, material_id) (SELECT $1, unnest($2::int[])))
	SELECT name FROM materials WHERE id = any($2::int[])`

	clothUpdate = `
	UPDATE clothes SET name = $2, type_id = $3, location = $4, designer = $5, condition = $6, size = $7, is_decor = $8, is_archived = $9 WHERE id = $1 
	RETURNING (SELECT name FROM clothes_types WHERE id = $3)`
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
	SELECT cl.id, cl.name, ct.name, cl.location, cl.designer, cl.condition, cl.size, 
	array_agg(c.name), array_agg(m.name), cl.is_decor, cl.is_archived
	FROM clothes cl
	INNER JOIN clothes_types ct ON cl.type_id = ct.id
	INNER JOIN clothes_colors cc ON cl.id = cc.cloth_id
	INNER JOIN colors c ON cc.color_id = c.id
	INNER JOIN clothes_materials cm ON cl.id = cm.cloth_id
	INNER JOIN materials m ON cm.material_id = m.id
	GROUP BY cl.id, ct.name
	ORDER BY id desc
	LIMIT $1 OFFSET $2`
	selectClothesByCostumeId = `
	SELECT cl.id, cl.name, ct.name, cl.location, cl.designer, cl.condition, cl.size, 
	array_agg(c.name), array_agg(m.name), cl.is_decor, cl.is_archived
	FROM clothes cl
	INNER JOIN clothes_types ct ON cl.type_id = ct.id
	INNER JOIN clothes_colors cc ON cl.id = cc.cloth_id
	INNER JOIN colors c ON cc.color_id = c.id
	INNER JOIN clothes_materials cm ON cl.id = cm.cloth_id
	INNER JOIN materials m ON cm.material_id = m.id
	INNER JOIN costumes_clothes ccl ON cl.id = ccl.cloth_id
	WHERE ccl.costume_id = $1
	GROUP BY cl.id, ct.name`
	selectClothesById = `
	SELECT cl.id, cl.name, ct.name, cl.location, cl.designer, cl.condition, cl.size, 
	array_agg(c.name), array_agg(m.name), , cl.is_decor, cl.is_archived
	FROM clothes cl
	INNER JOIN clothes_types ct ON cl.type_id = ct.id
	INNER JOIN clothes_colors cc ON cl.id = cc.cloth_id
	INNER JOIN colors c ON cc.color_id = c.id
	INNER JOIN clothes_materials cm ON cl.id = cm.cloth_id
	INNER JOIN materials m ON cm.material_id = m.id
	WHERE cl.id = $1
	GROUP BY cl.id, ct.name`
)
