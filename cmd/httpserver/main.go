package main

import (
	"database/sql"

	"github.com/afikrim/movie-api-go/internal/core/services/genre_service"
	"github.com/afikrim/movie-api-go/internal/handlers/genre_handler"
	"github.com/afikrim/movie-api-go/internal/repositories/genre_repository"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root@tcp(127.0.0.1:3306)/movie_api_go")
	if err != nil {
		panic(err.Error())
	}

	genreRepository := genre_repository.NewGenreRepository(db)
	genreService := genre_service.New(genreRepository)
	genreHandler := genre_handler.NewHTTPHandler(genreService)

	router := gin.New()
	router.POST("/genres", genreHandler.Create)
	router.GET("/genres", genreHandler.FindAll)
	router.GET("/genres/:id", genreHandler.FindOne)
	router.PATCH("/genres/:id", genreHandler.Update)
	router.DELETE("/genres/:id", genreHandler.Remove)

	router.Run(":8000")
}
