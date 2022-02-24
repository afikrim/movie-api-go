package http

import (
	"context"
	"net/http"
	"strconv"

	"github.com/afikrim/movie-api-go/internal/core/domain"
	"github.com/afikrim/movie-api-go/internal/core/ports"
	"github.com/labstack/echo/v4"
)

type genreHttpHandler struct {
	genreService ports.GenreService
}

func NewGenreHttpHandler(genreService ports.GenreService) *genreHttpHandler {
	return &genreHttpHandler{
		genreService: genreService,
	}
}

func (handler *genreHttpHandler) Create(e echo.Context) error {
	ctx := context.Background()

	var createGenreDto domain.CreateGenreDto
	if err := e.Bind(&createGenreDto); err != nil {
		return e.JSON(http.StatusBadRequest, &Response{Message: err.Error()})
	}

	genre, err := handler.genreService.Create(ctx, createGenreDto)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, &Response{Message: err.Error()})
	}

	return e.JSON(http.StatusOK, &Response{Message: "Successfully insert new genre", Data: map[string]interface{}{"genre": genre}})
}

func (handler *genreHttpHandler) FindAll(e echo.Context) error {
	ctx := context.Background()

	genres, err := handler.genreService.FindAll(ctx)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, &Response{Message: err.Error()})
	}

	return e.JSON(http.StatusOK, &Response{Message: "Success get all", Data: map[string]interface{}{"genres": genres}})
}

func (handler *genreHttpHandler) FindOne(e echo.Context) error {
	ctx := context.Background()

	id, err := strconv.ParseInt(e.Param("id"), 10, 64)
	if err != nil {
		return e.JSON(http.StatusBadRequest, &Response{Message: err.Error()})
	}

	genre, err := handler.genreService.FindOne(ctx, id)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, &Response{Message: err.Error()})
	}

	return e.JSON(http.StatusOK, &Response{Message: "Successfully get one", Data: map[string]interface{}{"genre": genre}})
}

func (handler *genreHttpHandler) Update(e echo.Context) error {
	ctx := context.Background()

	id, err := strconv.ParseInt(e.Param("id"), 10, 64)
	if err != nil {
		return e.JSON(http.StatusBadRequest, &Response{Message: err.Error()})
	}

	var updateGenreDto domain.UpdateGenreDto
	if err := e.Bind(updateGenreDto); err != nil {
		return e.JSON(http.StatusBadRequest, &Response{Message: err.Error()})
	}

	genre, err := handler.genreService.Update(ctx, id, updateGenreDto)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, &Response{Message: err.Error()})
	}

	return e.JSON(http.StatusOK, &Response{Message: "Successfully update one", Data: map[string]interface{}{"genre": genre}})
}

func (handler *genreHttpHandler) Remove(e echo.Context) error {
	ctx := context.Background()

	id, err := strconv.ParseInt(e.Param("id"), 10, 64)
	if err != nil {
		return e.JSON(http.StatusBadRequest, &Response{Message: err.Error()})
	}

	genre, err := handler.genreService.Remove(ctx, id)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, &Response{Message: err.Error()})
	}

	return e.JSON(http.StatusOK, &Response{Message: "Successfully remove one", Data: map[string]interface{}{"genre": genre}})
}
