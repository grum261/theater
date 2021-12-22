package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/grum261/theater/internal/pgdb"
	"github.com/grum261/theater/internal/rest"
	"github.com/grum261/theater/internal/service"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/joho/godotenv"
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/config"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
)

type zapTracerLogger struct {
	*zap.Logger
}

func (z *zapTracerLogger) Infof(msg string, args ...interface{}) {
	z.Logger.Info(fmt.Sprintf(msg, args...))
}

func (z *zapTracerLogger) Error(msg string) {
	z.Logger.Error(msg)
}

type serverConfig struct {
	Middlewares []func(*fiber.Ctx) error
	DB          *pgxpool.Pool
}

func main() {
	var envPath, port string

	flag.StringVar(&envPath, "env", "", "")
	flag.StringVar(&port, "port", "", "")
	flag.Parse()

	if err := godotenv.Load(envPath); err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()

	if err := run(ctx, port); err != nil {
		log.Fatal(err)
	}
}

func run(ctx context.Context, port string) error {
	logger, err := zap.NewProduction()
	if err != nil {
		return err
	}

	loggerMiddleware := func(c *fiber.Ctx) error {
		if err := c.Next(); err != nil {
			return err
		}

		statusCode := c.Response().StatusCode()

		fields := []zap.Field{
			zap.Time("time", time.Now()),
			zap.String("method", c.Method()),
			zap.String("uri", c.OriginalURL()),
			zap.Int("statusCode", statusCode),
		}

		switch {
		case statusCode >= 500, statusCode >= 400:
			errResponse := struct {
				Err string `json:"error"`
			}{}

			if err := c.BodyParser(&errResponse); err != nil {
				return err
			}

			fields = append(fields, zap.Error(errors.New(errResponse.Err)))

			if statusCode >= 500 {
				logger.Error("Ошибка сервера", fields...)
			} else {
				logger.Warn("Ошибка клиента", fields...)
			}
		case statusCode >= 300:
			logger.Warn("Редирект", fields...)
		default:
			logger.Info("Запрос выполнен успешно", fields...)
		}

		return nil
	}

	pool, err := newDB(ctx)
	if err != nil {
		return err
	}

	app := newApp(serverConfig{
		Middlewares: []func(*fiber.Ctx) error{loggerMiddleware},
		DB:          pool,
	})

	closer, err := newJaegerTracer(&zapTracerLogger{logger})
	if err != nil {
		return err
	}

	ctxSignal, stop := signal.NotifyContext(ctx, os.Interrupt, os.Kill, syscall.SIGTERM, syscall.SIGKILL)

	g, ctxGroup := errgroup.WithContext(ctxSignal)

	g.Go(func() error {
		<-ctxGroup.Done()

		logger.Info("Получен сигнал остановки сервера")

		ctxTimeout, cancel := context.WithTimeout(ctxGroup, time.Second*5)

		defer func() {
			_ = logger.Sync()

			pool.Close()
			closer.Close()
			stop()
			cancel()
		}()

		<-ctxTimeout.Done()

		if err := app.Shutdown(); err != nil {
			return err
		}

		logger.Info("Сервер успешно остановлен")

		return nil
	})

	g.Go(func() error {
		logger.Info(fmt.Sprintf("Начинаем слушать на %s", port))

		if err := app.Listen(port); err != nil {
			return err
		}

		return nil
	})

	if err := g.Wait(); err != nil {
		return err
	}

	return nil
}

func newApp(s serverConfig) *fiber.App {
	app := fiber.New(fiber.Config{
		ReadTimeout:  time.Second,
		WriteTimeout: time.Second,
		IdleTimeout:  time.Second,
	})

	for _, mw := range s.Middlewares {
		app.Use(mw)
	}

	store := pgdb.NewStore(s.DB)
	svc := service.NewServices(store.Tag, store.Cloth, store.Costume, store.Performance)
	h := rest.NewHandlers(svc.Tag, svc.Cloth, svc.Costume, svc.Performance)

	h.RegisterRoutes(app.Group("/api/v1"))

	app.Static("/", "./swagger-ui")

	return app
}

func newDB(ctx context.Context) (*pgxpool.Pool, error) {
	host := os.Getenv("POSTGRES_HOST")
	port := os.Getenv("POSTGRES_PORT")
	username := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbName := os.Getenv("POSTGRES_DB")

	dbURL := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s", username, password, host, port, dbName)

	config, err := pgxpool.ParseConfig(dbURL)
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

func newJaegerTracer(tracerLogger *zapTracerLogger) (io.Closer, error) {
	tracer, closer, err := config.Configuration{
		ServiceName: "theater",
		Sampler: &config.SamplerConfig{
			Type:  jaeger.SamplerTypeConst,
			Param: 1,
		},
		Reporter: &config.ReporterConfig{
			LogSpans: true,
		},
	}.NewTracer(config.Logger(tracerLogger))
	if err != nil {
		return nil, err
	}

	opentracing.SetGlobalTracer(tracer)

	return closer, nil
}
