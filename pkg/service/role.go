package service

import (
	"context"
	"errors"
	"fmt"
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
	r := strings.ToLower(rol.RoleName)
	r = strings.TrimSpace(r)
	fmt.Println(r)
	u, _ := s.repo.GetRoleByName(ctx, r)
	fmt.Println(u)
	if u.Id != 0 {
		return 0, errors.New("user already exists")
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

	r, err := s.repo.DeleteById(ctx, idRol)
	if err != nil {
		return 0, err
	}
	return r, nil
}
