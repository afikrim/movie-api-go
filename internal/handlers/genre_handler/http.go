package genre_handler

import (
	"strconv"

	"github.com/afikrim/movie-api-go/internal/core/ports"
	"github.com/gin-gonic/gin"
)

type HTTPHandler struct {
	genreService ports.GenreService
}

func NewHTTPHandler(genreService ports.GenreService) *HTTPHandler {
	return &HTTPHandler{
		genreService: genreService,
	}
}

func (handler *HTTPHandler) Create(ctx *gin.Context) {
	createGenreDto := CreateGenreDto{}
	ctx.BindJSON(&createGenreDto)

	genre, err := handler.genreService.Create(createGenreDto.Name)
	if err != nil {
		ctx.AbortWithStatusJSON(500, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(200, genre)
}

func (handler *HTTPHandler) FindAll(ctx *gin.Context) {
	genres, err := handler.genreService.FindAll()
	if err != nil {
		ctx.AbortWithStatusJSON(500, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(200, genres)
}

func (handler *HTTPHandler) FindOne(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.AbortWithStatusJSON(500, gin.H{"message": err.Error()})
		return
	}

	genre, err := handler.genreService.FindOne(id)
	if err != nil {
		ctx.AbortWithStatusJSON(500, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(200, genre)
}

func (handler *HTTPHandler) Update(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.AbortWithStatusJSON(500, gin.H{"message": err.Error()})
		return
	}

	updateGenreDto := UpdateGenreDto{}
	ctx.BindJSON(&updateGenreDto)

	genre, err := handler.genreService.Update(id, updateGenreDto.Name)
	if err != nil {
		ctx.AbortWithStatusJSON(500, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(200, genre)
}

func (handler *HTTPHandler) Remove(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.AbortWithStatusJSON(500, gin.H{"message": err.Error()})
		return
	}

	genre, err := handler.genreService.Remove(id)
	if err != nil {
		ctx.AbortWithStatusJSON(500, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(200, genre)
}
