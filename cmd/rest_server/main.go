package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/grum261/theater/internal/pgdb"
	"github.com/grum261/theater/internal/rest"
	"github.com/grum261/theater/internal/service"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/joho/godotenv"
)

func main() {
	var envPath string

	flag.StringVar(&envPath, "env", "", "")
	flag.Parse()

	if err := godotenv.Load(envPath); err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()

	db, err := newDB(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	s := pgdb.NewStore(db)
	svc := service.NewServices(s.Tag, s.Cloth, s.Costume, s.Performance)
	h := rest.NewHandlers(svc.Tag, svc.Cloth, svc.Costume, svc.Performance)

	app := fiber.New(fiber.Config{
		ReadTimeout:  time.Second,
		WriteTimeout: time.Second,
		IdleTimeout:  time.Second,
	})

	f, err := os.Create("./log/messages.log")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	app.Use(logger.New(logger.Config{
		Output: f,
	}))

	app.Static("/", "./assets/swagger-ui")

	h.RegisterRoutes(app.Group("/api/v1"))

	log.Fatal(app.Listen(":8000"))
}

func newDB(ctx context.Context) (*pgxpool.Pool, error) {
	host := os.Getenv("POSTGRES_HOST")
	port := os.Getenv("POSTGRES_PORT")
	username := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbName := os.Getenv("POSTGRES_DB")

	config, err := pgxpool.ParseConfig(fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s", username, password, host, port, dbName))
	if err != nil {
		return nil, err
	}

	conn, err := pgxpool.ConnectConfig(ctx, config)
	if err != nil {
		return nil, err
	}

	if err := conn.Ping(ctx); err != nil {
		return nil, err
	}

	return conn, nil
}
