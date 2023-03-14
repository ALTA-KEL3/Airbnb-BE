package data

import (
	"airbnb/features/homestay"
	"errors"
	"log"

	"gorm.io/gorm"
)

type homestayQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) homestay.HomestayData {
	return &homestayQuery{
		db: db,
	}
}

// Add implements homestay.HomestayData
func (hq *homestayQuery) Add(userID uint, newHomestay homestay.Core) (homestay.Core, error) {
	// if user.Role != "hoster" {
	// 	log.Println("access denied, cannot add product because you are not hoster")
	// 	return homestay.Core{}, errors.New("access denied, cannot add product because you are not hoster")
	// }
	cnv := CoreToModel(newHomestay)
	cnv.UserID = uint(userID)
	err := hq.db.Create(&cnv).Error
	if err != nil {
		log.Println("query error", err.Error())
		return homestay.Core{}, errors.New("server error")
	}
	newHomestay.ID = cnv.ID
	// newHomestay.Role = cnv.Role

	return newHomestay, nil

}

// Delete implements homestay.HomestayData
func (hq *homestayQuery) Delete(userID uint, homestayID uint) error {
	getID := Homestay{}
	err := hq.db.Where("id = ?", homestayID).First(&getID).Error
	if err != nil {
		log.Println("get homestay error : ", err.Error())
		return errors.New("failed to get homestay data")
	}

	if getID.ID != homestayID {
		log.Println("unauthorized request")
		return errors.New("unauthorized request")
	}
	qryDelete := hq.db.Delete(&Homestay{}, homestayID)
	affRow := qryDelete.RowsAffected

	if affRow <= 0 {
		log.Println("No rows affected")
		return errors.New("failed to delete homestay content, data not found")
	}

	return nil
}

// ShowAll implements homestay.HomestayData
func (hq *homestayQuery) ShowAll() ([]homestay.Core, error) {
	res := []Homestay{}
	err := hq.db.Find(&res).Error
	if err != nil {
		log.Println("data not found", err.Error())
		return []homestay.Core{}, errors.New("data not found")
	}
	result := []homestay.Core{}
	for _, val := range res {
		result = append(result, ModelToCore(val))

	}
	return result, nil
}

// ShowDetail implements homestay.HomestayData
func (hq *homestayQuery) ShowDetail(homestayID uint) (homestay.Core, error) {
	res := Homestay{}

	err := hq.db.Where("id = ?", homestayID).First(&res).Error
	if err != nil {
		log.Println("data not found", err.Error())
		return homestay.Core{}, errors.New("data not found")
	}

	result := ModelToCore(res)

	return result, nil
}

// Update implements homestay.HomestayData
func (hq *homestayQuery) Update(userID uint, homestayID uint, updateHomestay homestay.Core) (homestay.Core, error) {
	cnv := CoreToModel(updateHomestay)
	cnv.ID = uint(homestayID)

	qry := hq.db.Where("id = ?", homestayID).Updates(&cnv)
	affrows := qry.RowsAffected
	if affrows == 0 {
		log.Println("no rows affected")
		return homestay.Core{}, errors.New("no data updated")
	}
	err := qry.Error
	if err != nil {
		log.Println("update homestay query error", err.Error())
		return homestay.Core{}, errors.New("user not found")
	}
	return updateHomestay, nil
}
