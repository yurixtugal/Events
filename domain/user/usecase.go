package user

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/yurixtugal/Events/model"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	storage Storage
}

func New(s Storage) User {
	return User{storage: s}
}

func (u User) Create(m *model.User) error {
	ID, err := uuid.NewUUID()

	if err != nil {
		return fmt.Errorf("%s %w", "uuid.NewUUID()", err)
	}
	m.ID = ID

	passowrd, err := bcrypt.GenerateFromPassword([]byte(m.Password), bcrypt.DefaultCost)

	if err != nil {
		return fmt.Errorf("%s %w", "bcrypt.GenerateFromPassword", err)
	}

	m.Password = string(passowrd)

	if m.Details == nil {
		m.Details = []byte("{}")
	}

	err = u.storage.Create(m)

	if err != nil {
		return fmt.Errorf("%s %w", "u.storage.Create", err)
	}

	m.Password = ""

	return nil

}

func (u User) GetByEmail(email string) (model.User, error) {
	user, err := u.storage.GetByEmail(email)
	if err != nil {
		return model.User{}, fmt.Errorf("%s %w", "u.storage.GetByEmail", err)
	}
	return user, nil
}

func (u User) GetAll() (model.Users, error) {
	users, err := u.storage.GetAll()
	if err != nil {
		return nil, fmt.Errorf("%s %w", "u.storage.GetAll", err)
	}
	return users, nil
}
