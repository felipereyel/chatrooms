package controllers

import (
	"chatrooms/gosrc/models"
	"chatrooms/gosrc/repositories/broker"
	"chatrooms/gosrc/repositories/database"

	"github.com/google/uuid"
)

type PostsController struct {
	DbRepo     database.Database
	BrokerRepo broker.Broker
}

func NewPostsController(dbRepo database.Database, brokerRepo broker.Broker) *PostsController {
	return &PostsController{dbRepo, brokerRepo}
}

func (tc *PostsController) ListPosts(roomId string) ([]models.PostView, error) {
	posts, err := tc.DbRepo.ListPosts(roomId)
	if err != nil {
		return nil, err
	}

	return posts, nil
}

func (tc *PostsController) CreatePost(userId, roomId, content string) error {
	post := models.Post{
		Id:      uuid.New().String(),
		UserId:  userId,
		RoomId:  roomId,
		Content: content,
	}

	postview, err := tc.DbRepo.CreatePost(post)
	if err != nil {
		return err
	}

	return tc.BrokerRepo.Publish(roomId, postview)
}
