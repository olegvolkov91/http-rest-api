package teststore_test

import (
	"github.com/olegvolkov91/http-rest-api/internal/app/model"
	"github.com/olegvolkov91/http-rest-api/internal/app/store"
	"github.com/olegvolkov91/http-rest-api/internal/app/store/teststore"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserRepository_Create(t *testing.T) {
	s := teststore.New()
	u := model.TestUser(t)

	assert.NoError(t, s.User().Create(u))
	assert.NotNil(t, u)
}

func TestUserRepository_FindByEmail(t *testing.T) {
	s := teststore.New()

	correctEmail := "user@example.org"
	wrongEmail := "test@i.ua"
	_, err := s.User().FindByEmail(wrongEmail)
	assert.EqualError(t, err, store.ErrRecordNotFound.Error())

	u, err := s.User().FindByEmail(correctEmail)
	assert.NoError(t, err)
	assert.NotNil(t, u)
}
