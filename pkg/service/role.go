package service

import (
	"context"
	"errors"
	"strings"

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

func (s *RolesService) CreateRole(ctx context.Context, rol models.Role) (int, error) {
	role := strings.ToLower(rol.RoleName)
	role = strings.TrimSpace(role)
	r, _ := s.repo.GetRoleByName(ctx, role)
	if r.Id != 0 {
		return 0, errors.New("nombre de rol existe")
	}

	input := rol.UpperCase()

	i := entity.Role{
		RoleName:  input.RoleName,
		RoleNivel: input.RoleNivel,
	}

	return s.repo.CreateRole(ctx, i)
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

func (s *RolesService) DeleteById(ctx context.Context, idRol int) (int64, error) {
	ur, err := s.repo.GetUserByIdRole(ctx, idRol)
	if err != nil {
		return 0, err
	}
	if ur != nil {
		return 0, errors.New("no se puede eliminar el rol porque tiene usuarios asociados")
	}
	r, err := s.repo.DeleteById(ctx, idRol)
	if err != nil {
		return 0, err
	}
	return r, nil
}
