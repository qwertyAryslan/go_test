package service

import (
	"go_test/internal/domain"
	"go_test/internal/repository"
)

type Service struct {
	repo *repository.Repository
}

func NewService(repo *repository.Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) GetAllEntities() ([]*domain.Entity, error) {
	return s.repo.GetAllEntities()
}

func (s *Service) GetEntityById(id int) (*domain.Entity, error) {
	return s.repo.GetEntityById(id)
}

func (s *Service) CreateEntity(entity *domain.Entity) {
	s.repo.CreateEntity(entity)
}

func (s *Service) UpdateEntity(id int, entity *domain.Entity) (*domain.Entity, error) {
	return s.repo.UpdateEntity(id, entity)
}
