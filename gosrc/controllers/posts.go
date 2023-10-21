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

func (pc *PostsController) createCommand(roomId, payload string) error {
	commandView := models.CommandView{
		Id:      uuid.New().String(),
		Payload: payload,
		RoomId:  roomId,
	}

	return pc.brokerRepo.PublishCommand(commandView)
}

func (pc *PostsController) ListPosts(roomId string) ([]models.PostView, error) {
	posts, err := pc.dbRepo.ListPosts(roomId)
	if err != nil {
		return nil, err
	}

	return posts, nil
}

func (pc *PostsController) CreatePost(userId, roomId, content string) error {
	if models.IsCommand(content) {
		return pc.createCommand(roomId, content)
	}

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

	return pc.brokerRepo.PublishPost(roomId, postview)
}

func (pc *PostsController) SubscribePosts(roomId string) (broker.PostsSubscription, error) {
	return pc.brokerRepo.SubscribePosts(roomId)
}
