package services

import (
	"airbnb/app/config"
	"airbnb/features/user"
	"airbnb/helper"
	"errors"
	"log"
	"mime/multipart"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
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
func (us *userService) Delete(token interface{}) error {
	id := helper.ExtractToken(token)
	err := us.qry.Delete(uint(id))
	if err != nil {
		log.Println("query error", err.Error())
		return errors.New("query error, delete account fail")
	}
	return nil
}

// Login implements user.UserService
func (us *userService) Login(email string, password string) (string, user.Core, error) {
	res, err := us.qry.Login(email)
	if err != nil {
		msg := ""
		if strings.Contains(err.Error(), "not found") {
			msg = "data not found"
		} else {
			msg = "there is a problem with the server"
		}
		return "", user.Core{}, errors.New(msg)
	}

	if err := helper.ComparePassword(res.Password, password); err != nil {
		log.Println("login compare", err.Error())
		return "", user.Core{}, errors.New("password not matched")
	}

	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["userID"] = res.ID
	claims["exp"] = time.Now().Add(time.Hour * 48).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	useToken, _ := token.SignedString([]byte(config.JWTKey))

	return useToken, res, nil
}

// Profile implements user.UserService
func (us *userService) Profile(token interface{}) (user.Core, error) {
	userID := helper.ExtractToken(token)
	res, err := us.qry.Profile(uint(userID))
	if err != nil {
		log.Println("data not found")
		return user.Core{}, errors.New("query error, problem with server")
	}
	return res, nil
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
func (us *userService) Update(token interface{}, fileData multipart.FileHeader, updateData user.Core) (user.Core, error) {
	userID := helper.ExtractToken(token)

	hashed := helper.GeneratePassword(updateData.Password)
	updateData.Password = string(hashed)
	log.Println("size:", fileData.Size)

	url, err := helper.GetUrlImagesFromAWS(fileData, int(1))
	if err != nil {
		return user.Core{}, errors.New("validate: " + err.Error())
	}
	updateData.ProfilePicture = url
	res, err := us.qry.Update(uint(userID), updateData)
	if err != nil {
		msg := ""
		if strings.Contains(err.Error(), "not found") {
			msg = "account not registered"
		} else if strings.Contains(err.Error(), "email") {
			msg = "email duplicated"
		} else if strings.Contains(err.Error(), "access denied") {
			msg = "access denied"
		}
		return user.Core{}, errors.New(msg)
	}
	return res, nil
}
