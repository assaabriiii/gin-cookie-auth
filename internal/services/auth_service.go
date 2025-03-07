package services

import (
	"errors"

	"github.com/assaabriiii/gin-cookie-auth/internal/models"
	"github.com/assaabriiii/gin-cookie-auth/internal/repositories"
)

type AuthService struct {
	userRepo *repositories.UserRepository
}

func NewAuthService(userRepo *repositories.UserRepository) *AuthService {
	return &AuthService{userRepo: userRepo}
}

func (s *AuthService) Register(user *models.User) error {
	_, exists := s.userRepo.FindByUsername(user.Username)
	if exists {
		return errors.New("username already exists")
	}
	s.userRepo.Save(user)
	return nil
}

func (s *AuthService) Login(username, password string) (*models.User, error) {
	user, exists := s.userRepo.FindByUsername(username)
	if !exists || user.Password != password {
		return nil, errors.New("invalid credentials")
	}
	return user, nil
}
