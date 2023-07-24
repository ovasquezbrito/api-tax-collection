package service

import (
	"errors"
	"fmt"
	"strings"
	"time"

	baseapp "github.com/ovasquezbrito/base-app"
	"github.com/ovasquezbrito/base-app/pkg/repository"
	"github.com/ovasquezbrito/base-app/token"
	"github.com/ovasquezbrito/base-app/util"
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
type loginUserResponse struct {
	AccessToken string       `json:"access_"`
	UserLogin   baseapp.User `json:"user"`
}

func NewAuthService(repo repository.Authorization, tokenMaker token.Maker, config util.Config) *AuthService {
	return &AuthService{repo: repo, tokenMaker: tokenMaker, config: config}
}

func (s *AuthService) CreateUser(user baseapp.User) (int, error) {
	hashedPassword, err := util.HashPassword(user.Password)
	if err != nil {
		return 0, errors.New("invalid signing method")
	}
	user.Password = hashedPassword
	input := user.UpperCase()

	return s.repo.CreateUser(*input)
}

func (s *AuthService) UpdateUser(idUser int, user baseapp.User) (int, error) {
	hashedPassword, err := util.HashPassword(user.Password)
	if err != nil {
		return 0, errors.New("invalid signing method")
	}
	user.Password = hashedPassword
	input := user.UpperCase()
	return s.repo.UpdateUser(idUser, *input)
}

func (s *AuthService) GetUserById(IdUser int) (baseapp.User, error) {
	return s.repo.GetUserById(IdUser)
}

func (s *AuthService) GetUserByUserName(email string) (int, error) {
	return s.repo.GetUserByUserName(strings.ToLower(email))
}

func (s *AuthService) GetMenuOptionAll(idUser int) ([]baseapp.RoleUser, error) {
	return s.repo.GetMenuOptionAll(idUser)
}

func (s *AuthService) LoginUser(email, password string) (loginUserResponse, error) {

	user, err := s.repo.GetUserByUserEmail(strings.ToLower(email))
	if err != nil {
		return loginUserResponse{}, err
	}

	err = util.CheckPassword(password, user.Password)
	if err != nil {
		return loginUserResponse{}, err
	}

	fmt.Println(s.config.AccessTokenDuration)
	token, err := s.tokenMaker.CreateToken(email, s.config.AccessTokenDuration)
	if err != nil {
		return loginUserResponse{}, errors.New("No se pudo generar el token")
	}
	return loginUserResponse{
		UserLogin:   user,
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
