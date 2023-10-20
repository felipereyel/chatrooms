package controllers

import (
	"chatrooms/gosrc/models"
	"chatrooms/gosrc/repositories/database"

	"github.com/google/uuid"
)

type PostsController struct {
	DbRepo database.Database
}

func NewPostsController(dbRepo database.Database) *PostsController {
	return &PostsController{dbRepo}
}

func (tc *PostsController) ListPosts(roomId string) ([]models.PostView, error) {
	posts, err := tc.DbRepo.ListPosts(roomId)
	if err != nil {
		return nil, err
	}

	return posts, nil
}

func (tc *PostsController) CreatePost(userId, roomId, content string) (models.PostView, error) {
	post := models.Post{
		Id:      uuid.New().String(),
		UserId:  userId,
		RoomId:  roomId,
		Content: content,
	}

	postview, err := tc.DbRepo.CreatePost(post)
	if err != nil {
		return models.PostView{}, err
	}

	return postview, nil
}
