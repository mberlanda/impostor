package db

import (
	"sync"
	"testing"

	"github.com/go-playground/assert/v2"
	"github.com/mberlanda/impostor/api/models"
)

var userOne = &models.User{
	ID:       123,
	Email:    "foo@bar.baz",
	Password: "asupersecretpassword",
}

var userTwo = &models.User{
	ID:       456,
	Email:    "baz@bar.foo",
	Password: "anothersecretpassword",
}

func prepareUserDb(users ...*models.User) *UserDb {
	userDb := InitUserDb()
	wg := sync.WaitGroup{}

	for _, user := range users {
		wg.Add(1)
		u := user
		go func() {
			userDb.Add(u.ID, u)
			wg.Done()
		}()
	}

	wg.Wait()

	return &userDb
}

func TestPrepareUserDbWhenEmptyUserList(t *testing.T) {
	userDb := prepareUserDb()
	assert.Equal(t, 0, userDb.Size())
}

func TestAddUserDbWhenSingletonUserList(t *testing.T) {
	userDb := prepareUserDb(userOne)
	assert.Equal(t, 1, userDb.Size())
}

func TestAddUserDbWhenMultiItemsUserList(t *testing.T) {
	userDb := prepareUserDb(userOne, userTwo)
	assert.Equal(t, 2, userDb.Size())
}

func TestAddUserDbWhenDuplicatedUserList(t *testing.T) {
	userDb := prepareUserDb(userOne, userOne, userTwo)
	assert.Equal(t, 2, userDb.Size())
}

func TestGetUserDbWhenKeyPresent(t *testing.T) {
	userDb := prepareUserDb(userOne, userTwo)
	userFound, found := userDb.Get(userOne.ID)

	assert.Equal(t, true, found)
	assert.Equal(t, userOne.Email, userFound.Email)
}

func TestGetUserDbWhenKeyNotPresent(t *testing.T) {
	userDb := prepareUserDb(userTwo)
	userFound, found := userDb.Get(userOne.ID)

	assert.Equal(t, false, found)
	assert.Equal(t, nil, userFound)
}

func TestDeletetUserDbWhenKeyPresent(t *testing.T) {
	userDb := prepareUserDb(userOne, userTwo)

	_, found := userDb.Get(userOne.ID)
	assert.Equal(t, true, found)

	ok := userDb.Delete(userOne.ID)
	assert.Equal(t, true, ok)

	_, found = userDb.Get(userOne.ID)
	assert.Equal(t, false, found)
}

func TestDeleteUserDbWhenKeyNotPresent(t *testing.T) {
	userDb := prepareUserDb(userTwo)

	ok := userDb.Delete(userOne.ID)
	assert.Equal(t, false, ok)
}
