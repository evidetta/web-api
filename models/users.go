package models

import (
	"database/sql"
	"time"
)

type User struct {
	Model
	Name        string
	Address     string
	DateOfBirth time.Time
}

func CreateUser(db *sql.DB, name, address string, dob time.Time) (*User, error) {
	user := User{
		Name:        name,
		Address:     address,
		DateOfBirth: dob,
	}

	now := time.Now().UTC()
	user.CreatedAt = now
	user.UpdatedAt = now
	user.DeletedAt = nil

	err := db.QueryRow(
		"INSERT INTO users (name, address, date_of_birth, created_at, updated_at, deleted_at) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id",
		user.Name,
		user.Address,
		user.DateOfBirth,
		user.CreatedAt,
		user.UpdatedAt,
		user.DeletedAt,
	).Scan(&user.ID)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func GetUserById(db *sql.DB, id int64) (*User, error) {
	user := User{}
	user.ID = id

	err := db.QueryRow("SELECT name, address, date_of_birth, created_at, updated_at, deleted_at FROM users WHERE id=$1 AND deleted_at IS NULL", user.ID).Scan(
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

	_, err = db.Exec(
		"UPDATE users SET name=$1, address=$2, date_of_birth=$3, created_at=$4, updated_at=$5, deleted_at=$6 where id=$7",
		user.Name,
		user.Address,
		user.DateOfBirth,
		user.CreatedAt,
		user.UpdatedAt,
		user.DeletedAt,
		user.ID,
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
		"UPDATE users SET updated_at=$1, deleted_at=$2 WHERE id=$3",
		user.UpdatedAt,
		user.DeletedAt,
		user.ID,
	)

	if err != nil {
		return err
	}

	return nil
}

func (user *User) checkForSoftDelete(db *sql.DB) error {
	err := db.QueryRow("SELECT deleted_at FROM users WHERE id=$1", user.ID).Scan(
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
