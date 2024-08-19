package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Passport struct {
	FullName      string `form:"full_name" binding:"required"`
	DateOfBirth   string `form:"date_of_birth" binding:"required"`
	PassportIssue string `form:"passport_issue" binding:"required"`
	PassportFront string `form:"passport_front" binding:"required"`
	PassportBack  string `form:"passport_back" binding:"required"`
}

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Server is running")
	})

	r.POST("/upload", func(c *gin.Context) {
		var passport Passport

		if err := c.ShouldBind(&passport); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		frontFile, err := c.FormFile("passport_front")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Front photo is required"})
			return
		}

		backFile, err := c.FormFile("passport_back")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Back photo is required"})
			return
		}

		c.SaveUploadedFile(frontFile, "./uploads/"+frontFile.Filename)
		c.SaveUploadedFile(backFile, "./uploads/"+backFile.Filename)

		c.JSON(http.StatusOK, gin.H{
			"full_name":      passport.FullName,
			"date_of_birth":  passport.DateOfBirth,
			"passport_issue": passport.PassportIssue,
			"front_photo":    frontFile.Filename,
			"back_photo":     backFile.Filename,
		})
	})

	r.Run(":8080")
}
