package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	genre_service "github.com/afikrim/movie-api-go/internal/core/services/genre"
	"github.com/afikrim/movie-api-go/internal/handlers/http"
	genre_repository "github.com/afikrim/movie-api-go/internal/repositories/genre"
	"github.com/labstack/echo/v4"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {
	db, err := NewDatabaseInstance()
	if err != nil {
		panic(err.Error())
	}

	e := echo.New()
	e.Logger.SetLevel(log.LstdFlags)

	genreRepository := genre_repository.NewGenreRepository(db)
	genreService := genre_service.New(genreRepository)
	genreHandler := http.NewGenreHttpHandler(genreService)
	genreRouter := e.Group("/genres")

	genreRouter.POST("", genreHandler.Create)
	genreRouter.GET("", genreHandler.FindAll)
	genreRouter.GET("/:id", genreHandler.FindOne)
	genreRouter.PATCH("/:id", genreHandler.Update)
	genreRouter.DELETE("/:id", genreHandler.Remove)

	go func() {
		if err := e.Start(":8000"); err != nil {
			log.Fatalf("error starting server: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	e.Shutdown(ctx)
}

func NewDatabaseInstance() (*gorm.DB, error) {
	gormConf := &gorm.Config{}
	gormConf.Logger = logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Info,
			Colorful:      true,
		},
	)

	instance, err := gorm.Open(mysql.Open("root@tcp(127.0.0.1:3306)/movie_api_go?parseTime=true"), gormConf)
	if err != nil {
		return nil, err
	}

	return instance, nil
}
