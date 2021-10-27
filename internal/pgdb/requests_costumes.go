package pgdb

const (
	costumeInsert = `INSERT INTO costumes (
		name, description, image_front, image_back, image_sideway, image_details
	) VALUES ($1, $2, $3, $4, $5, $6)
	RETURNING id`
	costumeClothesInsert = `INSERT INTO costumes_clothes (costume_id, cloth_id) (SELECT $1, unnest($2::int[]))`

	costumeUpdate = `
	UPDATE costumes SET name = $2, desciption = $3, image_front = $7, image_back = $8, image_sideway = $9, image_details = $10 WHERE id = $1`
	costumeClothesUpdate = `
	WITH d AS (DELETE FROM costumes_clothes WHERE costume_id = $1)
	INSERT INTO costumes_clothes (costume_id, cloth_id) (SELECT $1, unnest($2::int[]))`

	costumeDelete = `DELETE FROM costumes WHERE id = $1`

	costumeSelectWithLimitOffset = `
	SELECT c.id, c.name, coalesce(c.description, ''), coalesce(c.image_front, ''), 
	coalesce(c.image_back, ''), coalesce(c.image_sideway, ''), coalesce(c.image_details, '')
	FROM costumes c
	ORDER BY c.id desc
	LIMIT $1 OFFSET $2`
	costumeSelectByPerformanceId = `
	SELECT c.id, c.name, coalesce(c.description, ''), coalesce(c.image_front, ''), 
	coalesce(c.image_back, ''), coalesce(c.image_sideway, ''), coalesce(c.image_details, '')
	FROM costumes c
	INNER JOIN performance_costumes pc ON c.id = pc.costume_id
	WHERE pc.performance_id = $1
	GROUP BY c.id`

	costumeWriteOff = `UPDATE costumes SET is_archived = true WHERE id = any($1::int[])`
)
