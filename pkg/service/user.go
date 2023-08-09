package service

import (
	"context"
	"errors"

	"github.com/ovasquezbrito/tax-collection/pkg/entity"
	"github.com/ovasquezbrito/tax-collection/pkg/models"
	"github.com/ovasquezbrito/tax-collection/pkg/repository"
)

type UsersService struct {
	repo  repository.UserRepository
	repo2 repository.Authorization
	repo3 repository.RoleRepository
}

func NewUsersService(
	repo repository.UserRepository,
	repo2 repository.Authorization,
	repo3 repository.RoleRepository,
) *UsersService {
	return &UsersService{repo: repo, repo2: repo2, repo3: repo3}
}

func (s *UsersService) UpdateRoleToUser(ctx context.Context, userRole models.AsociateRoleToUser) error {
	idUser := userRole.IdUser
	idRole := userRole.IdRole
	u, _ := s.repo2.GetUserById(ctx, idUser)
	if u.Id != 0 {
		return errors.New("usuario no existe")
	}

	r, _ := s.repo3.GetById(ctx, idRole)
	if r.Id != 0 {
		return errors.New("rol no existe")
	}

	return s.repo.UpdateRoleToUser(ctx, idUser, idRole)
}

func (s *UsersService) GetAll(ctx context.Context, query models.QueryParameter) ([]models.UserResponse, int, error) {
	queryUpper := query.UpperCase()
	q := &entity.QueryParameter{
		Page:   queryUpper.Page,
		Limit:  queryUpper.Limit,
		Search: queryUpper.Search,
	}
	uu, total, err := s.repo.GetAll(ctx, *q)
	if err != nil {
		return nil, 0, err

	}

	mu := []models.UserResponse{}

	for _, u := range uu {
		mu = append(mu, models.UserResponse{
			Id:         u.Id,
			FirstLast:  u.FirstLast,
			Email:      u.Email,
			AvatarUser: u.AvatarUser,
			Status:     u.Status,
			IsAdmin:    u.IsAdmin,
		})
	}

	return mu, total, nil
}
