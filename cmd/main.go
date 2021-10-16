package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/grum261/theater/internal/models/service"
	"github.com/grum261/theater/internal/pgdb"
	"github.com/grum261/theater/internal/rest"
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

	c := pgdb.NewCostume(db)
	svc := service.NewCostume(c)
	h := rest.NewCostumeHandler(svc)

	app := fiber.New(fiber.Config{
		ReadTimeout:  time.Second,
		WriteTimeout: time.Second,
		IdleTimeout:  time.Second,
	})

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
