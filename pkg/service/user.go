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

func (s *UsersService) AddRoleToUser(ctx context.Context, userRole models.AsociateRoleToUser) (int64, error) {
	idUser := userRole.IdUser
	idRole := userRole.IdRole
	_, err := s.repo2.GetUserById(ctx, idUser)
	if err != nil {
		return 0, errors.New("user no existe")
	}

	_, err = s.repo3.GetById(ctx, idRole)
	if err != nil {
		return 0, errors.New("rol no existe")
	}

	af, err := s.repo.AddRoleToUser(ctx, idUser, idRole)
	if err != nil {
		return 0, err
	}
	return af, nil
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
			FkRole:     u.FkRole,
		})
	}

	return mu, total, nil
}
