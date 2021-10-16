package pgdb

const (
	insertCostumePerformances = `
	INSERT INTO costumes_performances (costume_id, performace_id) (SELECT $1, unnest($2::int[]))
	RETURNING (SELECT name FROM performaces WHERE id = any($2::int[]))`

	insertCostume = `
	INSERT INTO costumes (name, description, is_decor, condition, size, created_at) 
	VALUES ($1, $2, $3, $4, $5, now(), $6) 
	RETURNING id`
	updateCostume = `
	UPDATE costumes SET name = $2, description = $3, is_decor = $4, condition = $5, size = $6, per WHERE id = $1`
	selectCostumeById = `
	SELECT c.name, c.description, c.is_decor, c.condition, c.size, coalsce(c.perfomance_id, 0), 
	array_agg(cc.name), array_agg(cm.name), array_agg(cd.name), array_agg(pc.name)
	FROM costumes c
	INNER JOIN costumes_colors cc ON c.id = cc.costume_id
	INNER JOIN costumes_materials cm ON c.id = cc.costume_id
	INNER JOIN costumes_designers cd ON c.id = cc.costume_id
	INNER JOIN performances_costumes pc ON c.id = pc.costume_id
	WHERE c.id = $1`

	insertCostumeColors = `
	INSERT INTO costumes_colors (costume_id, color_id) (SELECT $1, unnest($2::int[])) 
	RETURNING (SELECT name FROM colors WHERE id = color_id)`
	updateCostumeColors = `
	WITH d AS (DELETE FROM costumes_colors WHERE costume_id = $1)
	INSERT INTO costumes_colors (costume_id, color_id) (SELECT $1, unnest($2::int[]))
	RETURNING (SELECT name FROM colors WHERE id = color_id)`

	insertColorNames = `INSERT INTO colors (name) (SELECT unnest($2::varchar[]))`
	updateColorName  = `UPDATE colors SET name = $2 WHERE id = $1`

	insertCostumeDesigners = `
	INSERT INTO costumes_designers (costume_id, designer_id) (SELECT $1, unnest($2::int[]))
	RETURNING (SELECT name FROM designers WHERE id = any($2::int[]))`
	updateCostumeDesigners = `
	WITH d AS (DELETE FROM costumes_designers WHERE costume_id = $1)
	INSERT INTO costumes_designers (costume_id, designer_id) (SELECT $1, unnest($2::int[]))
	RETURNING (SELECT name FROM designers WHERE id = any($2::int[]))`

	insertDesignerName = `INSERT INTO designers (name) (SELECT unnest($2::varchar[]))`
	updateDesignerName = `UPDATE designers SET name = $2 WHERE id = $1`

	insertCostumeMaterials = `
	INSERT INTO costumes_materials (costume_id, material_id) (SELECT $1, unnest($2::int[]))
	RETURNING (SELECT name FROM materials WHERE id = any($2::int[]))`
	updateCostumeMaterials = `
	WITH d AS (DELETE FROM costumes_materials WHERE costume_id = $1)
	INSERT INTO costumes_materials (costume_id, material_id) (SELECT $1, unnest($2::int[]))
	RETURNING (SELECT name FROM materials WHERE id = any($2::int[]))`

	insertMaterialNames = `INSERT INTO materials (name) (SELECT unnest($2::varchar[]))`
	updateMaterialName  = `UPDATE materials SET name = $2 WHERE id = $1`
)
