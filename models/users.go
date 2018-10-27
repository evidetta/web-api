package models

import (
	"database/sql"
	"time"
)

type User struct {
	Model
	Name        string    `json:"name"`
	Address     string    `json:"address"`
	DateOfBirth time.Time `json:"date_of_birth"`
}

func CreateUser(db *sql.DB, user *User) (*User, error) {

	now := time.Now().UTC()
	user.CreatedAt = now
	user.UpdatedAt = now
	user.DeletedAt = nil

	user.DateOfBirth = time.Date(user.DateOfBirth.Year(), user.DateOfBirth.Month(), user.DateOfBirth.Day(), 0, 0, 0, 0, time.UTC)

	err := db.QueryRow(
		"INSERT INTO users (name, address, date_of_birth, created_at, updated_at, deleted_at) VALUES ($1, $2, $3, $4, $5, $6) RETURNING tag",
		user.Name,
		user.Address,
		user.DateOfBirth,
		user.CreatedAt,
		user.UpdatedAt,
		user.DeletedAt,
	).Scan(&user.Tag)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func GetUsers(db *sql.DB, pageSize, pageNumber int) ([]*User, error) {
	users := []*User{}

	limit := pageSize
	offset := (pageNumber - 1) * pageSize
	rows, err := db.Query("SELECT tag, name, address, date_of_birth, created_at, updated_at, deleted_at FROM users WHERE deleted_at IS NULL ORDER BY created_at LIMIT $1 OFFSET $2", limit, offset)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		user := User{}
		err = rows.Scan(
			&user.Tag,
			&user.Name,
			&user.Address,
			&user.DateOfBirth,
			&user.CreatedAt,
			&user.UpdatedAt,
			&user.DeletedAt,
		)

		if err != nil {
			return nil, err
		}

		users = append(users, &user)
	}

	return users, nil
}

func GetUserByTag(db *sql.DB, tag string) (*User, error) {
	user := User{}
	user.Tag = tag

	err := db.QueryRow("SELECT name, address, date_of_birth, created_at, updated_at, deleted_at FROM users WHERE tag=$1 AND deleted_at IS NULL", user.Tag).Scan(
		&user.Name,
		&user.Address,
		&user.DateOfBirth,
		&user.CreatedAt,
		&user.UpdatedAt,
		&user.DeletedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrorEntryNotFound
		}
		return nil, err
	}

	return &user, nil
}

func (user *User) Update(db *sql.DB) error {

	err := user.checkForSoftDelete(db)
	if err != nil {
		return err
	}

	now := time.Now().UTC()
	user.UpdatedAt = now

	user.DateOfBirth = time.Date(user.DateOfBirth.Year(), user.DateOfBirth.Month(), user.DateOfBirth.Day(), 0, 0, 0, 0, time.UTC)

	_, err = db.Exec(
		"UPDATE users SET name=$1, address=$2, date_of_birth=$3, created_at=$4, updated_at=$5, deleted_at=$6 where tag=$7",
		user.Name,
		user.Address,
		user.DateOfBirth,
		user.CreatedAt,
		user.UpdatedAt,
		user.DeletedAt,
		user.Tag,
	)

	if err != nil {
		return err
	}

	return nil
}

func (user *User) Delete(db *sql.DB) error {
	err := user.checkForSoftDelete(db)
	if err != nil {
		return err
	}

	now := time.Now().UTC()
	user.UpdatedAt = now
	user.DeletedAt = &now

	_, err = db.Exec(
		"UPDATE users SET updated_at=$1, deleted_at=$2 WHERE tag=$3",
		user.UpdatedAt,
		user.DeletedAt,
		user.Tag,
	)

	if err != nil {
		return err
	}

	return nil
}

func (user *User) checkForSoftDelete(db *sql.DB) error {
	err := db.QueryRow("SELECT deleted_at FROM users WHERE tag=$1", user.Tag).Scan(
		&user.DeletedAt,
	)

	if err != nil {
		return err
	}

	if user.DeletedAt != nil {
		return ErrorEntryNotFound
	}

	return nil
}
