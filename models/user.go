package models

import (
	"errors"

	"github.com/JarGad23/go-rest-api/db"
	"github.com/JarGad23/go-rest-api/utils"
)

type User struct {
	Id       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (u User) Save() error {
	query := `
		INSERT INTO users(email, password) VALUES(?, ?)
	`
	sqlStmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer sqlStmt.Close()

	hashedPassword, err := utils.HashPassword(u.Password)

	if err != nil {
		return err
	}

	result, err := sqlStmt.Exec(u.Email, hashedPassword)

	if err != nil {
		return err
	}

	userId, err := result.LastInsertId()

	u.Id = userId
	return err
}

func (u User) ValidateCredentials() error {
	query := "SELECT id, password FROM users WHERE email = ?"

	row := db.DB.QueryRow(query, u.Email)

	var retrivedPassword string
	err := row.Scan(&u.Id, &retrivedPassword)

	if err != nil {
		return errors.New("Invalid credentials")
	}

	passwordIsValid := utils.CheckPasswordHash(u.Password, retrivedPassword)

	if !passwordIsValid {
		return errors.New("Invalid credentials")
	}

	return nil
}
