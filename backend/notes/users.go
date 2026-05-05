package notes

import (
	"database/sql"
	"errors"
)

func RegisterUser(username, email, password string) (*User, error) {
	hashedPassword, err := HashPassword(password)
	if err != nil {
		return nil, err
	}

	var userID uint
	err = DB.QueryRow(
		"INSERT INTO users (username, email, password_hash) VALUES ($1, $2, $3) RETURNING id",
		username, email, hashedPassword,
	).Scan(&userID)

	if err != nil {
		if err.Error() == "pq: duplicate key value violates unique constraint \"users_username_key\"" {
			return nil, errors.New("username already exists")
		}
		if err.Error() == "pq: duplicate key value violates unique constraint \"users_email_key\"" {
			return nil, errors.New("email already exists")
		}
		return nil, err
	}

	return &User{
		ID:       userID,
		Username: username,
	}, nil
}

func LoginUser(email, password string) (*User, error) {
	var user User
	var passwordHash string

	err := DB.QueryRow(
		"SELECT id, username, password_hash FROM users WHERE email = $1",
		email,
	).Scan(&user.ID, &user.Username, &passwordHash)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	if !CheckPassword(passwordHash, password) {
		return nil, errors.New("incorrect password")
	}

	return &user, nil
}

func GetUserByID(userID uint) (*User, error) {
	var user User

	err := DB.QueryRow(
		"SELECT id, username FROM users WHERE id = $1",
		userID,
	).Scan(&user.ID, &user.Username)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	return &user, nil
}
