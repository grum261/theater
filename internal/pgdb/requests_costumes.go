package pgdb

const (
	costumeInsert = `INSERT INTO costumes (
		name, description, image_front, image_back, image_sideway, image_details, designer, 
		location, size, condition, is_decor, is_archived, comment
	) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)
	RETURNING id`
	costumeClothesInsert = `INSERT INTO costumes_clothes (costume_id, cloth_id) (SELECT $1, unnest($2::int[]))`

	costumeUpdate = `
	UPDATE costumes SET name = $2, desciption = $3, image_front = $4, image_back = $5, image_sideway = $6, image_details = $7,
	designer = $8, location = $9, $size = $10, condition = $11, is_decor = $12, is_archived = $13, comment = $14
	WHERE id = $1`
	costumeClothesUpdate = `
	WITH d AS (DELETE FROM costumes_clothes WHERE costume_id = $1)
	INSERT INTO costumes_clothes (costume_id, cloth_id) (SELECT $1, unnest($2::int[]))`

	costumeDelete = `DELETE FROM costumes WHERE id = $1`

	costumeSelectWithLimitOffset = `
	SELECT c.id, c.name, coalesce(c.description, ''), coalesce(c.image_front, ''), 
	coalesce(c.image_back, ''), coalesce(c.image_sideway, ''), coalesce(c.image_details, ''), c.designer, 
	coalesce(c.location, ''), c.size, c.condition, c.is_decor, c.is_archived, coalesce(c.comment, '')
	FROM costumes c
	ORDER BY c.id desc
	LIMIT $1 OFFSET $2`
	costumeSelectByPerformanceId = `
	SELECT c.id, c.name, coalesce(c.description, ''), coalesce(c.image_front, ''), 
	coalesce(c.image_back, ''), coalesce(c.image_sideway, ''), coalesce(c.image_details, ''), c.designer, 
	coalesce(c.location, ''), c.size, c.condition, c.is_decor, c.is_archived, coalesce(c.comment, '')
	FROM costumes c
	INNER JOIN performance_costumes pc ON c.id = pc.costume_id
	WHERE pc.performance_id = $1`

	costumeWriteOff = `UPDATE costumes SET is_archived = true WHERE id = any($1::int[])`
)
