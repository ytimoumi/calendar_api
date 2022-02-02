package services

import (
	"giskard_api/database"
	"giskard_api/models"
	"giskard_api/repositories"
	"log"
)

func GetCalendarInfo(clientName string) (response []*models.Calendar, err error) {
	// here we will get all available slots
	var (
		er error
	)
	db, err := database.GetDb(clientName)
	if err != nil {
		log.Println(err)
		return response, err
	}
	response, er = repositories.GetCalendar(db)
	if er != nil {
		log.Println(er)
		return response, er
	}
	return response, nil
}

func CreateReservation(clientName string, idAvailability int, input models.ToReserve) (response *models.Response, err error) {
	// here we will create reservation
	var (
		er           error
		availability models.Calendar
		errAv        error
	)
	db, err := database.GetDb(clientName)
	if err != nil {
		log.Println(err)
		return response, err
	}
	//get availability by id
	availability, errAv = repositories.GetAvailabilityById(db, idAvailability)
	if errAv != nil {
		log.Println(errAv)
		return response, errAv
	}
	startDateAvailability := availability.Start
	endDateAvailability := availability.End

	if input.Start.Before(startDateAvailability) || input.End.After(endDateAvailability) {
		response = models.NewResponse(400, "You can't reserve in this slot, please choose another !!")
		return response, nil
	}

	//create object
	newReservation := models.NewToReserve(input.Start, input.End, input.Title, input.Email)
	// insert in database
	_, er = repositories.InsertReserv(db, newReservation)
	if er != nil {
		log.Println(er)
		return response, er
	}

	// if start date of reservation is After than the start date of availability
	// AND the end date of reservation id Before than the end date of availability
	// we will create another availability with the rest of time.
	restTime := endDateAvailability.Sub(input.End)
	newEndDate := input.End.Add(restTime)
	newStartDate := input.End
	sub := newEndDate.Sub(newStartDate)

	log.Println("--------------sub--------------", sub.Minutes())
	// I will choose 30 minutes as a min time for availability
	if sub.Minutes() >= 30 {
		NewAvailability := models.NewAvailability(newStartDate, newEndDate)
		_, err = CreateAvailability(clientName, NewAvailability)
		if err != nil {
			log.Println(err)
			response = models.NewResponse(400, err.Error())
		}
	}

	return response, err
}

func CreateAvailability(clientName string, input models.Availability) (response *models.Response, err error) {
	// here we will create availability
	var (
		er error
	)
	db, err := database.GetDb(clientName)
	if err != nil {
		log.Println(err)
		return response, err
	}
	//create object
	newAvailability := models.NewAvailability(input.Start, input.End)
	// insert in database
	_, er = repositories.InsertAvailability(db, newAvailability)
	if er != nil {
		log.Println(er)
		return response, er
	}

	return response, err
}

func DeleteReservation(clientName string, id int, email string) (response *models.Response, err error) {
	var (
		er error
	)
	db, err := database.GetDb(clientName)
	if err != nil {
		log.Println(err)
		return response, err
	}
	// insert in database
	_, er = repositories.DeleteReservation(db, id, email)
	if er != nil {
		log.Println(er)
		return response, er
	}

	return response, err
}

func DeleteAvailability(clientName string, id int) (response *models.Response, err error) {
	var (
		er error
	)
	db, err := database.GetDb(clientName)
	if err != nil {
		log.Println(err)
		return response, err
	}
	// insert in database
	_, er = repositories.DeleteAvailability(db, id)
	if er != nil {
		log.Println(er)
		return response, er
	}

	return response, err
}
