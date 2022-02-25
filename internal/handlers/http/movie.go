package http

import (
	"context"
	"net/http"
	"strconv"

	"github.com/afikrim/movie-api-go/internal/core/domain"
	"github.com/afikrim/movie-api-go/internal/core/ports"
	"github.com/labstack/echo/v4"
)

type movieHttpHandler struct {
	movieService ports.MovieService
}

func NewMovieHttpHandler(movieService ports.MovieService) *movieHttpHandler {
	return &movieHttpHandler{
		movieService: movieService,
	}
}

func (handler *movieHttpHandler) Create(e echo.Context) error {
	ctx := context.Background()

	var createMovieDto domain.CreateMovieDto
	if err := e.Bind(&createMovieDto); err != nil {
		return e.JSON(http.StatusBadRequest, &Response{Message: err.Error()})
	}

	movie, err := handler.movieService.Create(ctx, createMovieDto)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, &Response{Message: err.Error()})
	}

	return e.JSON(http.StatusOK, &Response{Message: "Successfully insert new movie", Data: map[string]interface{}{"movie": movie}})
}

func (handler *movieHttpHandler) FindAll(e echo.Context) error {
	ctx := context.Background()

	movies, err := handler.movieService.FindAll(ctx)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, &Response{Message: err.Error()})
	}

	return e.JSON(http.StatusOK, &Response{Message: "Success get all", Data: map[string]interface{}{"movies": movies}})
}

func (handler *movieHttpHandler) FindOne(e echo.Context) error {
	ctx := context.Background()

	id, err := strconv.ParseInt(e.Param("id"), 10, 64)
	if err != nil {
		return e.JSON(http.StatusBadRequest, &Response{Message: err.Error()})
	}

	movie, err := handler.movieService.FindOne(ctx, id)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, &Response{Message: err.Error()})
	}

	return e.JSON(http.StatusOK, &Response{Message: "Successfully get one", Data: map[string]interface{}{"movie": movie}})
}

func (handler *movieHttpHandler) Update(e echo.Context) error {
	ctx := context.Background()

	id, err := strconv.ParseInt(e.Param("id"), 10, 64)
	if err != nil {
		return e.JSON(http.StatusBadRequest, &Response{Message: err.Error()})
	}

	var updateMovieDto domain.UpdateMovieDto
	if err := e.Bind(updateMovieDto); err != nil {
		return e.JSON(http.StatusBadRequest, &Response{Message: err.Error()})
	}

	movie, err := handler.movieService.Update(ctx, id, updateMovieDto)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, &Response{Message: err.Error()})
	}

	return e.JSON(http.StatusOK, &Response{Message: "Successfully update one", Data: map[string]interface{}{"movie": movie}})
}

func (handler *movieHttpHandler) Remove(e echo.Context) error {
	ctx := context.Background()

	id, err := strconv.ParseInt(e.Param("id"), 10, 64)
	if err != nil {
		return e.JSON(http.StatusBadRequest, &Response{Message: err.Error()})
	}

	movie, err := handler.movieService.Remove(ctx, id)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, &Response{Message: err.Error()})
	}

	return e.JSON(http.StatusOK, &Response{Message: "Successfully remove one", Data: map[string]interface{}{"movie": movie}})
}
