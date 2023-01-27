package model_test

import (
	"github.com/olegvolkov91/http-rest-api/internal/app/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUser_Validate(t *testing.T) {
	testCases := []struct {
		name    string
		user    func() *model.User
		isValid bool
	}{
		{
			name: "valid",
			user: func() *model.User {
				return model.TestUser(t)
			},
			isValid: true,
		},
		{
			name: "empty email",
			user: func() *model.User {
				u := model.TestUser(t)
				u.Email = ""
				return u
			},
			isValid: false,
		},
		{
			name: "invalid email",
			user: func() *model.User {
				u := model.TestUser(t)
				u.Email = "wrong_format_email.org"
				return u
			},
			isValid: false,
		},
		{
			name: "empty password",
			user: func() *model.User {
				u := model.TestUser(t)
				u.Password = ""
				return u
			},
			isValid: false,
		},
		{
			name: "short pssword",
			user: func() *model.User {
				u := model.TestUser(t)
				u.Password = "12345"
				return u
			},
			isValid: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.isValid {
				assert.NoError(t, tc.user().Validate())
			} else {
				assert.Error(t, tc.user().Validate())
			}
		})
	}
}

func TestUser_HashPassword(t *testing.T) {
	u := model.TestUser(t)
	assert.NoError(t, u.HashPassword())
	assert.NotEmpty(t, u.EncryptedPassword)
}
