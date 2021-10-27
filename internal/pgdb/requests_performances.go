package pgdb

const (
	performanceInsert         = `INSERT INTO performances (name, location, starting_at, duration) VALUES ($1, $2, $3, $4) RETURNING id`
	performanceUpdate         = `UPDATE performances SET name = $2, location = $3, starting_at = $4, duration = $5 WHERE id = $1`
	performanceCostumesInsert = `INSERT INTO performances_costumes (performance_id, costume_id) (SELECT $1, unnest($2::int[]))`
	performanceCostumesUpdate = `
	WITH d AS (DELETE FROM performances_costumes WHERE performance_id = $1)
	INSERT INTO performances_costumes (performance_id, costume_id) (SELECT $1, unnest($2::int[]))`
	performanceDelete = `DELETE FROM performances WHERE id = $1`
)
