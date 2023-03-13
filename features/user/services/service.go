package services

import (
	"airbnb/features/user"
	"airbnb/helper"
	"errors"
	"mime/multipart"
	"strings"
)

type userService struct {
	qry user.UserData
}

func New(ud user.UserData) user.UserService {
	return &userService{
		qry: ud,
	}
}

// Delete implements user.UserService
func (*userService) Delete(token interface{}) error {
	panic("unimplemented")
}

// Login implements user.UserService
func (*userService) Login(email string, password string) (string, user.Core, error) {
	panic("unimplemented")
}

// Profile implements user.UserService
func (*userService) Profile(token interface{}) (user.Core, error) {
	panic("unimplemented")
}

// Register implements user.UserService
func (us *userService) Register(newUser user.Core) (user.Core, error) {
	hashed := helper.GeneratePassword(newUser.Password)
	newUser.Password = string(hashed)

	res, err := us.qry.Register(newUser)
	if err != nil {
		msg := ""
		if strings.Contains(err.Error(), "duplicated") {
			msg = "email already registered"
		} else {
			msg = "server error"
		}
		return user.Core{}, errors.New(msg)
	}

	return res, nil
}

// Update implements user.UserService
func (*userService) Update(token interface{}, fileData multipart.FileHeader, updateData user.Core) (user.Core, error) {
	panic("unimplemented")
}
