package database

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
	"giskard_api/common"
)

var dbBoConnections = make(map[string]*gorm.DB)

func GetDb(clientName string) (*gorm.DB, error) {

	fmt.Printf("Requested client DB : %s", clientName)

	if database, ok := dbBoConnections[clientName]; ok {
		return database, nil
	} else {
		fmt.Printf("Establishing new DB connection for the client %s\n", clientName)
		DB, err := connectToBoDatabase(clientName)
		if err != nil {
			return nil, err
		}
		dbBoConnections[clientName] = DB
		return dbBoConnections[clientName], nil
	}
}

func connectToBoDatabase(client string) (*gorm.DB, error) {

	parameters, err := common.GetDataBaseConfig(client)
	fmt.Println("******* params : ", parameters)
	if err != nil {
		return nil, err
	}
	db, err := gorm.Open(parameters.DbDriver, fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", parameters.DbUser, parameters.DbPassword, parameters.DbHost, parameters.DbPort, parameters.DatabaseName))
	log.Println("-----------------*************d b************-------------------", db)
	if err != nil {
		fmt.Println("failed to connect  to db client ", client, ":", parameters.DatabaseName)
		return nil, err
	}
	fmt.Println("connect to db client ", client, ":", parameters.DatabaseName)
	db.LogMode(true)
	return db, err
}