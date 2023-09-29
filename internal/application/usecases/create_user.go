package usecases

import (
    "my-saas-app/internal/domain/entities"
    "my-saas-app/internal/domain/repositories"
)

type CreateUserInput struct {
    Name     string
    Email    string
    Password string
}

type CreateUserOutput struct {
    User *entities.User
}

type CreateUserUseCase struct {
    userRepository repositories.UserRepository
}

func NewCreateUserUseCase(userRepository repositories.UserRepository) *CreateUserUseCase {
    return &CreateUserUseCase{userRepository: userRepository}
}

func (uc *CreateUserUseCase) Execute(input *CreateUserInput) (*CreateUserOutput, error) {
    user := &entities.User{
        Name:     input.Name,
        Email:    input.Email,
        Password: input.Password,
    }

    if err := uc.userRepository.Create(user); err != nil {
        return nil, err
    }

    return &CreateUserOutput{User: user}, nil
}