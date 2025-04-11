package domain

import (
	"time"

)

type User struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserRepository interface {
	Create(user *User) error
	FindByEmail(email string) (*User, error)
	FindByID(id uint) (*User, error)
	Update(user *User) error
	Delete(id uint) error
	Fetch() ([]User, error)
}

type UserUsecase interface {
	Register(name, email, password string) (*User, error)
	Login(email, password string) (*User, error) // returns JWT
	GetUser(id uint) (*User, error)
	UpdateUser(user *User) error
	DeleteUser(id uint) error
	ListUsers() ([]User, error)
}
