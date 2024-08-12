package usecases

import (
	"errors"
	"my-saas-app/src/domain/entities"
	"my-saas-app/src/domain/repositories"
)

type CreateUserInput struct {
	Name     string
	Email    string
	Password string
}

type CreateUserOutput struct {
	User *entities.User
}

type UserUseCase struct {
	userRepository repositories.UserRepository
}

func NewUserUseCase(userRepository repositories.UserRepository) *UserUseCase {
	return &UserUseCase{userRepository: userRepository}
}

func (uc *UserUseCase) Create(input *entities.User) (int, error) {
	if nameInvalid(input.Name) {
		return 0, errors.New("name is invalid")
	}

	user := &entities.User{
		Name:     input.Name,
		Email:    input.Email,
		Password: input.Password,
	}

	userId, err := uc.userRepository.Create(user)

	if err != nil {
		return 0, err
	}

	return userId, nil
}

func (uc *UserUseCase) GetUserByID(id int) (*entities.User, error) {
	return uc.userRepository.FindByID(id)
}
