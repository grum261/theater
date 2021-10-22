package pgdb

const (
	costumeInsert = `INSERT INTO costumes (
		name, description, designer, condition, is_decor, location, is_archived, size, 
		image_front, image_back, image_sideway, image_details
	) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)
	RETURNING id`
	costumeClothesInsert = `INSERT INTO costumes_clothes (costume_id, cloth_id) (SELECT $1, unnest($2::int[]))`

	costumeUpdate = `
	UPDATE costumes SET name = $2, desciption = $3, condition = $4, is_decor = $5, location = $6, is_archived = $7, size = $8, 
	image_front = $9, image_back = $10, image_sideway = $11, image_details = $12, designer = $13 WHERE id = $1`
	costumeClothesUpdate = `
	WITH d AS (DELETE FROM costumes_clothes WHERE costume_id = $1)
	INSERT INTO costumes_clothes (costume_id, cloth_id) (SELECT $1, unnest($2::int[]))`

	costumeDelete = `DELETE FROM costumes WHERE id = $1`

	costumeSelectWithLimitOffset = `
	SELECT c.id, c.name, c.size, c.condition, c.location, c.description, c.image_front, 
	c.image_back, c.image_sideway, c.image_details, c.is_decor, c.is_archived, c.designer, array_agg(ccl.costume_id)
	FROM costumes c
	INNER JOIN costumes_clothes ccl ON c.id = ccl.costume_id
	GROUP BY c.id
	ORDER BY c.id desc
	LIMIT $1 OFFSET $2`

	costumeWriteOff = `UPDATE costumes SET is_archived = true WHERE id = $1`
)
