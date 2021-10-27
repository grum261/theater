package pgdb

const (
	costumeInsert = `INSERT INTO costumes (
		name, description, is_decor, is_archived, 
		image_front, image_back, image_sideway, image_details
	) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
	RETURNING id`
	costumeClothesInsert = `INSERT INTO costumes_clothes (costume_id, cloth_id) (SELECT $1, unnest($2::int[]))`

	costumeUpdate = `
	UPDATE costumes SET name = $2, desciption = $3, is_decor = $5, is_archived = $6, 
	image_front = $7, image_back = $8, image_sideway = $9, image_details = $10 WHERE id = $1`
	costumeClothesUpdate = `
	WITH d AS (DELETE FROM costumes_clothes WHERE costume_id = $1)
	INSERT INTO costumes_clothes (costume_id, cloth_id) (SELECT $1, unnest($2::int[]))`

	costumeDelete = `DELETE FROM costumes WHERE id = $1`

	costumeSelectWithLimitOffset = `
	SELECT c.id, c.name, c.description, c.image_front, 
	c.image_back, c.image_sideway, c.image_details, c.is_decor, c.is_archived, array_agg(ccl.clothes_id)
	FROM costumes c
	INNER JOIN costumes_clothes ccl ON c.id = ccl.costume_id
	GROUP BY c.id
	ORDER BY c.id desc
	LIMIT $1 OFFSET $2`

	costumeWriteOff = `UPDATE costumes SET is_archived = true WHERE id = any($1::int[])`
)
