package store_test

import (
	"github.com/olegvolkov91/http-rest-api/internal/app/model"
	"github.com/olegvolkov91/http-rest-api/internal/app/store"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserRepository_Create(t *testing.T) {
	s, teardown := store.TestStore(t, databaseURL)
	defer teardown("users")

	u, err := s.User().Create(model.TestUser(t))

	assert.NoError(t, err)
	assert.NotNil(t, u)
}

func TestUserRepository_FindByEmail(t *testing.T) {
	s, teardown := store.TestStore(t, databaseURL)
	defer teardown("users")

	correctEmail := "user@example.org"
	wrongEmail := "test@i.ua"
	_, err := s.User().FindByEmail(wrongEmail)
	assert.Error(t, err)

	u, err := s.User().FindByEmail(correctEmail)
	assert.NoError(t, err)
	assert.NotNil(t, u)
}
