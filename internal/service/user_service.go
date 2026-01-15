package service

import (
	"errors"
	"time"
)

var (
	ErrNotFound      = errors.New("user not found")
	ErrAlreadyExists = errors.New("user already exists")
)

type User struct {
	ID          string
	Name        string
	DisplayName string
	Email       string
	Active      bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type Store interface {
	Create(user User) (User, error)
	GetByID(id string) (User, error)
	//GetByEmail(email string) (User, error)
	Update(id string, displayName *string) (User, error)
	//Delete(id string) error
}
