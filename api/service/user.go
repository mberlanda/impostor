package service

import (
	"errors"
	"strconv"
	"time"

	"github.com/mberlanda/impostor/api/models"
	"github.com/mberlanda/impostor/api/repository"
)

type UserService struct {
	repository *repository.UserRepository
}

func NewUserService(r *repository.UserRepository) UserService {
	return UserService{
		repository: r,
	}
}

func (us *UserService) Create(u *models.User) (*models.User, error) {
	u.ID = int(time.Now().UnixMilli())
	if u.Email == "" {
		return nil, errors.New("email cannot be empty")
	}
	user, _ := us.repository.AddUser(u)
	return user, nil
}

func (us *UserService) GetAllUsers() []*models.User {
	return us.repository.GetAllUser()
}

func (us *UserService) GetUser(id string) (*models.User, error) {
	userId, err := strconv.Atoi(id)
	if err != nil {
		return nil, errors.New("invalid user id")
	}
	user, found := us.repository.GetUser(userId)
	if !found {
		return nil, errors.New("user not found")
	}
	return user, nil
}
