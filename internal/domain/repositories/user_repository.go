package repositories


import "my-saas-app/internal/domain/entities"


type UserRepository interface {
    FindByID(id int) (*entities.User, error)
    FindByEmail(email string) (*entities.User, error)
    Create(user *entities.User) (int, error)
    Update(user *entities.User) error
    Delete(id int) error
}