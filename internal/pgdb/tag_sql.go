package pgdb

import "context"

type tagReturn struct {
	Id   int
	Name string
}

func (q *Queries) insertTags(ctx context.Context, names []string, table string) ([]tagReturn, error) {
	var request string

	switch table {
	case "tags":

	case "types":
		request = clothesTypesInsert
	}

	rows, err := q.db.Query(ctx, request, names)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var r []tagReturn

	for i := 0; rows.Next(); i++ {
		var id int

		if err := rows.Scan(&id); err != nil {
			return nil, err
		}

		r = append(r, tagReturn{
			Id:   id,
			Name: names[i],
		})
	}

	return r, nil
}

func (q *Queries) updateTag(ctx context.Context, id int, name, table string) error {
	var request string

	switch table {
	case "tags":

	case "types":
		request = clothesTypesInsert
	}

	if _, err := q.db.Exec(ctx, request, id, name); err != nil {
		return err
	}

	return nil
}

func (q *Queries) selectTags(ctx context.Context, table string) ([]tagReturn, error) {
	var request string

	switch table {
	case "tags":

	case "types":
		request = clothesTypesInsert
	}

	rows, err := q.db.Query(ctx, request)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var r []tagReturn

	for rows.Next() {
		var t tagReturn

		if err := rows.Scan(&t.Id, &t.Name); err != nil {
			return nil, err
		}

		r = append(r, t)
	}

	return r, nil
}

func (q *Queries) deleteTag(ctx context.Context, id int, table string) error {
	var request string

	switch table {
	case "tags":

	case "types":
		request = clothesTypesInsert
	}

	if _, err := q.db.Exec(ctx, request, id); err != nil {
		return err
	}

	return nil
}
