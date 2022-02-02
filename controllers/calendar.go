package controllers

import (
	"errors"
	"fmt"
	"giskard_api/common"
	"giskard_api/common/logger"
	"giskard_api/models"
	"giskard_api/services"
	"github.com/gin-gonic/gin"
)

func GetCalendar(gc *gin.Context) (response []*models.Calendar, err error) {
	clientName, err := common.GetClientName(gc)
	if err != nil {
		logger.Log("internal", gc.ClientIP(), errors.New(fmt.Sprintf("[GetCalendar] Failed to get client name : %s", err)), true)
		return response, err
	}
	response, err = services.GetCalendarInfo(clientName)
	if err != nil {
		logger.Log(clientName, gc.ClientIP(), errors.New(fmt.Sprintf("[GetCalendar] Failed to retrieve calendar info : %s", err)), true)
		return response, err
	}
	return response, nil
}
func CreateReservation(gc *gin.Context,idAvailability int, input models.ToReserve) (response *models.Response, err error) {
	clientName, err := common.GetClientName(gc)
	if err != nil {
		logger.Log("internal", gc.ClientIP(), errors.New(fmt.Sprintf("[CreateReservation] Failed to get client name : %s", err)), true)
		response = models.NewResponse(400, err.Error())
		return response, err
	}
	response, err = services.CreateReservation(clientName, idAvailability, input)
	if err != nil {
		logger.Log(clientName, gc.ClientIP(), errors.New(fmt.Sprintf("[CreateReservation] Failed to retrieve calendar info : %s", err)), true)
		response = models.NewResponse(400, err.Error())
		return response, err
	}
	response = models.NewResponse(200, "Reservation created")
	return response, nil
}

func CreateAvailability(gc *gin.Context, input models.Availability) (response *models.Response, err error) {
	clientName, err := common.GetClientName(gc)
	if err != nil {
		logger.Log("internal", gc.ClientIP(), errors.New(fmt.Sprintf("[CreateAvailability] Failed to get client name : %s", err)), true)
		response = models.NewResponse(400, err.Error())
		return response, err
	}
	response, err = services.CreateAvailability(clientName, input)
	if err != nil {
		logger.Log(clientName, gc.ClientIP(), errors.New(fmt.Sprintf("[CreateAvailability] Failed to retrieve calendar info : %s", err)), true)
		response = models.NewResponse(400, err.Error())
		return response, err
	}
	response = models.NewResponse(200, "Availability created")
	return response, nil
}

func DeleteReservation(gc *gin.Context, id int, email string) (response *models.Response, err error) {
	clientName, err := common.GetClientName(gc)
	if err != nil {
		logger.Log("internal", gc.ClientIP(), errors.New(fmt.Sprintf("[DeleteReservation] Failed to get client name : %s", err)), true)
		response = models.NewResponse(400, err.Error())
		return response, err
	}
	response, err = services.DeleteReservation(clientName, id, email)
	if err != nil {
		logger.Log(clientName, gc.ClientIP(), errors.New(fmt.Sprintf("[DeleteReservation] Failed to retrieve calendar info : %s", err)), true)
		response = models.NewResponse(400, err.Error())
		return response, err
	}
	response = models.NewResponse(200, "Reservation deleted")
	return response, nil
}

func DeleteAvailability(gc *gin.Context, id int) (response *models.Response, err error) {
	clientName, err := common.GetClientName(gc)
	if err != nil {
		logger.Log("internal", gc.ClientIP(), errors.New(fmt.Sprintf("[DeleteAvailability] Failed to get client name : %s", err)), true)
		response = models.NewResponse(400, err.Error())
		return response, err
	}
	response, err = services.DeleteAvailability(clientName, id)
	if err != nil {
		logger.Log(clientName, gc.ClientIP(), errors.New(fmt.Sprintf("[DeleteAvailability] Failed to retrieve calendar info : %s", err)), true)
		response = models.NewResponse(400, err.Error())
		return response, err
	}
	response = models.NewResponse(200, "Availability deleted")
	return response, nil
}
