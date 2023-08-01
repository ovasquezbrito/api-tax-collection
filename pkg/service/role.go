package service

import (
	"context"

	"github.com/ovasquezbrito/tax-collection/pkg/entity"
	"github.com/ovasquezbrito/tax-collection/pkg/models"
	"github.com/ovasquezbrito/tax-collection/pkg/repository"
)

type RolesService struct {
	repo repository.RoleRepository
}

func NewRoleService(repo repository.RoleRepository) *RolesService {
	return &RolesService{repo: repo}
}

func (s *RolesService) GetAll(ctx context.Context, query models.QueryParameter) ([]models.Role, int, error) {
	queryUpper := query.UpperCase()
	q := &entity.QueryParameter{
		Page:   queryUpper.Page,
		Limit:  queryUpper.Limit,
		Search: queryUpper.Search,
	}
	rr, total, err := s.repo.GetAll(ctx, *q)
	if err != nil {
		return nil, 0, err

	}

	roles := []models.Role{}

	for _, r := range rr {
		roles = append(roles, models.Role{
			IdRole:    r.Id,
			RoleName:  r.RoleName,
			CreatedAt: r.CreatedAt,
			UpdatedAt: r.UpdatedAt,
			Status:    r.Status,
		})
	}

	return roles, total, nil
}

func (s *RolesService) GetById(ctx context.Context, idRol int) (*models.Role, error) {
	r, err := s.repo.GetById(ctx, idRol)
	if err != nil {
		return nil, err
	}
	return &models.Role{
		IdRole:    r.Id,
		RoleName:  r.RoleName,
		CreatedAt: r.CreatedAt,
		UpdatedAt: r.UpdatedAt,
		Status:    r.Status,
	}, nil
}
