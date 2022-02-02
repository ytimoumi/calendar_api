package repositories

import (
	"giskard_api/models"
	"github.com/jinzhu/gorm"
	"log"
)

func InsertReserv(db *gorm.DB, input models.ToReserve) (response models.Response, err error) {
	if er := db.Table("reservations").Create(&input).Error; er != nil {
		log.Println(er)
		return response, er
	}
	return response, nil
}

func InsertAvailability(db *gorm.DB, input models.Availability) (response models.Response, err error) {
	if er := db.Table("availabilities").Create(&input).Error; er != nil {
		log.Println(er)
		return response, er
	}
	return response, nil
}

func DeleteReservation(db *gorm.DB, id int, email string) (response models.Response, err error) {
	var input models.ToReserve
	if er := db.Table("reservations").Where("id = ? AND email = ?", id, email).Delete(&input).Error; er != nil {
		log.Println(er)
		return response, er
	}
	return response, nil
}

func GetCalendar(db *gorm.DB) ([]*models.Calendar, error) {
	var (
		availabilities []*models.Calendar
		err            error
	)

	err = db.Table("availabilities").Find(&availabilities).Error
	if err != nil {
		availabilities = make([]*models.Calendar, 0)
	}

	return availabilities, nil

}

func DeleteAvailability(db *gorm.DB, id int) (response models.Response, err error) {
	var input models.ToReserve
	if er := db.Table("availabilities").Where("id = ?", id).Delete(&input).Error; er != nil {
		log.Println(er)
		return response, er
	}
	return response, nil
}

func GetAvailabilityById(db *gorm.DB, id int) (models.Calendar, error) {
	var (
		availability models.Calendar
		err            error
	)
	err = db.Table("availabilities").Where("id = ?", id).Find(&availability).Error
	if err != nil {
		return availability, err
	}

	return availability, nil

}