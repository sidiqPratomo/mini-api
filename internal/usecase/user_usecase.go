package usecase

import (
	"errors"
	"time"

	"github.com/sidiqPratomo/mini-api/internal/domain"
	"golang.org/x/crypto/bcrypt"
)

type userUsecase struct {
	userRepo domain.UserRepository
}

func NewUserUsecase(repo domain.UserRepository) domain.UserUsecase {
	return &userUsecase{userRepo: repo}
}

func (uc *userUsecase) Register(name, email, password string) (*domain.User, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &domain.User{
		Name:     name,
		Email:    email,
		Password: string(hashed),
	}

	err = uc.userRepo.Create(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (uc *userUsecase) Login(email, password string) (*domain.User, error) {
	user, err := uc.userRepo.FindByEmail(email)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("user not found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, errors.New("invalid password")
	}

	return user, nil
}

func (uc *userUsecase) GetUser(id uint) (*domain.User, error) {
	return uc.userRepo.FindByID(id)
}

func (uc *userUsecase) ListUsers() ([]domain.User, error) {
	return uc.userRepo.Fetch()
}

func (uc *userUsecase) UpdateUser(user *domain.User) error {
	user.UpdatedAt = time.Now()
	return uc.userRepo.Update(user)
}

func (uc *userUsecase) DeleteUser(id uint) error {
	return uc.userRepo.Delete(id)
}
