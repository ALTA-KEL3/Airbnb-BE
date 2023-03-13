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
func (hq *homestayQuery) Add(userRole string, userID uint, newHomestay homestay.Core) (homestay.Core, error) {
	if userRole != "hoster" {
		log.Println("access denied, cannot add product because you are not hoster")
		return homestay.Core{}, errors.New("access denied, cannot add product because you are not hoster")
	}
	cnv := CoreToModel(newHomestay)
	err := hq.db.Create(&cnv).Error
	if err != nil {
		log.Println("query error", err.Error())
		return homestay.Core{}, errors.New("server error")
	}

	newHomestay.ID = cnv.ID
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
func (*homestayQuery) ShowAll() ([]homestay.Core, error) {
	panic("unimplemented")
}

// ShowDetail implements homestay.HomestayData
func (*homestayQuery) ShowDetail(homestayID uint) (homestay.Core, error) {
	panic("unimplemented")
}

// Update implements homestay.HomestayData
func (*homestayQuery) Update(userID uint, homestayID uint, updateHomestay homestay.Core) (homestay.Core, error) {
	panic("unimplemented")
}
