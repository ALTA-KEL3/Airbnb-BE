package user

import (
	"mime/multipart"

	"github.com/labstack/echo/v4"
)

type Core struct {
	ID             uint
	Name           string
	Email          string `validate:"required,email,unique"`
	Password       string `validate:"required"`
	Role           string `validate:"required"`
	ProfilePicture string
	Phone          string
	Address        string
	// Homestay []Homestay
}

// type Homestay struct {
// 	ID              uint
// 	Name            string
// 	HomestayAddress string
// 	Phone           string
// 	Price           int
// 	Facility        string
// 	Image1          string
// }

type UserHandler interface {
	Register() echo.HandlerFunc
	Login() echo.HandlerFunc
	Profile() echo.HandlerFunc
	Update() echo.HandlerFunc
	Delete() echo.HandlerFunc
}

type UserService interface {
	Register(newUser Core) (Core, error)
	Login(email, password string) (string, Core, error)
	Profile(token interface{}) (Core, error)
	Update(token interface{}, fileData multipart.FileHeader, updateData Core) (Core, error)
	Delete(token interface{}) error
}

type UserData interface {
	Register(newUser Core) (Core, error)
	Login(email string) (Core, error)
	Profile(userID uint) (Core, error)
	Update(updateData Core, userID uint) (Core, error)
	Delete(userID uint) error
}
