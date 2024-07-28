package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go_test/internal/domain"
	"go_test/internal/service"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{service: service}
}

// GetMethod godoc
// @Summary Get all entities
// @Description get all entities
// @Tags entity
// @Accept  json
// @Produce  json
// @Router /entity [get]
func (h *Handler) GetMethod(c *gin.Context) {
	entities, err := h.service.GetAllEntities()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, entities)
}

// GetMethodById godoc
// @Summary Get an entity by ID
// @Description get entity by ID
// @Tags entity
// @Accept  json
// @Produce  json
// @Param id path int true "Entity ID"
// @Router /entity/{id} [get]
func (h *Handler) GetMethodById(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	entity, err := h.service.GetEntityById(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, entity)
}

// PostMethod godoc
// @Summary Create a new entity
// @Description create entity
// @Tags entity
// @Accept  json
// @Produce  json
// @Router /entity [post]
func (h *Handler) PostMethod(c *gin.Context) {
	var entity domain.Entity
	if err := c.ShouldBindJSON(&entity); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	h.service.CreateEntity(&entity)
	c.JSON(http.StatusCreated, entity)
}

// PutMethod godoc
// @Summary Update an entity by ID
// @Description update entity by ID
// @Tags entity
// @Accept  json
// @Produce  json
// @Param id path int true "Entity ID"
// @Router /entity/{id} [put]swag init -g main.go
func (h *Handler) PutMethod(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var entity domain.Entity
	if err := c.ShouldBindJSON(&entity); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	updatedEntity, err := h.service.UpdateEntity(id, &entity)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, updatedEntity)
}
