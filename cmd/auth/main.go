package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"ride-microservices-go/config"

	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
   
    cfg := config.Load()

    dsn := "postgres://" + cfg.DBUser + ":" + cfg.DBPassword + "@" + cfg.DBHost + ":" + cfg.DBPort + "/" + cfg.DBName
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    dbpool, err := pgxpool.New(ctx, dsn)
    if err != nil {
        log.Fatalf("Unable to connect to DB: %v", err)
    }
    defer dbpool.Close()

    if err := dbpool.Ping(ctx); err != nil {
        log.Fatalf("DB ping failed: %v", err)
    }
    log.Println("✅ DB connection successful")

   
    srv := &http.Server{
        Addr:         ":" + cfg.Port,
        ReadTimeout:  10 * time.Second,
        WriteTimeout: 10 * time.Second,
        IdleTimeout:  120 * time.Second,
    }

    log.Printf("🚀 Auth service running on port %s\n", cfg.Port)
    log.Fatal(srv.ListenAndServe())
}
