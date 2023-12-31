package models

import "github.com/JarGad23/go-rest-api/db"

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

	result, err := sqlStmt.Exec(u.Email, u.Password)

	if err != nil {
		return err
	}

	userId, err := result.LastInsertId()

	u.Id = userId
	return err
}
