package user

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	// service concentra la lógica de negocio y el acceso a datos.
	service *Service
}

// NewHandler construye el handler conectado al service del dominio.
func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

// Create recibe un JSON, lo valida y crea un usuario.
func (handler *Handler) Create(context *gin.Context) {
	var request CreateRequest
	if err := context.ShouldBindJSON(&request); err != nil {
		// Si el body no cumple con la estructura esperada, devolvemos 400.
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Delegamos la creación al service.
	response, err := handler.service.Create(request)
	if err != nil {
		// Un error interno del flujo termina en 500.
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Si salió bien, respondemos con 201 Created.
	context.JSON(http.StatusCreated, response)
}

// FindAll devuelve la lista completa de usuarios.
func (handler *Handler) FindAll(context *gin.Context) {
	// El handler solo traduce la request HTTP; la lógica vive en el service.
	responses, err := handler.service.FindAll()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 200 OK con el listado.
	context.JSON(http.StatusOK, responses)
}

// FindByID busca un usuario usando el parámetro id de la URL.
func (handler *Handler) FindByID(context *gin.Context) {
	// context.Param("id") toma el valor desde una ruta como /users/3.
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		// Si el id no es un número válido, la request está mal formada.
		context.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	// El service resuelve la búsqueda.
	response, err := handler.service.FindByID(id)
	if err != nil {
		// NotFound se traduce a 404 para que el cliente entienda que no existe.
		if err == ErrNotFound {
			context.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		// Cualquier otro fallo inesperado se trata como error interno.
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Si existe, devolvemos 200 con el usuario.
	context.JSON(http.StatusOK, response)
}

// Update reemplaza los datos principales de un usuario existente.
func (handler *Handler) Update(context *gin.Context) {
	// Primero leemos el id desde la URL.
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	// Después validamos el JSON con los datos nuevos.
	var request UpdateRequest
	if err := context.ShouldBindJSON(&request); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// El service aplica la regla de negocio y persiste el cambio.
	response, err := handler.service.Update(id, request)
	if err != nil {
		// Si el usuario no existe, respondemos con 404.
		if err == ErrNotFound {
			context.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		// Cualquier otro error se considera interno.
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 200 OK con el recurso actualizado.
	context.JSON(http.StatusOK, response)
}

// Delete elimina un usuario por id.
func (handler *Handler) Delete(context *gin.Context) {
	// Convertimos el parámetro de ruta a número.
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	// La eliminación real ocurre en el service.
	if err := handler.service.Delete(id); err != nil {
		// Si el usuario no existe, devolvemos 404.
		if err == ErrNotFound {
			context.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		// Si falla por otra razón, devolvemos 500.
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 204 significa que la operación se hizo bien, pero no hay cuerpo para devolver.
	context.Status(http.StatusNoContent)
}
