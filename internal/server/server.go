package server

import (
	"assignment/internal/controller"
	"assignment/internal/repository"
	"assignment/internal/routes"
	"assignment/internal/service"
	"assignment/pkg/utils"
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

type Server struct {
	conn *sqlx.DB
	gin  *gin.Engine
}

func NewServer(conn *sqlx.DB) Server {
	r := gin.New()

	// Explicit middleware (NO redirect logic)
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// 🔴 Disable Gin redirects (critical)
	r.RedirectTrailingSlash = false
	r.RedirectFixedPath = false
	r.RemoveExtraSlash = false
	return Server{gin: r, conn: conn}
}

func (s *Server) Run() {

	validator := utils.NewValidator()

	dogRepo := repository.NewRepository(s.conn)
	dogService := service.NewService(dogRepo, validator)
	dogController := controller.NewController(dogService)

	quit := make(chan os.Signal, 1)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	server := &http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: s.gin,
	}

	routes.DogRoutes(s.gin.Group("/api/"), dogController)

	s.gin.GET("/", func(c *gin.Context) {
		c.File("./app/index.html")
	})

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Failed to initialize http server: %v\n", err)
		}
	}()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	log.Println("Http Server Listening on port 8080")

	<-quit
	log.Println("Shutting down http server...")
	server.Shutdown(ctx)

}
