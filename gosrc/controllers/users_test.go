package controllers

import (
	"chatrooms/gosrc/repositories/database"
	"testing"
)

func TestUsercontroller(t *testing.T) {
	botUsername := "bot"
	dbRepo := database.FakeDatabaseRepo()
	userController := NewUserController(dbRepo, botUsername)

	// Happy cases

	_, err := userController.Register(UserRequest{
		Username: "user1",
		Password: "123456",
	})
	if err != nil {
		t.Fatalf("Error when user register: %v", err)
	}

	_, err = userController.Login(UserRequest{
		Username: "user1",
		Password: "123456",
	})
	if err != nil {
		t.Fatalf("Error when user login: %v", err)
	}

	// Bad cases - Login

	_, err = userController.Login(UserRequest{
		Username: botUsername,
		Password: "123456",
	})
	if err == nil {
		t.Fatalf("Expected error when bot login, got nil")
	}

	_, err = userController.Login(UserRequest{
		Username: "user1",
		Password: "1234567",
	})
	if err == nil {
		t.Fatalf("Expected error when user login with wrong password, got nil")
	}

	// Bad cases - Register

	_, err = userController.Register(UserRequest{
		Username: botUsername,
		Password: "123456",
	})
	if err == nil {
		t.Fatalf("Expected error when bot register, got nil")
	}

	_, err = userController.Register(UserRequest{
		Username: "user1",
		Password: "123456",
	})
	if err == nil {
		t.Fatalf("Expected error when user register with duplicate username, got nil")
	}
}
