package entities

import (
	"errors"
)


func (u *User) Validate(user *User) (User, error) {

	if user.Email == "" {
		return User{}, errors.New("email is invalid")
	}

	if user.Password == "" {
		return User{}, errors.New("password is invalid")
	}

	userValidated := User{
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	}

	return userValidated, nil
}

// func (u *User) ValidateUpdate(user *User) (User, error) {
// 	if nameValidate(user.Name) {
// 		return User{}, errors.New("name is required")
// 	}

// 	if (user.Email == "") {
// 		return User{}, errors.New("email is required")
// 	}

// 	userValidated := User{
// 		Name:     user.Name,
// 		Email:    user.Email,
// 		Password: user.Password,
// 	}

// 	return userValidated, nil
// }

func (u *User) ValidateLogin(user *User) (User, error) {
	if (user.Email == "") {
		return User{}, errors.New("email is required")
	}

	if (user.Password == "") {
		return User{}, errors.New("password is required")
	}

	userValidated := User{
		Email:    user.Email,
		Password: user.Password,
	}

	return userValidated, nil
}

