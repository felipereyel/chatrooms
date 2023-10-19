package database

import (
	"chatrooms/gosrc/config"
	"chatrooms/gosrc/models"
	"chatrooms/gosrc/utils"
	"database/sql"
	"errors"

	_ "github.com/lib/pq"
)

type database struct {
	conn *sql.DB
}

func NewDatabaseRepo() (*database, error) {
	conn, err := sql.Open("postgres", config.Configs.DatabaseConnString)
	if err != nil {
		return nil, err
	}

	return &database{conn}, nil
}

func (db *database) Close() error {
	return db.conn.Close()
}

func (db *database) UserLogin(username string, password string) (models.User, error) {
	query := `SELECT id, username, pass FROM users WHERE username = $1`
	row := db.conn.QueryRow(query, username)

	var user models.User
	err := row.Scan(&user.Id, &user.Username, &user.Pass)
	if err != nil {
		// TODO handle not found
		return models.User{}, err
	}

	// slow checking -> design feature of bcrypt
	if !utils.CheckPasswordHash(password, user.Pass) {
		return models.User{}, errors.New("bad username or password")
	}

	return user, nil
}

func (db *database) UserRegister(user models.User) error {
	query := `INSERT INTO users (id, username, pass) VALUES ($1, $2, $3)`
	_, err := db.conn.Exec(query, user.Id, user.Username, user.Pass)
	if err != nil {
		// TODO handle conflict
		return err
	}

	return nil
}
