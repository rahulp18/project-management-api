package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"

	"project-management/internal/auth"
	"project-management/internal/config"
	"project-management/internal/database"
	"project-management/internal/user"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	cfg := config.Load()
	log.Println("DATABASE_URL from config:", cfg.DBUrl)
	db, err := database.New(cfg.DBUrl)
	if err != nil {
		log.Fatalf("Failed to connect to DB: %v", err)
	}

	//  Initialize Router
	r := chi.NewRouter()
	// middleware
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})
	userRepo := user.NewRepository(db.Pool)
	authService := auth.NewService(userRepo)
	authHandler := auth.NewHandler(authService)
	auth.AuthRoutes(r, authHandler)
	// USER ROUTES
	userService := user.NewService(userRepo)
	userHandler := user.NewHandler(userService)
	user.UserRoutes(r, userHandler)
	server := &http.Server{
		Addr:         ":" + cfg.Port,
		Handler:      r,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  60 * time.Second,
	}
	go func() {
		log.Printf("Server running on port %s\n", cfg.Port)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server failed %v", err)
		}
	}()
	//   Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit
	db.Close()
	log.Println("Shutting down server...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("Server exited properly")
}
