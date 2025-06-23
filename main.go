package main

import (
	initializers "GoginOrmDocker/Initializers"
	"GoginOrmDocker/controllers"
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()

	// Try to connect to database, but don't fail if it's not available
	var err error

	// Check environment and use appropriate connection method
	switch {
	case os.Getenv("DIRECT_DB") == "true":
		// Use direct connection - no file system operations at all
		err = initializers.ConnectToDBDirect()
	case os.Getenv("MINIMAL_DB") == "true":
		// Use minimal configuration - minimal file system operations
		err = initializers.ConnectToDBMinimal()
	case os.Getenv("CONTAINER_ENV") == "true":
		// Use container defaults with minimal file system ops
		err = initializers.ConnectToDBContainer()
	default:
		// Use standard connection with fallback defaults
		err = initializers.ConnectToDB()
	}

	if err != nil {
		log.Printf("Warning: Database connection failed: %v", err)
		log.Println("Application will start without database connection")
	}
}

func main() {
	r := gin.Default()
	r.POST("/user", controllers.PostCreate)
	r.GET("/user", controllers.PostIndex)
	r.GET("/user/:id", controllers.GetSpecific)
	r.DELETE("/user/:id", controllers.Delete)
	r.PUT("/user/:id", controllers.UpdateUser)

	// Use custom HTTP server instead of gin.Engine.Run() to avoid file system operations
	if os.Getenv("MINIMAL_SERVER") == "true" {
		startServerMinimal(r)
	} else {
		startServer(r)
	}
}

func startServer(router *gin.Engine) {
	// Get port from environment or use default
	port := getEnvWithDefault("PORT", "8080")
	addr := ":" + port

	// Create custom HTTP server
	srv := &http.Server{
		Addr:         addr,
		Handler:      router,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	// Start server in a goroutine
	go func() {
		log.Printf("Starting server on port %s", port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Printf("Server error: %v", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down server...")

	// Create a deadline for server shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Attempt graceful shutdown
	if err := srv.Shutdown(ctx); err != nil {
		log.Printf("Server forced to shutdown: %v", err)
	}

	log.Println("Server exited")
}

func startServerMinimal(router *gin.Engine) {
	// Use hardcoded port to avoid environment variable access
	addr := ":8080"

	// Create minimal HTTP server - no file system operations
	srv := &http.Server{
		Addr:    addr,
		Handler: router,
	}

	// Start server in a goroutine
	go func() {
		log.Printf("Starting minimal server on port 8080")
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Printf("Server error: %v", err)
		}
	}()

	// Wait for interrupt signal
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down minimal server...")

	// Graceful shutdown with minimal timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Printf("Server forced to shutdown: %v", err)
	}

	log.Println("Minimal server exited")
}

func getEnvWithDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
