package repository

import (
	"github.com/mberlanda/impostor/api/db"
	"github.com/mberlanda/impostor/api/models"
)

type UserRepository struct {
	db db.UserDb
}

func NewUserRepository(userDb *db.UserDb) UserRepository {
	return UserRepository{
		db: *userDb,
	}
}

func (r *UserRepository) AddUser(u *models.User) (*models.User, bool) {
	// TODO: refactor changing the input type and using this method to
	// map and trasform to some user db model
	return r.db.Add(u.ID, u)
}

func (r *UserRepository) GetUser(id int) (*models.User, bool) {
	// TODO: refactor changing the input type and using this method to
	// map and trasform to some user db model
	return r.db.Get(id)
}

func (r *UserRepository) DeleteUser(id int) bool {
	// TODO: refactor changing the input type and using this method to
	// map and trasform to some user db model
	return r.db.Delete(id)
}

func (r *UserRepository) GetAllUser() []*models.User {
	// TODO: refactor changing the input type and using this method to
	// map and trasform to some user db model
	return r.db.All()
}
