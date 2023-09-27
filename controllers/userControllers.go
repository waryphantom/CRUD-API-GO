package controllers

import (
	"strings"

	"main/initialize"
	"main/models"

	"github.com/gin-gonic/gin"
)

var TotalTicket uint = 100

func BookTicket(c *gin.Context) {
	var body struct {
		FirstName   string
		LastName    string
		Email       string
		UserTickets uint
	}

	c.Bind(&body)

	user := models.User{FirstName: body.FirstName, LastName: body.LastName, Email: body.Email, UserTickets: body.UserTickets}
	result := initialize.DB.Create(&user)

	TotalTicket = TotalTicket - body.UserTickets

	isValid := len(body.FirstName) > 2 && len(body.LastName) > 2 && strings.Contains(body.Email, "@") && body.UserTickets > 0

	if result.Error != nil {
		c.Status(400)
		return
	}

	if TotalTicket == 0 {
		c.JSON(200, gin.H{
			"Message": "All tickets got booked, Thank you!",
		})
	} else if !isValid {
		c.JSON(200, gin.H{
			"Error Message": "Enter valid input data",
		})
	} else {
		c.JSON(200, gin.H{
			"user":         user,
			"Message":      "Thank you for booking tickets",
			"Tickets left": TotalTicket,
		})
	}
}

func BookedUsers(c *gin.Context) {
	var users []models.User
	initialize.DB.Find(&users)

	c.JSON(200, gin.H{
		"users": users,
	})
}

func UserDetails(c *gin.Context) {
	Email := c.Param("email")
	var user models.User
	initialize.DB.Where("email = ?", Email).First(&user)

	c.JSON(200, gin.H{
		"user": user,
	})
}

func CancelTicket(c *gin.Context) {
	Email := c.Param("email")
	var user models.User
	// var bookedTickets uint

	// initialize.DB.Select("user_tickets").Where("email = ?", Email).First(&bookedTickets)
	// TotalTicket = TotalTicket + bookedTickets

	initialize.DB.Where("email = ?", Email).Delete(&user)

	c.JSON(200, gin.H{
		// "Tickets left": TotalTicket,
		"Message": "Ticket Cancelled for user : " + Email,
	})
}
