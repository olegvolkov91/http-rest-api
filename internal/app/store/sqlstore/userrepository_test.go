package sqlstore_test

import (
	"github.com/olegvolkov91/http-rest-api/internal/app/model"
	"github.com/olegvolkov91/http-rest-api/internal/app/store"
	"github.com/olegvolkov91/http-rest-api/internal/app/store/sqlstore"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserRepository_Create(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("users")

	s := sqlstore.New(db)
	u := model.TestUser(t)

	assert.NoError(t, s.User().Create(u))
	assert.NotNil(t, u)
}

func TestUserRepository_FindByEmail(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("users")
	s := sqlstore.New(db)

	correctEmail := "user@example.org"
	wrongEmail := "test@i.ua"
	_, err := s.User().FindByEmail(wrongEmail)
	assert.EqualError(t, err, store.ErrRecordNotFound.Error())

	u := model.TestUser(t)
	u.Email = correctEmail
	s.User().Create(u)
	u, err = s.User().FindByEmail(correctEmail)
	assert.NoError(t, err)
	assert.NotNil(t, u)
}
