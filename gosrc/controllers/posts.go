package controllers

import (
	"chatrooms/gosrc/models"
	"chatrooms/gosrc/repositories/broker"
	"chatrooms/gosrc/repositories/database"

	"github.com/google/uuid"
)

type PostsController struct {
	dbRepo     database.Database
	brokerRepo broker.Broker
}

func NewPostsController(dbRepo database.Database, brokerRepo broker.Broker) *PostsController {
	return &PostsController{dbRepo, brokerRepo}
}

func (pc *PostsController) ListPosts(roomId string) ([]models.PostView, error) {
	posts, err := pc.dbRepo.ListPosts(roomId)
	if err != nil {
		return nil, err
	}

	return posts, nil
}

func (pc *PostsController) CreatePost(userId, roomId, content string) error {
	post := models.Post{
		Id:      uuid.New().String(),
		UserId:  userId,
		RoomId:  roomId,
		Content: content,
	}

	postview, err := pc.dbRepo.CreatePost(post)
	if err != nil {
		return err
	}

	return pc.brokerRepo.Publish(roomId, postview)
}

type MessageWriter func(data []byte) error

func (pc *PostsController) SubscribeMessages(roomId string, writer MessageWriter) error {
	subscription, err := pc.brokerRepo.Subscribe(roomId)
	if err != nil {
		return err
	}
	defer subscription.Close()

	for msg := range subscription.Channel() {
		body := msg.Body
		if err := writer(body); err != nil {
			return err
		}
	}

	return nil
}
