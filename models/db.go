package models


type DatabaseParameters struct {
	DbUser       string
	DbPassword   string
	DbHost       string
	DbPort       int
	DatabaseName string
	DbDriver     string
}
