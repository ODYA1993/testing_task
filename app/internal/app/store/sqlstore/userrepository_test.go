package sqlstore_test

import (
	"github.com/DmitryOdintsov/rest-api/app/internal/app/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserRepository_Create(t *testing.T) {
	db, teardown := TestDB(t, databaseUrl)
	defer teardown("user")

	s := NewStore(db)
	u := model.TestUser(t)
	assert.NoError(t, s.User().Create(u))
	assert.NotNil(t, u)
}

func TestUserRepository_FindByEmail(t *testing.T) {
	db, teardown := TestDB(t, databaseUrl)
	defer teardown("user")

	s := NewStore(db)
	u1 := model.TestUser(t)
	_, err := s.User().FindByEmail(u1.Email)
	//assert.EqualError(t, err, store.ErrRecordNotFound.Error())

	s.User().Create(u1)
	u2, err := s.User().FindByEmail(u1.Email)
	assert.NoError(t, err)
	assert.NotNil(t, u2)
}

func TestUserRepository_Find(t *testing.T) {
	db, teardown := TestDB(t, databaseUrl)
	defer teardown("user")

	s := NewStore(db)
	u1 := model.TestUser(t)
	_ = s.User().Create(u1)
	u2, err := s.User().Find(u1.ID)
	assert.NoError(t, err)
	assert.NotNil(t, u2)
}
