package service

import (
	"fmt"

	"github.com/MuhammedAshifVnr/user_service/internal/models"
	"github.com/MuhammedAshifVnr/user_service/internal/repo"
)

// UserService defines business logic methods for user operations
type UserService struct {
	repo repo.UserRepository
}

// NewUserService creates a new UserService
func NewUserService(repo repo.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) GetUserByID(id uint) (*models.User, error) {
	return s.repo.GetUserByID(id)
}

func (s *UserService) GetUsersByIDs(ids []uint) ([]models.User, error) {
	return s.repo.GetUsersByIDs(ids)
}

func (s *UserService) SearchUsers(city, phone, query string, married bool, limit, offset int) ([]models.User, error) {
	return s.repo.SearchUsers(city, phone, query, married, limit, offset)
}

func (s *UserService) CreateUser(user *models.User) error {
	if err := user.Validate(); err != nil {
		fmt.Println("==",err)
		return err
	}
	return s.repo.CreateUser(user)
}
