package data

import (
	"airbnb/features/user"
	"gorm.io/gorm"
)

func UserCoreToUser(userCore user.Core) User {
	return User{
		Model:          gorm.Model{ID: userCore.ID},
		Name:           userCore.Name,
		Email:          userCore.Email,
		Password:       userCore.Password,
		Role:           userCore.Role,
		ProfilePicture: userCore.ProfilePicture,
		Phone:          userCore.Phone,
		Address:        userCore.Address,
	}
}

func UserToUserCore(users User) user.Core {

	return user.Core{
		ID:             users.ID,
		Name:           users.Name,
		Email:          users.Email,
		Password:       users.Password,
		Role:           users.Role,
		ProfilePicture: users.ProfilePicture,
		Phone:          users.Phone,
		Address:        users.Address,
	}
}

func ListUserToUserCore(users []User) []user.Core {
	var userCore []user.Core
	for _, v := range users {
		userCore = append(userCore, UserToUserCore(v))
	}
	return userCore
}