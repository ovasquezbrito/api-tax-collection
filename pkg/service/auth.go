package service

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/ovasquezbrito/tax-collection/pkg/entity"
	"github.com/ovasquezbrito/tax-collection/pkg/models"
	"github.com/ovasquezbrito/tax-collection/pkg/repository"
	"github.com/ovasquezbrito/tax-collection/token"
	"github.com/ovasquezbrito/tax-collection/util"
)

const (
	saltUser       = "hjgrhjgw1234567ajfhajs"
	signingKeyUser = "grkjk#4#%35FSFJlja#4353KsfjH"
	tokenTTLUser   = 12 * time.Hour
)

type AuthService struct {
	repo       repository.Authorization
	tokenMaker token.Maker
	config     util.Config
}

func NewAuthService(repo repository.Authorization, tokenMaker token.Maker, config util.Config) *AuthService {
	return &AuthService{repo: repo, tokenMaker: tokenMaker, config: config}
}

func (s *AuthService) CreateUser(ctx context.Context, user models.User) (int, error) {
	e := strings.ToLower(user.Email)
	u, _ := s.repo.GetUserByUserEmail(ctx, e)
	if u.Id != 0 {
		return 0, errors.New("user already exists")
	}

	hashedPassword, err := util.HashPassword(user.Password)
	if err != nil {
		return 0, errors.New("invalid signing method")
	}
	user.Password = hashedPassword
	input := user.UpperCase()

	i := entity.User{
		FirstLast:  input.FirstLast,
		Email:      input.Email,
		Password:   input.Password,
		AvatarUser: input.AvatarUser,
	}

	return s.repo.CreateUser(ctx, i)
}

func (s *AuthService) UpdateUser(ctx context.Context, idUser int, user models.User) (int, error) {
	hashedPassword, err := util.HashPassword(user.Password)
	if err != nil {
		return 0, errors.New("invalid signing method")
	}
	user.Password = hashedPassword
	input := user.UpperCase()
	i := entity.User{
		Id:         idUser,
		FirstLast:  input.FirstLast,
		Email:      input.Email,
		Password:   input.Password,
		AvatarUser: input.AvatarUser,
	}
	return s.repo.UpdateUser(ctx, idUser, i)
}

func (s *AuthService) GetUserById(ctx context.Context, IdUser int) (*models.User, error) {
	u, err := s.repo.GetUserById(ctx, IdUser)
	if err != nil {
		return nil, err
	}

	return &models.User{
		Id:         u.Id,
		FirstLast:  u.FirstLast,
		Email:      u.Email,
		Password:   u.Password,
		AvatarUser: u.AvatarUser,
	}, nil
}

func (s *AuthService) GetMenuOptionAll(ctx context.Context, idUser int) ([]models.RoleUser, error) {
	mm, err := s.repo.GetMenuOptionAll(ctx, idUser)
	if err != nil {
		return nil, err
	}

	menuRoles := []models.RoleUser{}

	for _, m := range mm {
		menuRoles = append(menuRoles, models.RoleUser{
			IdRole:       m.Id,
			NameOpcion:   m.NameOpcion,
			Icon:         m.Icon,
			ComponentUri: m.ComponentUri,
			PageUrl:      m.PageUrl,
			OrderBy:      m.OrderBy,
			TypeOpcion:   m.TypeOpcion,
			NivelOpcion:  m.NivelOpcion,
			CreatedAt:    m.CreatedAt,
			UpdatedAt:    m.UpdatedAt,
			Status:       m.Status,
		})
	}

	return menuRoles, nil
}

func (s *AuthService) LoginUser(ctx context.Context, email, password string) (models.LoginUserResponse, error) {

	user, err := s.repo.GetUserByUserEmail(ctx, strings.ToLower(email))
	if err != nil {
		return models.LoginUserResponse{}, err
	}

	err = util.CheckPassword(password, user.Password)
	if err != nil {
		return models.LoginUserResponse{}, err
	}

	fmt.Println(s.config.AccessTokenDuration)
	token, err := s.tokenMaker.CreateToken(email, s.config.AccessTokenDuration)
	if err != nil {
		return models.LoginUserResponse{}, errors.New("No se pudo generar el token")
	}

	u := &models.User{
		Id:         user.Id,
		FirstLast:  user.FirstLast,
		Email:      user.Email,
		Password:   user.Password,
		AvatarUser: user.AvatarUser,
	}

	return models.LoginUserResponse{
		UserLogin:   u,
		AccessToken: token,
	}, nil
}

func (s *AuthService) ParseToken(accessToken string) (string, error) {
	payload, err := s.tokenMaker.VerifyToken(accessToken)
	if err != nil {
		return "", errors.New("invalid signing method")
	}

	return payload.Email, nil

}

func (s *AuthService) VerifyToken(accessToken string) (*token.Payload, error) {
	payload, err := s.tokenMaker.VerifyToken(accessToken)
	if err != nil {
		return &token.Payload{}, err
	}

	return payload, nil

}
