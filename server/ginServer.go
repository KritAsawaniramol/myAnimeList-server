package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-contrib/cors"
	limits "github.com/gin-contrib/size"
	"github.com/gin-contrib/timeout"

	"github.com/gin-gonic/gin"
	"github.com/kritAsawaniramol/myAnimeList-server/config"
	"github.com/kritAsawaniramol/myAnimeList-server/module/middleware/middlewareHandler"
	"github.com/kritAsawaniramol/myAnimeList-server/module/middleware/middlewareRepository"
	"github.com/kritAsawaniramol/myAnimeList-server/module/middleware/middlewareUsecase"
	"gorm.io/gorm"
)

type ginServer struct {
	app        *gin.Engine
	db         *gorm.DB
	cfg        *config.Config
	middleware middlewareHandler.MiddlewareHandlerService
}

func newMiddleware(cfg *config.Config, db *gorm.DB) middlewareHandler.MiddlewareHandlerService {
	repo := middlewareRepository.NewMiddlewareRepository(db)
	usecase := middlewareUsecase.NewMiddlewareUsecase(repo)
	return middlewareHandler.NewMiddlewareHandler(cfg, usecase)
}

func NewGinServer(cfg *config.Config, db *gorm.DB) Server {
	return &ginServer{
		app:        gin.New(),
		db:         db,
		cfg:        cfg,
		middleware: newMiddleware(cfg, db),
	}
}

func timeoutMiddleware() gin.HandlerFunc {
	return timeout.New(
		timeout.WithTimeout(30*time.Second),
		timeout.WithHandler(func(c *gin.Context) {
			c.Next()
		}),
		timeout.WithResponse(timeoutErrorResponse),
	)
}

func gracefulShutdown(pctx context.Context, srv *http.Server) {
	log.Printf("Start service...")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server...")
	ctx, cancel := context.WithTimeout(pctx, 10*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	// catching ctx.Done(). timeout of 10 seconds.
	<-ctx.Done()
	log.Println("timeout of 10 seconds.")
}

// Start implements Server.
func (g *ginServer) Start() {

	// Define your allowed origins
	allowedOrigins := []string{
		fmt.Sprintf("http://%s:%d",g.cfg.Client.Host, g.cfg.Client.Port),
	}

	// Configure CORS middleware with multiple allowed origins
	config := cors.Config{
		AllowOrigins:     allowedOrigins,
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Authorization", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}

	// Apply the CORS middleware to the Gin router
	g.app.Use(cors.New(config))
	// request's size limit = 10MB
	g.app.Use(limits.RequestSizeLimiter(10 << 20))

	g.app.Use(timeoutMiddleware())
	g.app.Use(gin.Logger())
	g.app.Use(gin.Recovery())

	g.InitialzieHttpHandler()
	g.authService()
	g.userService()
	g.commentService()
	g.animeListService()
	serverUrl := fmt.Sprintf(":%d", g.cfg.App.Port)
	srv := &http.Server{
		Addr:    serverUrl,
		Handler: g.app,
	}

	ctx := context.Background()
	go gracefulShutdown(ctx, srv)

	fmt.Printf("listen on: %s:%d\n", g.cfg.App.Host, g.cfg.App.Port)
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("listen: %s\n", err)
	}

}

func timeoutErrorResponse(c *gin.Context) {
	c.String(http.StatusRequestTimeout, "timeout")
}

func (g *ginServer) InitialzieHttpHandler() {
	g.app.POST("/", g.healthCheckService)

}
