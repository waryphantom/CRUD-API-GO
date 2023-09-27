package main

import (
	"main/initialize"
	"main/models"
)

func init() {
	initialize.LoadEnvVariables()
	initialize.ConnectToDB()
}

func main() {
	initialize.DB.AutoMigrate(&models.User{})
}
