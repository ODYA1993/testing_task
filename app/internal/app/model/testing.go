package model

import "testing"

func TestUser(t *testing.T) *User {
	t.Helper()
	return &User{
		Name:        "Dima",
		Surname:     "Odintsov",
		Email:       "user@example.org",
		Password:    "password",
		Patronymic:  "Sergeevith",
		Age:         30,
		Gender:      "Мужской",
		Nationality: "Русский",
	}
}
