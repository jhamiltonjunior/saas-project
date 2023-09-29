package service

import (
    "my-saas-app/internal/domain/entities"
    "my-saas-app/internal/domain/repositories"
)

type UserService struct {
    userRepository repositories.UserRepository
}

func NewUserService(userRepository repositories.UserRepository) *UserService {
    return &UserService{userRepository: userRepository}
}

func (s *UserService) GetUserByID(id int) (*entities.User, error) {
    return s.userRepository.FindByID(id)
}

func (s *UserService) CreateUser(user *entities.User) error {
    return s.userRepository.Create(user)
}

func (s *UserService) UpdateUser(user *entities.User) error {
    return s.userRepository.Update(user)
}

func (s *UserService) DeleteUser(id int) error {
    return s.userRepository.Delete(id)
}