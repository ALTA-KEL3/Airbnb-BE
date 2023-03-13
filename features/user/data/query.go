package data

import (
	"airbnb/features/user"
	"errors"
	"log"

	"gorm.io/gorm"
)

type userQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) user.UserData {
	return &userQuery{
		db: db,
	}
}

func (q *userQuery) Register(newUser user.Core) (user.Core, error) {
	dupEmail := UserCoreToUser(newUser)
	err := q.db.Where("email = ?", newUser.Email).First(&dupEmail).Error
	if err == nil {
		log.Println("duplicated")
		return user.Core{}, errors.New("email duplicated")
	}

	newUser.ProfilePicture = "https://cdn.pixabay.com/photo/2015/10/05/22/37/blank-profile-picture-973460_1280.png"

	users := UserCoreToUser(newUser)
	err = q.db.Create(&users).Error
	if err != nil {
		log.Println("query error", err.Error())
		return user.Core{}, errors.New("server error")
	}

	newUser.ID = users.ID
	return newUser, nil
}

func (q *userQuery) Login(username string) (user.Core, error) {
	if username == "" {
		log.Println("data empty, query error")
		return user.Core{}, errors.New("username not allowed empty")
	}
	res := User{}
	if err := q.db.Where("username = ?", username).First(&res).Error; err != nil {
		log.Println("login query error", err.Error())
		return user.Core{}, errors.New("data not found")
	}

	return UserToUserCore(res), nil
}

func (q *userQuery) Profile(userID uint) (user.Core, error) {
	res := User{}
	err := q.db.Where("id = ?", userID).First(&res).Error
	if err != nil {
		log.Println("query err", err.Error())
		return user.Core{}, errors.New("account not found")
	}
	return UserToUserCore(res), nil
}

func (q *userQuery) Update(updateData user.Core, userID uint) (user.Core, error) {
	if updateData.Email != "" {
		dupEmail := User{}
		err := q.db.Where("email = ?", updateData.Email).First(&dupEmail).Error
		if err == nil {
			log.Println("duplicated")
			return user.Core{}, errors.New("email duplicated")
		}
	}

	users := UserCoreToUser(updateData)
	qry := q.db.Model(&User{}).Where("id = ?", userID).Updates(&users)
	affrows := qry.RowsAffected
	if affrows == 0 {
		log.Println("no rows affected")
		return user.Core{}, errors.New("no data updated")
	}
	err := qry.Error
	if err != nil {
		log.Println("update user query error", err.Error())
		return user.Core{}, errors.New("user not found")
	}
	result := UserToUserCore(users)
	result.ID = userID
	return result, nil
}

func (q *userQuery) Delete(userID uint) error {
	getID := User{}
	err := q.db.Where("id = ?", userID).First(&getID).Error
	if err != nil {
		log.Println("get user error : ", err.Error())
		return errors.New("failed to get user data")
	}

	if getID.ID != userID {
		log.Println("unauthorized request")
		return errors.New("unauthorized request")
	}
	qryDelete := q.db.Delete(&User{}, userID)
	affRow := qryDelete.RowsAffected

	if affRow <= 0 {
		log.Println("No rows affected")
		return errors.New("failed to delete user content, data not found")
	}

	return nil
}
