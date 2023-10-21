package database

import (
	"chatrooms/gosrc/models"
	"chatrooms/gosrc/utils"
	"errors"

	_ "github.com/lib/pq"
)

type fakeDatabase struct {
	users map[string]models.User
	rooms map[string]models.Room
	posts map[string]models.Post
}

func FakeDatabaseRepo() (Database, error) {
	users := make(map[string]models.User)
	rooms := make(map[string]models.Room)
	posts := make(map[string]models.Post)

	return &fakeDatabase{users, rooms, posts}, nil
}

func (db *fakeDatabase) Close() error {
	return nil
}

func (db *fakeDatabase) userGetName(userId string) (string, error) {
	user, ok := db.users[userId]
	if !ok {
		return "", errors.New("user not found")
	}

	return user.Username, nil
}

func (db *fakeDatabase) UserGetId(username string) (string, error) {
	for _, user := range db.users {
		if user.Username == username {
			return user.Id, nil
		}
	}

	return "", errors.New("user not found")
}

func (db *fakeDatabase) UserLogin(username string, password string) (models.User, error) {
	userId, err := db.UserGetId(username)
	if err != nil {
		return models.User{}, err
	}

	user, ok := db.users[userId]
	if !ok {
		return models.User{}, errors.New("user not found")
	}

	if !utils.CheckPasswordHash(password, user.Pass) {
		return models.User{}, errors.New("bad username or password")
	}

	return user, nil
}

func (db *fakeDatabase) UserRegister(user models.User) error {
	db.users[user.Id] = user
	return nil
}

func (db *fakeDatabase) ListRooms() ([]models.Room, error) {
	rooms := []models.Room{}
	for _, room := range db.rooms {
		rooms = append(rooms, room)
	}

	return rooms, nil
}

func (db *fakeDatabase) CreateRoom(room models.Room) error {
	db.rooms[room.Id] = room
	return nil
}

func (db *fakeDatabase) GetRoom(roomId string) (models.Room, error) {
	room, ok := db.rooms[roomId]
	if !ok {
		return models.Room{}, errors.New("room not found")
	}

	return room, nil
}

func (db *fakeDatabase) ListPosts(roomId string) ([]models.PostView, error) {
	var posts []models.Post
	for _, post := range db.posts {
		if post.RoomId == roomId {
			posts = append(posts, post)
		}
	}

	var postViews []models.PostView
	for _, post := range posts {
		username, err := db.userGetName(post.UserId)
		if err != nil {
			return nil, err
		}

		postView := models.PostView{
			Id:       post.Id,
			Content:  post.Content,
			RoomId:   post.RoomId,
			Username: username,
		}

		postViews = append(postViews, postView)
	}

	return postViews, nil
}

func (db *fakeDatabase) CreatePost(post models.Post) (models.PostView, error) {
	db.posts[post.Id] = post
	postView := models.PostView{
		Id:      post.Id,
		Content: post.Content,
		RoomId:  post.RoomId,
	}

	username, err := db.userGetName(post.UserId)
	if err != nil {
		return models.PostView{}, err
	}

	postView.Username = username
	return postView, nil
}
