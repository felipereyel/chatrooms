package controllers

import (
	"chatrooms/gosrc/models"
	"chatrooms/gosrc/repositories/broker"
	"chatrooms/gosrc/repositories/database"
	"testing"
	"time"
)

func setup(t *testing.T) *PostsController {
	dbRepo := database.FakeDatabaseRepo()
	brokerRepo := broker.FakeBrokerRepo(3)
	postsController := NewPostsController(dbRepo, brokerRepo)

	room := models.Room{
		Id:   "room1",
		Name: "room1",
	}

	err := dbRepo.CreateRoom(room)
	if err != nil {
		t.Fatalf("[Setup] Error when create room: %v", err)
	}

	user := models.User{
		Id:       "user1",
		Username: "user1",
	}

	err = dbRepo.UserRegister(user)
	if err != nil {
		t.Fatalf("[Setup] Error when register user: %v", err)
	}

	return postsController
}

func TestPostscontroller(t *testing.T) {
	postsController := setup(t)

	// posting

	if err := postsController.CreatePost("user1", "room1", "Hello world"); err != nil {
		t.Fatalf("Error when user create post: %v", err)
	}

	posts, err := postsController.ListPosts("room1")
	if err != nil {
		t.Fatalf("Error when list posts: %v", err)
	}

	if len(posts) != 1 {
		t.Fatalf("Expected 1 post, got %d", len(posts))
	}

	// subscription

	subs, err := postsController.SubscribePosts("room1")
	if err != nil {
		subs.Close()
		t.Fatalf("Error when subscribe messages: %v", err)
	}
	defer subs.Close()

	select {
	case <-subs.Channel():
		// ok
	case <-time.After(2 * time.Second):
		t.Fatal("User subscribe messages timeout")
	}
}
