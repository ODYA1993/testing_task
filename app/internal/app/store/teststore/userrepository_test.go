package teststore

import (
	"github.com/DmitryOdintsov/rest-api/app/internal/app/model"
	"github.com/DmitryOdintsov/rest-api/app/internal/app/store"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserRepository_Create(t *testing.T) {
	s := NewStore()
	u := model.TestUser(t)
	assert.NoError(t, s.User().Create(u))
	assert.NotNil(t, u)
}

func TestUserRepository_FindByEmail(t *testing.T) {

	s := NewStore()
	u1 := model.TestUser(t)
	_, err := s.User().FindByEmail(u1.Email)
	assert.EqualError(t, err, store.ErrRecordNotFound.Error())

	err = s.User().Create(u1)
	if err != nil {
		return
	}
	u2 := model.TestUser(t)
	_, err = s.User().FindByEmail(u2.Email)
	assert.NoError(t, err)
	assert.NotNil(t, u2)
}

func TestUserRepository_Find(t *testing.T) {
	s := NewStore()
	u1 := model.TestUser(t)
	_ = s.User().Create(u1)
	u2, err := s.User().Find(u1.ID)
	assert.NoError(t, err)
	assert.NotNil(t, u2)
}
