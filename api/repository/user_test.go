package repository

import (
	"testing"

	"github.com/go-playground/assert/v2"
	"github.com/mberlanda/impostor/api/db"
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

func TestRepositoryLifeCycle(t *testing.T) {
	udb := db.InitUserDb()
	r := NewUserRepository(&udb)
	r.AddUser(userOne)

	assert.Equal(t, 1, len(r.GetAllUser()))

	_, found := r.GetUser(userOne.ID)
	assert.Equal(t, true, found)
	_, found = r.GetUser(userTwo.ID)
	assert.Equal(t, false, found)

	ok := r.DeleteUser(userOne.ID)
	assert.Equal(t, true, ok)
	ok = r.DeleteUser(userOne.ID)
	assert.Equal(t, false, ok)

	assert.Equal(t, 0, len(r.GetAllUser()))
}
