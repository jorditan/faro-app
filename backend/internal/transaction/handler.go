package transaction

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

func (handler *Handler) Create(context *gin.Context) {
	var request CreateRequest
	if err := context.ShouldBindJSON(&request); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response, err := handler.service.Create(request)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, response)
}

func (handler *Handler) FindAll(context *gin.Context) {
	responses, err := handler.service.FindAll()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, responses)
}

func (handler *Handler) FindByID(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	response, err := handler.service.FindByID(id)
	if err != nil {
		if err == ErrNotFound {
			context.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, response)
}

func (handler *Handler) Update(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	var request UpdateRequest
	if err := context.ShouldBindJSON(&request); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response, err := handler.service.Update(id, request)
	if err != nil {
		if err == ErrNotFound {
			context.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, response)
}

func (handler *Handler) Delete(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	if err := handler.service.Delete(id); err != nil {
		if err == ErrNotFound {
			context.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.Status(http.StatusNoContent)
}
