package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"

	"scheduler/internal/api"
	"scheduler/internal/db"
	"scheduler/internal/scheduler"
)

func main() {
	dbURL := "postgres://octavio:Scr@ppy1121@localhost:5432/scheduler"
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()

	pool, err := pgxpool.New(ctx, dbURL)
	if err != nil {
		log.Fatalf("failed to connect to DB: %v", err)
	}
	defer pool.Close()

	db.InitDB(pool)

	go func() {
		ticker := time.NewTicker(1 * time.Minute)
		defer ticker.Stop()

		for {
			select {
			case <-ctx.Done():
				log.Println("scheduler shutting down background scheduler...")
				return
			case <-ticker.C:
				log.Println("scheduler checking for scheduled jobs...")
				if err := scheduler.RunScheduledJobs(); err != nil {
					log.Printf("scheduler error: %v\n", err)
				}
			}
		}
	}()

	r := api.NewRouter()
	srv := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	go func() {
		log.Println("api server starting on :8080")
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("failed to start Gin server: %v", err)
		}
	}()

	<-ctx.Done()
	log.Println("interrupt received: shutting down...")

	shutdownCtx, cancelShutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelShutdown()

	if err := srv.Shutdown(shutdownCtx); err != nil {
		log.Fatalf("gin server forced to shutdown: %v", err)
	}

	log.Println("server exited cleanly")
}
