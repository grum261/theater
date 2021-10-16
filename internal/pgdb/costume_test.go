package pgdb_test

// import (
// 	"context"
// 	"fmt"
// 	"testing"

// 	"github.com/grum261/theater/internal/pgdb"

// 	"github.com/jackc/pgx/v4/pgxpool"
// )

// func Test_InsertCostume(t *testing.T) {
// 	ctx := context.Background()

// 	conn, err := newDB(ctx)
// 	if err != nil {
// 		t.Fatalf("ошибка: %v", err)
// 	}
// 	defer conn.Close()

// 	q := pgdb.NewQueries(conn)

// 	c, err := q.InsertCostume(ctx, pgdb.InsertCostumeParams{
// 		Name:        "Новый костюм",
// 		Description: "Первый костюм",
// 		IsDecor:     true,
// 		Condition:   "хорошее",
// 		Colors:      []int{1, 2},
// 		Materials:   []int{1},
// 		Designers:   []int{2},
// 	})
// 	if err != nil {
// 		t.Fatalf("выполнение с ошибкой: %v", err)
// 	}

// 	fmt.Println(c)
// }
