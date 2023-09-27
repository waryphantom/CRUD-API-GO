package main

import (
	"main/controllers"
	"main/initialize"

	"github.com/gin-gonic/gin"
)

func init() {
	initialize.LoadEnvVariables()
	initialize.ConnectToDB()
}

func main() {
	server := gin.Default()
	server.POST("/bookticket", controllers.BookTicket)
	server.GET("/bookedusers", controllers.BookedUsers)
	server.GET("/userdetails/:email", controllers.UserDetails)
	server.GET("/cancelticket/:email", controllers.CancelTicket)
	server.Run(":3000")
}
