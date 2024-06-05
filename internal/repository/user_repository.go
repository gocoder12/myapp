package repository

import (
	"database/sql"
	"myapp/internal/model"

	_ "github.com/lib/pq"
)

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db}
}

func (r *userRepository) FetchAllUsers() ([]model.User, error) {
	rows, err := r.db.Query("SELECT id, name, email, password FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := []model.User{}
	for rows.Next() {
		var user model.User
		if err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.Password); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (r *userRepository) FetchUserByID(id string) (model.User, error) {
	var user model.User
	err := r.db.QueryRow("SELECT id, name, email, password FROM users WHERE id=$1", id).Scan(&user.ID, &user.Name, &user.Email, &user.Password)
	if err != nil {
		return user, err
	}
	return user, nil
}

func (r *userRepository) CreateUser(user model.User) error {
	_, err := r.db.Exec("INSERT INTO users (id, name, email, password) VALUES ($1, $2, $3, $4)", user.ID, user.Name, user.Email, user.Password)
	return err
}

func (r *userRepository) UpdateUser(id string, user model.User) error {
	_, err := r.db.Exec("UPDATE users SET name=$1, email=$2, password=$3 WHERE id=$4", user.Name, user.Email, user.Password, id)
	return err
}

func (r *userRepository) DeleteUser(id string) error {
	_, err := r.db.Exec("DELETE FROM users WHERE id=$1", id)
	return err
}
