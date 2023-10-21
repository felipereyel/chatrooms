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

func NewDatabaseRepo() (Database, error) {
	if config.Configs.DatabaseConnString == "" {
		return nil, errors.New("DatabaseConnString is not set")
	}

	conn, err := sql.Open("postgres", config.Configs.DatabaseConnString)
	if err != nil {
		return nil, err
	}

	return &database{conn}, nil
}

func (db *database) Close() error {
	return db.conn.Close()
}

func (db *database) userGetName(userId string) (string, error) {
	query := `SELECT username FROM users WHERE id = $1`
	row := db.conn.QueryRow(query, userId)

	var username string
	err := row.Scan(&username)
	if err != nil {
		// TODO handle not found
		return "", err
	}

	return username, nil
}

func (db *database) UserGetId(username string) (string, error) {
	query := `SELECT id FROM users WHERE username = $1`
	row := db.conn.QueryRow(query, username)

	var userId string
	err := row.Scan(&userId)
	if err != nil {
		// TODO handle not found
		return "", err
	}

	return userId, nil
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

func (db *database) ListRooms() ([]models.Room, error) {
	query := `SELECT id, name FROM rooms`
	rows, err := db.conn.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	rooms := []models.Room{}
	for rows.Next() {
		var room models.Room
		err := rows.Scan(&room.Id, &room.Name)
		if err != nil {
			return nil, err
		}

		rooms = append(rooms, room)
	}

	return rooms, nil
}

func (db *database) CreateRoom(room models.Room) error {
	query := `INSERT INTO rooms (id, name) VALUES ($1, $2)`
	_, err := db.conn.Exec(query, room.Id, room.Name)
	if err != nil {
		// TODO handle conflict
		return err
	}

	return nil
}

func (db *database) GetRoom(roomId string) (models.Room, error) {
	query := `SELECT id, name FROM rooms WHERE id = $1`
	row := db.conn.QueryRow(query, roomId)

	var room models.Room
	err := row.Scan(&room.Id, &room.Name)
	if err != nil {
		// TODO handle not found
		return models.Room{}, err
	}

	return room, nil
}

func (db *database) ListPosts(roomId string) ([]models.PostView, error) {
	query := `SELECT posts.id, posts.content, users.username, posts.created_at, posts.room_id FROM posts INNER JOIN users ON posts.user_id = users.id WHERE posts.room_id = $1 ORDER BY posts.created_at DESC LIMIT 50`
	rows, err := db.conn.Query(query, roomId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	posts := []models.PostView{}
	for rows.Next() {
		var post models.PostView
		err := rows.Scan(&post.Id, &post.Content, &post.Username, &post.CreatedAt, &post.RoomId)
		if err != nil {
			return nil, err
		}

		posts = append(posts, post)
	}

	return posts, nil
}

func (db *database) CreatePost(post models.Post) (models.PostView, error) {
	postView := models.PostView{
		Id:      post.Id,
		Content: post.Content,
		RoomId:  post.RoomId,
	}

	query := `INSERT INTO posts (id, content, user_id, room_id) VALUES ($1, $2, $3, $4) RETURNING created_at`
	row := db.conn.QueryRow(query, post.Id, post.Content, post.UserId, post.RoomId)
	if err := row.Scan(&postView.CreatedAt); err != nil {
		// TODO handle conflict
		return models.PostView{}, err
	}

	username, err := db.userGetName(post.UserId)
	if err != nil {
		return models.PostView{}, err
	}

	postView.Username = username
	return postView, nil
}
