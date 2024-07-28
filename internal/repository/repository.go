package repository

import (
	"errors"
	"go_test/internal/domain"
)

type Repository struct {
	entities map[int]*domain.Entity
}

func NewRepository() *Repository {
	return &Repository{
		entities: make(map[int]*domain.Entity),
	}
}

func (r *Repository) GetAllEntities() ([]*domain.Entity, error) {
	var entities []*domain.Entity
	for _, entity := range r.entities {
		entities = append(entities, entity)
	}
	return entities, nil
}

func (r *Repository) GetEntityById(id int) (*domain.Entity, error) {
	entity, exists := r.entities[id]
	if !exists {
		return nil, errors.New("entity not found")
	}
	return entity, nil
}

func (r *Repository) CreateEntity(entity *domain.Entity) {
	r.entities[entity.ID] = entity
}

func (r *Repository) UpdateEntity(id int, entity *domain.Entity) (*domain.Entity, error) {
	_, exists := r.entities[id]
	if !exists {
		return nil, errors.New("entity not found")
	}
	r.entities[id] = entity
	return entity, nil
}
