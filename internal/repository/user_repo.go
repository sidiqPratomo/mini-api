package repository

import (
	"database/sql"
	"errors"

	"github.com/sidiqPratomo/mini-api/internal/domain"
)

type userRepo struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) domain.UserRepository {
	return &userRepo{db}
}

func (r *userRepo) Create(user *domain.User) error {
	query := `INSERT INTO users (name, email, password, created_at, updated_at)
			  VALUES ($1, $2, $3, NOW(), NOW()) RETURNING id`

	err := r.db.QueryRow(query, user.Name, user.Email, user.Password).Scan(&user.ID)
	return err
}

func (r *userRepo) FindByEmail(email string) (*domain.User, error) {
	query := `SELECT id, name, email, password, created_at, updated_at FROM users WHERE email=$1 LIMIT 1`
	row := r.db.QueryRow(query, email)

	var user domain.User
	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}

func (r *userRepo) FindByID(id uint) (*domain.User, error) {
	query := `SELECT id, name, email, password, created_at, updated_at FROM users WHERE id=$1 LIMIT 1`
	row := r.db.QueryRow(query, id)

	var user domain.User
	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}

func (r *userRepo) Update(user *domain.User) error {
	query := `UPDATE users SET name=$1, email=$2, password=$3, updated_at=NOW() WHERE id=$4`
	_, err := r.db.Exec(query, user.Name, user.Email, user.Password, user.ID)
	return err
}

func (r *userRepo) Delete(id uint) error {
	query := `DELETE FROM users WHERE id=$1`
	_, err := r.db.Exec(query, id)
	return err
}

func (r *userRepo) Fetch() ([]domain.User, error) {
	query := `SELECT id, name, email, password, created_at, updated_at FROM users ORDER BY id`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []domain.User
	for rows.Next() {
		var user domain.User
		err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}
