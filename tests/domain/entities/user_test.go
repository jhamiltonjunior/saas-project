package entities_test

import (
	"my-saas-app/src/domain/entities"
	"testing"
)

func TestUser(t *testing.T) {
	user := entities.User{
		ID:       1,
		Name:     "John Doe",
		Email:    "johndoe@example.com",
		Password: "password",
	}

	t.Run("TestUserFields", func(t *testing.T) {
		if user.ID != 1 {
			t.Errorf("Expected ID to be 1, but got %d", user.ID)
		}

		if user.Name != "John Doe" {
			t.Errorf("Expected Name to be 'John Doe', but got '%s'", user.Name)
		}

		if user.Email != "johndoe@example.com" {
			t.Errorf("Expected Email to be 'johndoe@example.com', but got '%s'", user.Email)
		}

		if user.Password != "password" {
			t.Errorf("Expected Password to be 'password', but got '%s'", user.Password)
		}
	})
}
