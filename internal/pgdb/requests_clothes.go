package pgdb

const (
	clothInsert = `
	INSERT INTO clothes (name, type_id) VALUES ($1, $2) 
	RETURNING id, (SELECT name FROM clothes_types WHERE id = $2)`

	clothUpdate = `
	UPDATE clothes SET name = $2, type_id = $3 WHERE id = $1 
	RETURNING (SELECT name FROM clothes_types WHERE id = $3)`

	clothDelete = `DELETE FROM clothes WHERE id = $1`

	selectClothesLimitOffset = `
	SELECT cl.id, cl.name, ct.name FROM clothes cl
	INNER JOIN clothes_types ct ON cl.type_id = ct.id
	ORDER BY id desc
	LIMIT $1 OFFSET $2`
	selectClothesByCostumeId = `
	SELECT cl.id, cl.name, ct.name FROM clothes cl
	INNER JOIN clothes_types ct ON cl.type_id = ct.id
	INNER JOIN costumes_clothes ccl ON cl.id = ccl.cloth_id
	WHERE ccl.costume_id = $1`
	selectClothesById = `
	SELECT cl.id, cl.name, ct.name FROM clothes cl
	INNER JOIN clothes_types ct ON cl.type_id = ct.id
	WHERE cl.id = $1`
)
