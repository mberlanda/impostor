package db

import (
	"sync"

	"github.com/mberlanda/impostor/api/models"
)

type UserDb struct {
	m   map[int]*models.User
	mux *sync.RWMutex
}

func InitUserDb() UserDb {
	return UserDb{
		m:   make(map[int]*models.User),
		mux: &sync.RWMutex{},
	}
}

func (u *UserDb) Add(key int, user *models.User) (*models.User, bool) {
	u.mux.Lock()
	// TODO: missing error handling and overrides
	u.m[(*user).ID] = user
	u.mux.Unlock()

	return user, true
}

func (u *UserDb) Get(key int) (*models.User, bool) {
	u.mux.RLock()
	user := u.m[key]
	u.mux.RUnlock()

	return user, (user != nil)
}

func (u *UserDb) Delete(key int) bool {
	u.mux.Lock()
	_, found := u.m[key]
	if found {
		delete(u.m, key)
	}
	u.mux.Unlock()

	return found
}

func (u *UserDb) Size() int {
	if u.m == nil {
		return 0
	}
	return len(u.m)
}

func (u *UserDb) All() []*models.User {
	users := make([]*models.User, 0, u.Size())
	for _, u := range u.m {
		users = append(users, u)
	}
	return users
}
