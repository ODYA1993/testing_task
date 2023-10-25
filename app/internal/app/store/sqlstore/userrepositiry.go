package sqlstore

import (
	"database/sql"
	"github.com/DmitryOdintsov/rest-api/app/internal/app/model"
	"github.com/DmitryOdintsov/rest-api/app/internal/app/store"
)

type UserRepository struct {
	store *Store
}

func (r *UserRepository) Create(u *model.User) error {
	if err := u.Validate(); err != nil {
		return err
	}

	if err := u.BeforeCreate(); err != nil {
		return err
	}

	return r.store.db.QueryRow(
		`INSERT INTO "user" (name, surname, email, encrypted_password, patronymic, age, gender, nationality) 
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id`,
		u.Name, u.Surname, u.Email, u.EncryptedPassword, u.Patronymic, u.Age, u.Gender, u.Nationality,
	).Scan(&u.ID)
}

func (r *UserRepository) Find(id int) (*model.User, error) {
	u := &model.User{}
	if err := r.store.db.QueryRow(
		`SELECT id, name, surname, email, encrypted_password, patronymic, age, gender, nationality
	  		   FROM "user" WHERE id = $1`,
		id).Scan(
		&u.ID,
		&u.Name,
		&u.Surname,
		&u.Email,
		&u.EncryptedPassword,
		&u.Patronymic,
		&u.Age,
		&u.Gender,
		&u.Nationality,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, store.ErrRecordNotFound
		}
		return nil, err
	}

	return u, nil
}

func (r *UserRepository) FindByEmail(email string) (*model.User, error) {
	u := &model.User{}
	if err := r.store.db.QueryRow(
		`SELECT id, name, surname, email, encrypted_password, patronymic, age, gender, nationality
	  		   FROM "user" WHERE email = $1`,
		email).Scan(
		&u.ID,
		&u.Name,
		&u.Surname,
		&u.Email,
		&u.EncryptedPassword,
		&u.Patronymic,
		&u.Age,
		&u.Gender,
		&u.Nationality,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, store.ErrRecordNotFound
		}
		return nil, err
	}

	return u, nil
}
