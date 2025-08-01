package main

import (
	"context"
	"coupon/cmd/middleware"
	"coupon/internal/data"
	"database/sql"
	"flag"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

type config struct {
	port       int
	jwt_secret string
	db         struct {
		dsn          string
		maxOpenConns int
		maxIdleConns int
		maxIdleTime  string
	}
}

type application struct {
	config config
	models data.Models
}

func openDB(cfg config) (*sql.DB, error) {
	db, err := sql.Open("postgres", cfg.db.dsn)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func main() {

	var cfg config
	flag.IntVar(&cfg.port, "port", 8000, "API server port")
	flag.StringVar(&cfg.db.dsn, "db-dsn", "postgres://coupon:coupon@db:5432/coupon_db?sslmode=disable", "PostgreSQL DSN")
	flag.StringVar(&cfg.jwt_secret, "jwt-secret", "thisisasamplejwtsecret", "Secret for signing jwt secret")
	flag.Parse()

	db, err := openDB(cfg)
	if err != nil {
		fmt.Println(cfg.db.dsn)
		log.Fatal(err)
	}

	app := &application{
		config: cfg,
		models: data.NewModels(db),
	}

	router := gin.Default()

	router.Use(middleware.RateLimiter())
	router.Use(middleware.CORSMiddleware())

	app.register_routes(router)

	err = router.Run(fmt.Sprintf("0.0.0.0:%s", strconv.Itoa(app.config.port)))
	log.Fatal(err)
}
